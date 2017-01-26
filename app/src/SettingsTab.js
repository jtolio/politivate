/* @flow */
"use strict";

import React, { Component } from 'react';
import { ScrollView, RefreshControl, View, Text } from 'react-native';
import { ErrorView, TabHeader, Button } from './common';

export default class SettingsTab extends Component {
  render() {
    return (
      <View>
        <TabHeader>Settings</TabHeader>
        <Button onPress={this.props.appstate.logout} title="Log out" />
      </View>
    );
  }
}
