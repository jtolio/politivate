"use strict";

import React, { Component } from 'react';
import { ScrollView, RefreshControl, View, Text, Button } from 'react-native';
import { styles, ErrorView } from './common';

export default class SettingsTab extends Component {
  render() {
    return (
      <View>
        <View style={styles.tabheader}>
          <Text>Settings</Text>
        </View>
        <Button onPress={this.props.appstate.logout} title="Log out" />
      </View>
    );
  }
}
