"use strict";

import React from 'react';
import { Text, View } from 'react-native';
import ScrollableTabView from 'react-native-scrollable-tab-view';
import DefaultTabBar from 'react-native-scrollable-tab-view/DefaultTabBar';
import Button from 'react-native-scrollable-tab-view/Button';
import CausesTab from './CausesTab';
import ChallengesTab from './ChallengesTab';
import ProfileTab from './ProfileTab';
import SettingsTab from './SettingsTab';
import { styles, theme } from './common';
import Icon from 'react-native-vector-icons/Entypo';

export default class Tabs extends React.Component {
  constructor(props) {
    super(props);
    this.renderTab = this.renderTab.bind(this);
    this.icons = {
      "Challenges": "stopwatch",
      "Causes": "awareness-ribbon",
      "Profile": "v-card",
      "Settings": "cog"};
  }

  renderTab(name, page, isTabActive, onPressHandler) {
    const textColor = (isTabActive ?
        theme.topTabBarActiveTextColor : theme.topTabBarTextColor);
    const fontWeight = (isTabActive ? "bold" : "normal");
    return (<Button
                style={{flex: 1}}
                key={name}
                accessible={true}
                accessibilityLabel={name}
                accessibilityTraits="button"
                onPress={() => onPressHandler(page)}>
              <View
                  style={[{flex: 1,
                           alignItems: "center",
                           justifyContent: "center",
                           paddingBottom: 10,
                           paddingTop: 5}, styles.tabBar]}>
                <Icon name={this.icons[name]} size={30}
                  style={[{color: textColor}, theme.tabBarText]}/>
              </View>
            </Button>);
  }

  render() {
    return (
      <View style={{flex:1}}>
        <View style={styles.appheader} alignItems="center">
          <Text style={styles.appheadertext}>Politivate</Text>
        </View>
        <ScrollableTabView
            tabBarPosition="bottom"
            locked={true}
            renderTabBar={(props) => <DefaultTabBar
                    renderTab={this.renderTab} {...props}/>}
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
