"use strict";

import React from 'react';
import { Text, Image, TouchableOpacity, View } from 'react-native';
import ChallengePage from './ChallengePage';
import List from './List';
import { LoadingView, ErrorView, colors, TabHeader } from './common';
import Icon from 'react-native-vector-icons/Entypo';

const MONTHS = ["JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG",
                "SEP", "OCT", "NOV", "DEC"];

class ChallengeStat extends React.Component {
  render() {
    return (
      <View style={[{borderWidth: 1, borderColor: colors.secondary.val,
                     justifyContent: "center", alignItems: "center",
                     width: 50, height: 50, borderRadius: 10,
                     margin: 1},
                   this.props.style]}>
        {this.props.children}
      </View>
    );
  }
}

class ChallengeStatType extends React.Component {
  render() {
    if (!this.props.icon) {
      return null;
    }
    return (
      <ChallengeStat>
        <Icon name={this.props.icon} size={50}
              style={{color: colors.secondary.val}}/>
      </ChallengeStat>
    );
  }
}

class ChallengeStatPoints extends React.Component {
  render() {
    if (this.props.points <= 0) {
      return null;
    }
    return (
      <ChallengeStat>
        <Text style={{
            fontSize: 30, lineHeight: 30,
            color: colors.secondary.val}}>
          {this.props.points}
        </Text>
        <Text style={{
            fontSize: 8, lineHeight: 8,
            color: colors.secondary.val}}>
          POINTS
        </Text>
      </ChallengeStat>
    );
  }
}

class ChallengeStatDeadline extends React.Component {
  render() {
    if (!this.props.deadline) {
      return null;
    }
    let d = new Date(this.props.deadline);
    if (!d.getDate()) {
      return null;
    }
    return (
      <ChallengeStat>
        <Text style={{
            fontSize: 15, lineHeight: 13, color: colors.secondary.val}}>
          {MONTHS[d.getMonth()]}
        </Text>
        <Text style={{
            fontSize: 25, lineHeight: 25, color: colors.secondary.val}}>
          {d.getDate()}
        </Text>
      </ChallengeStat>
    );
  }
}

class ChallengeEntry extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: true,
      cause: null,
      error: null
    };
    this.typeIcon = {
                      "phonecall": "phone",
                      "location": "location-pin",
                    }[this.props.challenge.type];
    this.update = this.update.bind(this);
  }

  componentDidMount() {
    this.update();
  }

  async update() {
    try {
      this.setState({loading: true, error: null});
      let cause = await this.props.appstate.request(
          "GET", "/v1/cause/" + this.props.challenge.cause_id)
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
    let chal = this.props.challenge;
    let cause = this.state.cause;

    return (
      <TouchableOpacity onPress={() => this.props.appstate.navigator.push({
              component: ChallengePage,
              passProps: {challenge: chal, cause: cause}})}>
        <View style={{
            flexDirection: "row",
            alignItems: "center",
            flex: 1}}>
          { cause.icon_url ?
            <Image source={{uri: cause.icon_url}}
                   style={{width: 50, height: 50, borderRadius: 10}}/> :
            <Icon name="awareness-ribbon" size={50}
                  style={{backgroundColor: colors.primary.val,
                          color: colors.background.val,
                          borderRadius: 10}}/> }
          <View style={{paddingLeft: 10, flex: 1}}>
            <Text style={{fontWeight: "bold"}}>{chal.title}</Text>
          </View>
          <ChallengeStatType icon={this.typeIcon}/>
          <ChallengeStatPoints points={chal.points}/>
          <ChallengeStatDeadline deadline={chal.event_end}/>
        </View>
      </TouchableOpacity>
    );
  }
}

export default class ChallengesList extends React.Component {
  constructor(props) {
    super(props);
    this.renderRow = this.renderRow.bind(this);
  }

  renderRow(row) {
    return <ChallengeEntry appstate={this.props.appstate} challenge={row}/>;
  }

  render() {
    return (
      <List resource={this.props.resource} renderRow={this.renderRow}
            appstate={this.props.appstate}>
        <View style={{flex: 1, alignItems: "center", paddingTop: 30}}>
          {this.props.children ? this.props.children : [
            (<Text key={0} style={{fontWeight: "bold"}}>
               There are no challenges available!
             </Text>),
            (<Text key={1}>
               Consider following some more causes?
             </Text>)]}
        </View>
      </List>
    );
  }
}
