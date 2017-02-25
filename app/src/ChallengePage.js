/* @flow */
"use strict";

import React from 'react';
import { Text, Image, View, ScrollView, Button } from 'react-native';
import Subpage from './Subpage';
import LoadablePage from './LoadablePage';
import { Link, ErrorView, LoadingView, phonecall } from './common';
import FollowButton from './FollowButton';

class ChallengeActions extends React.Component {
  render() {
    let chal = this.props.challenge;
    if (chal.database == "direct") {
      let action = {"phonecall": "Call", "location": "Check In"}[chal.type];
      return (
        <View>
          <View style={{paddingTop: 10}}/>
          <Button title={action} onPress={() => {}}/>
        </View>
      );
    }
    let result = [];
    for (var legislator of chal.legislators) {
      let phonenumber = legislator.phone;
      let action = {
          "phonecall": "Call " + phonenumber,
          "location": "Check In",
        }[chal.type];
      let title = {"senate": "Sen.", "house": "Rep."}[legislator.chamber];
      let name = legislator.first_name + " " + legislator.last_name;
      let message = title + " " + name + ": " + action;
      result.push(<View key={"view-" + legislator.votesmart_id}
                      style={{paddingTop: 10}}/>);
      let onPress = () => {};
      if (chal.type == "phonecall") {
        onPress = () => { phonecall(phonenumber); };
      }
      result.push(<Button key={"button-" + legislator.votesmart_id}
                      title={message} onPress={onPress}/>);
    }
    return <View>{result}</View>;
  }
}

export default class ChallengePage extends React.Component {
  resourceURL() {
    return "/v1/cause/" + this.props.challenge.cause_id +
           "/challenge/" + this.props.challenge.id;
  }

  renderLoaded(chal) {
    return (
      <View style={{
          padding: 20
        }}>
        <View style={{
            flexDirection: "row",
            alignItems: "center",
            paddingBottom: 10}}>
          <Image
            source={{uri: this.props.cause.icon_url}}
            style={{width: 50, height: 50, borderRadius: 10}}/>
          <View style={{paddingLeft: 10, flex: 1}}>
            <Text style={{fontWeight: "bold"}}>{this.props.cause.name}</Text>
            <Link url={this.props.cause.url}>
              {this.props.cause.url}
            </Link>
          </View>
          <FollowButton cause={this.props.cause}
                appstate={this.props.appstate} />
        </View>
        <Text>{chal.description}</Text>
        <View style={{paddingTop: 10}}/>
        <ChallengeActions challenge={chal}/>
      </View>
    );
  }

  render() {
    return (
      <Subpage appstate={this.props.appstate}
               title={this.props.challenge.title}>
        <LoadablePage renderLoaded={this.renderLoaded.bind(this)}
                      resourceURL={this.resourceURL()}
                      appstate={this.props.appstate}/>
      </Subpage>
    );
  }
}
