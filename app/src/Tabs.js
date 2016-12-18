"use strict";

import React, { Component } from 'react';
import { Text } from 'react-native';
import { H1, View } from 'native-base';
var ScrollableTabView = require('react-native-scrollable-tab-view');
import CausesTab from './CausesTab';
import ChallengesTab from './ChallengesTab';
import ProfileTab from './ProfileTab';
import SettingsTab from './SettingsTab';
import { styles, theme } from './common';

export default class Tabs extends Component {
  render() {
    return (
      <View>
        <View style={styles.appheader} alignItems="center">
          <Text style={styles.appheadertext}>Politivate</Text>
        </View>
        <ScrollableTabView
            tabBarPosition="bottom"
            locked={true}
            tabBarUnderlineStyle={styles.tabBarUnderline}
            tabBarBackgroundColor={theme.tabDefaultBg}
            tabBarActiveTextColor={theme.topTabBarActiveTextColor}
            tabBarInactiveTextColor={theme.topTabBarTextColor}
            tabBarTextStyle={styles.tabBarText}
            style={styles.tabBar}>
          <ChallengesTab
              tabLabel="Challenges" appstate={this.props.appstate}
              navigator={this.props.navigator}/>
          <CausesTab
              tabLabel="Causes" appstate={this.props.appstate}
              navigator={this.props.navigator}/>
          <ProfileTab
              tabLabel="Profile" appstate={this.props.appstate}
              navigator={this.props.navigator}/>
          <SettingsTab
              tabLabel="Settings" appstate={this.props.appstate}
              navigator={this.props.navigator}/>
        </ScrollableTabView>
      </View>
    );
  }
}
