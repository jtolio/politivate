/* @flow */
"use strict";

import React from 'react';
import {
  Text, Image, View, ScrollView, TouchableOpacity, Button, Linking
} from 'react-native';
import Subpage from './Subpage';
import { Link, ErrorView, LoadingView } from './common';
import FollowButton from './FollowButton';

class ChallengeActions extends React.Component {
  render() {
    let info = this.props.data.challenge.info;
    let data = this.props.data.challenge.data;
    if (data.database == "direct") {
      let action = {"phonecall": "Call", "location": "Check In"}[info.type];
      return (
        <View>
          <View style={{paddingTop: 10}}/>
          <Button title={action} onPress={() => {}}/>
        </View>
      );
    }
    let result = [];
    for (var legislator of this.props.data.legislators) {
      let action = {
          "phonecall": "Call " + legislator.phone,
          "location": "Check In",
        }[info.type];
      let title = {"senate": "Sen.", "house": "Rep."}[legislator.chamber];
      let name = legislator.first_name + " " + legislator.last_name;
      let message = title + " " + name + ": " + action;
      result.push(<View key={"view-" + legislator.votesmart_id}
                      style={{paddingTop: 10}}/>);
      let onPress = () => {};
      if (info.type == "phonecall") {
        onPress = () => {
          Linking.openURL("tel:" + legislator.phone);
        };
      }
      result.push(<Button key={"button-" + legislator.votesmart_id}
                      title={message} onPress={onPress}/>);
    }
    return <View>{result}</View>;
  }
}

export default class ChallengePage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: true,
      data: null,
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
      let data = await this.props.appstate.request("GET",
          "/v1/cause/" + this.props.challenge.cause_id +
          "/challenge/" + this.props.challenge.id);
      this.setState({loading: false, data});
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
    let chal = this.state.data.challenge;
    return (
      <Subpage appstate={this.props.appstate} title={chal.info.title}>
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
                <TouchableOpacity onPress={() => Linking.openURL(
                    this.props.cause.url).catch(err => {})}>
                  <Link>
                    {this.props.cause.url}
                  </Link>
                </TouchableOpacity>
              </View>
              <FollowButton cause={this.props.cause}
                    appstate={this.props.appstate} />
            </View>
            <Text>{chal.data.description}</Text>
            <View style={{paddingTop: 10}}/>
            <ChallengeActions data={this.state.data}/>
          </View>
        </ScrollView>
      </Subpage>
    );
  }
}
