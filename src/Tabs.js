"use strict";

import React, { Component } from 'react';
import { H1, View } from 'native-base';
var ScrollableTabView = require('react-native-scrollable-tab-view');
import CausesTab from './CausesTab';
import ChallengesTab from './ChallengesTab';
import ProfileTab from './ProfileTab';
import { styles } from './common';

export default class Tabs extends Component {
  render() {
    return (
      <View>
        <View style={styles.header}>
          <H1>PolitiForce</H1>
        </View>
        <ScrollableTabView
            tabBarPosition="bottom"
            locked={true}>
          <ChallengesTab
              tabLabel="Challenges"
              navigator={this.props.navigator}/>
          <CausesTab
              tabLabel="Causes"
              navigator={this.props.navigator}/>
          <ProfileTab
              tabLabel="Profile"
              navigator={this.props.navigator}/>
        </ScrollableTabView>
      </View>
    );
  }
}
