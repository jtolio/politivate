/* @flow */
"use strict";

import React, { Component } from 'react';
import { ScrollView, RefreshControl, View, Text } from 'react-native';
import { ErrorView, TabHeader, Button } from './common';

export default class SettingsTab extends Component {
  render() {
    return (
      <View style={{flex: 1}}>
        <TabHeader>Settings</TabHeader>
        <View style={{flex: 1, padding: 20, paddingTop: 5, paddingBottom: 5}}>
          <Button onPress={this.props.appstate.logout} title="Log out" />
        </View>
      </View>
    );
  }
}
