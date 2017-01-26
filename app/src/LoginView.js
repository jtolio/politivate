"use strict";

import React, { Component } from 'react';
import { ScrollView, RefreshControl, Linking, View, Text, Button } from 'react-native';
import { styles, ErrorView } from './common';

export default class LoginView extends Component {
  render() {
    return (
      <View>
        <View style={styles.tabheader}>
          <Text>Politivate</Text>
        </View>
        <Button
          onPress={() => Linking.openURL("https://www.politivate.org/app/login").catch(err => {})}
          title="Login" />
      </View>
    );
  }
}
