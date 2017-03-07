"use strict";

import React from 'react';
import { View } from 'react-native';
import ChallengesList from './ChallengesList';
import { TabHeader } from './common';

export default class ChallengesTab extends React.Component {
  render() {
    return (
      <View style={{flex:1}}>
        <TabHeader>Challenges</TabHeader>
        <ChallengesList
            resourceFn={this.props.appstate.resources.getChallenges}
            appstate={this.props.appstate}/>
      </View>
    );
  }
}
