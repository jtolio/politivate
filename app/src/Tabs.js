"use strict";

import React from 'react';
import { Text, View, Image, StyleSheet } from 'react-native';
import ScrollableTabView from 'react-native-scrollable-tab-view';
import DefaultTabBar from 'react-native-scrollable-tab-view/DefaultTabBar';
import Button from 'react-native-scrollable-tab-view/Button';
import CausesTab from './CausesTab';
import ChallengesTab from './ChallengesTab';
import ProfileTab from './ProfileTab';
import SettingsTab from './SettingsTab';
import { colors, AppHeader, Separator } from './common';
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
        colors.secondary.val : colors.primary.val);
    const fontWeight = (isTabActive ? "bold" : "normal");
    return (<Button
                style={{flex: 1}}
                key={"tab-" + name}
                accessible={true}
                accessibilityLabel={name}
                accessibilityTraits="button"
                onPress={() => onPressHandler(page)}>
              <View style={styles.tabBar}>
                <Icon name={this.icons[name]} size={30}
                    style={[{color: textColor}]}/>
              </View>
            </Button>);
  }

  render() {
    return (
      <View style={{flex:1}}>
        <AppHeader/>
        <Separator/>
        <ScrollableTabView
            tabBarPosition="bottom"
            locked={true}
            renderTabBar={(props) => <DefaultTabBar
                    renderTab={this.renderTab} {...props}/>}
            tabBarUnderlineStyle={styles.tabBarUnderline}>
          <ChallengesTab
              tabLabel="Challenges" appstate={this.props.appstate}/>
          <CausesTab
              tabLabel="Causes" appstate={this.props.appstate}/>
          <ProfileTab
              tabLabel="Profile" appstate={this.props.appstate}/>
          <SettingsTab
              tabLabel="Settings" appstate={this.props.appstate}/>
        </ScrollableTabView>
      </View>
    );
  }
}

var styles = StyleSheet.create({
  tabBarUnderline: {
    backgroundColor: colors.secondary.val,
  },
  tabBar: {
    borderWidth: 1,
    borderBottomWidth: 0,
    borderLeftWidth: 0,
    borderRightWidth: 0,
    borderColor: colors.primary.val,
    flex: 1,
    alignItems: "center",
    justifyContent: "center",
    paddingBottom: 10,
    paddingTop: 5,
  },
});
