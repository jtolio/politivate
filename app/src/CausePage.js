/* @flow */
"use strict";

import React from 'react';
import {
  Linking, View, Text, Image, TouchableOpacity, ScrollView, Button
} from 'react-native';
import Subpage from './Subpage';
import ChallengesList from './ChallengesList';
import { ErrorView, Link, LoadingView } from './common';
import Icon from 'react-native-vector-icons/Entypo';

export default class Cause extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: true,
      cause: null,
      error: null
    };
    this.update = this.update.bind(this);
  }

  componentDidMount() {
    this.update();
  }

  async update() {
    try {
      this.setState({loading: true, error: null});
      let cause = await this.props.appstate.request("GET",
          "/v1/cause/" + this.props.cause.id);
      this.setState({loading: false, cause});
    } catch(error) {
      this.setState({loading: false, error});
    }
  }

  render() {
    if (this.state.loading) {
      return <LoadingView/>;
    }
    if (this.state.error) {
      return <ErrorView msg={this.state.error}/>;
    }
    let cause = this.state.cause;
    return (
      <Subpage appstate={this.props.appstate} title={cause.name}>
        <ScrollView>
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
                <TouchableOpacity onPress={() => Linking.openURL(cause.url).catch(err => {})}>
                  <Link>
                    {cause.url}
                  </Link>
                </TouchableOpacity>
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
        </ScrollView>
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
