"use strict";

import React, { Component } from 'react';
import { ScrollView, RefreshControl, View, Text } from 'react-native';
import { Button } from 'native-base';
import { styles, ErrorView } from './common';

export default class SettingsTab extends Component {
  render() {
    return (
      <View>
        <View style={styles.tabheader}>
          <Text>Settings</Text>
        </View>
        <Button block onPress={this.props.appstate.logout}>Log out</Button>
      </View>
    );
  }
}
