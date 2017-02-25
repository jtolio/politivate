/* @flow */
"use strict";

import React from 'react';
import { Linking, View, Text, Image, Button } from 'react-native';
import Subpage from './Subpage';
import LoadablePage from './LoadablePage';
import ChallengesList from './ChallengesList';
import { Link } from './common';
import Icon from 'react-native-vector-icons/Entypo';

export default class Cause extends React.Component {
  resourceURL() { return "/v1/cause/" + this.props.cause.id; }

  renderLoaded(cause) {
    return (
      <View style={{
          padding: 20
        }}>
        <View style={{
            flexDirection: "row",
            alignItems: "center",
            paddingBottom: 10}}>
          <Image
            source={{uri: cause.icon_url}}
            style={{width: 50, height: 50, borderRadius: 10}}/>
          <View style={{paddingLeft: 10, flex: 1}}>
            <Text style={{fontWeight: "bold"}}>{cause.name}</Text>
            <Link url={cause.url}>
              {cause.url}
            </Link>
          </View>
          {this.props.followButton}
        </View>
        <Text>{cause.description}</Text>
        <View style={{paddingTop: 20}}/>
        <Button title="Challenges"
            onPress={() => this.props.appstate.navigator.push(
              {component: CauseChallengePage, passProps: {cause}})}/>
        <View style={{paddingTop: 10}}/>
        <Button title="Leaderboard"
            onPress={() => this.props.appstate.navigator.push(
              {component: LeaderboardPage, passProps: {cause}})}/>
      </View>
    );
  }

  render() {
    return (
      <Subpage appstate={this.props.appstate} title={this.props.cause.name}>
        <LoadablePage renderLoaded={this.renderLoaded.bind(this)}
                      resourceURL={this.resourceURL()}
                      appstate={this.props.appstate}/>
      </Subpage>
    );
  }
}

class LeaderboardPage extends React.Component {
  render() {
    let cause = this.props.cause;
    return (
      <Subpage appstate={this.props.appstate} title={cause.name}>
        <View style={{flex: 1, alignItems: "center", paddingTop: 30}}>
          <Text style={{fontWeight: "bold"}}>
            Leaderboard coming soon!
          </Text>
        </View>
      </Subpage>
    );
  }
}

class CauseChallengePage extends React.Component {
  render() {
    let cause = this.props.cause;
    return (
      <Subpage appstate={this.props.appstate} title={cause.name}>
        <ChallengesList
            resource={"/v1/cause/" + cause.id + "/challenges/"}
            appstate={this.props.appstate}>
          <Text style={{fontWeight: "bold"}}>
            This cause has not listed any challenges yet!
          </Text>
        </ChallengesList>
      </Subpage>
    );
  }
}
