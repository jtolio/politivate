/* @flow */
"use strict";

import React from 'react';
import {
  Linking, View, Text, Image, TouchableOpacity, ScrollView, Button
} from 'react-native';
import Subpage from './Subpage';
import ChallengesList from './ChallengesList';
import { ErrorView, Link } from './common';
import Icon from 'react-native-vector-icons/Entypo';

export default class Cause extends React.Component {
  render() {
    let row = this.props.cause;
    return (
      <Subpage appstate={this.props.appstate} title={row.name}>
        <ScrollView>
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
                <TouchableOpacity onPress={() => Linking.openURL(row.url).catch(err => {})}>
                  <Link>
                    {row.url}
                  </Link>
                </TouchableOpacity>
              </View>
              {this.props.followButton}
            </View>
            <Text>{row.description}</Text>
            <View style={{paddingTop: 20}}/>
            <Button title="Challenges"
                onPress={() => this.props.appstate.navigator.push(
                  {component: CauseChallengePage, passProps: {
                    cause: row}})}/>
            <View style={{paddingTop: 10}}/>
            <Button title="Leaderboard"
                onPress={() => this.props.appstate.navigator.push(
                  {component: LeaderboardPage, passProps: {
                    cause: row}})}/>
          </View>
        </ScrollView>
      </Subpage>
    );
  }
}

class LeaderboardPage extends React.Component {
  render() {
    let row = this.props.cause;
    return (
      <Subpage appstate={this.props.appstate} title={row.name}>
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
    return (
      <Subpage appstate={this.props.appstate} title={this.props.cause.name}>
        <ChallengesList
            resource={"/v1/cause/" + this.props.cause.id + "/challenges/"}
            appstate={this.props.appstate}>
          <Text style={{fontWeight: "bold"}}>
            This cause has not listed any challenges yet!
          </Text>
        </ChallengesList>
      </Subpage>
    );
  }
}
