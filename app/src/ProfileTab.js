"use strict";

import React from 'react';
import LoadablePage from './LoadablePage';
import { View, Text, Image } from 'react-native';
import { TabHeader, colors } from './common';

class NutritionFacts extends React.Component {
  render() {
    return (
      <View style={{paddingBottom: 20}}>
        <View style={{borderWidth: 1, borderColor: colors.primary.val,
                      borderRadius: 10, padding: 10, paddingTop: 6}}>
          <Text style={{fontWeight: "bold", fontSize: 25,
                        color: colors.primary.val}}>{this.props.header}</Text>
          {Object.keys(this.props.values).map((cause_id) => (
            <View style={{justifyContent: "space-between", flexDirection: "row",
                          borderTopWidth: 1, borderColor: colors.primary.val}}
                  key={cause_id}>
              <Text style={{fontWeight: "bold", fontSize: 20}}>
                {this.props.causes[cause_id].cause.name}
              </Text>
              <Text style={{fontWeight: "bold", fontSize: 20}}>
                {this.props.values[cause_id]}
              </Text>
            </View>
          ))}
        </View>
      </View>
    );
  }
}

export default class ProfileTab extends React.Component {
  renderHeader() { return <TabHeader>Profile</TabHeader>; }

  async process(profile) {
    var causes = {};
    if (profile.month_actions) {
      for (var action of profile.month_actions) {
        causes[action.cause_id] = null;
      }
      for (var cause_id of Object.keys(causes)) {
        causes[cause_id] = {
          cause: (await this.props.appstate.resources.getPartialCause(
                                                          cause_id)),
          challenges: {},
        };
      }
      for (var action of profile.month_actions) {
        causes[action.cause_id].challenges[action.challenge_id] = null;
      }
      for (var cause_id of Object.keys(causes)) {
        for (var challenge_id of Object.keys(causes[cause_id].challenges)) {
          causes[cause_id].challenges[challenge_id] = (
              await this.props.appstate.resources.getPartialChallenge(
                                                challenge_id, cause_id));
        }
      }
    }
    return {causes, profile};
  }

  renderLoaded(result) {
    let causes = result.causes;
    let profile = result.profile;

    let phonecalls = {};
    let phonechals = {};
    let events = {};
    let checkins = {};
    let challenges = {};

    for (var cause_id of Object.keys(causes)) {
      let cause = causes[cause_id];
      for (var chal_id of Object.keys(cause.challenges)) {
        let chal = cause.challenges[chal_id];
        if (chal.type == "phonecall") {
          if (!phonechals[cause_id]) {
            phonechals[cause_id] = 1;
          } else {
            phonechals[cause_id] += 1;
          }
        } else if (chal.type == "location") {
          if (!events[cause_id]) {
            events[cause_id] = 1;
          } else {
            events[cause_id] += 1;
          }
        }
        if (!challenges[cause_id]) {
          challenges[cause_id] = 1;
        } else {
          challenges[cause_id] += 1;
        }
      }
    }

    for (var action of profile.month_actions) {
      let chal = causes[action.cause_id].challenges[action.challenge_id];
      if (chal.type == "phonecall") {
        if (!phonecalls[action.cause_id]) {
          phonecalls[action.cause_id] = 1;
        } else {
          phonecalls[action.cause_id] += 1;
        }
      } else if (chal.type == "location") {
        if (!checkins[action.cause_id]) {
          checkins[action.cause_id] = 1;
        } else {
          checkins[action.cause_id] += 1;
        }
      }
    }

    return (
      <View style={{
          padding: 20,
          paddingTop: 5,
          paddingBottom: 5
        }}>
        <View style={{
            flexDirection: "row",
            alignItems: "center",
            paddingBottom: 30}}>
          <Image
            source={{uri: profile.avatar_url}}
            style={{width: 50, height: 50, borderRadius: 10}}/>
          <View style={{paddingLeft: 10}}>
            <Text style={{fontWeight: "bold"}}>{profile.name}</Text>
            <Text>Profile text!</Text>
          </View>
        </View>
        { Object.keys(events).length > 0 ?
          <NutritionFacts
              header="Challenges this month"
              causes={causes} values={challenges}/> : null }
        { Object.keys(checkins).length > 0 ?
          <NutritionFacts
              header="Checkins this month"
              causes={causes} values={checkins}/> : null }
        { Object.keys(phonecalls).length > 0 ?
          <NutritionFacts
              header="Calls this month"
              causes={causes} values={phonecalls}/> : null }
      </View>
    );
  }

  render() {
    return (
      <LoadablePage renderLoaded={this.renderLoaded.bind(this)}
        renderHeader={this.renderHeader.bind(this)}
        process={this.process.bind(this)} appstate={this.props.appstate}
        resourceFn={() =>
            this.props.appstate.resources.request("GET", "/v1/profile")}/>
    );
  }
}
