"use strict";

import React, { Component } from 'react';
import { ScrollView, RefreshControl } from 'react-native';
import { H2, View, Text, Button } from 'native-base';
import { styles, ErrorView } from './common';

export default class SettingsTab extends Component {
  render() {
    return (
      <View>
        <View style={styles.tabheader}>
          <H2>Settings</H2>
        </View>
        <Button block onPress={this.props.appstate.logout}>Log out</Button>
      </View>
    );
  }
}
