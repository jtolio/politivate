"use strict";

import React, { Component } from 'react';
import { ScrollView, RefreshControl, Linking } from 'react-native';
import { H2, View, Text, Button } from 'native-base';
import { styles, ErrorView } from './common';

export default class LoginView extends Component {
  render() {
    return (
      <View>
        <View style={styles.tabheader}>
          <H2>Politivate</H2>
        </View>
        <Button block onPress={() =>
            Linking.openURL("https://www.politivate.org/app/login").catch(err => {})
        }>Login</Button>
      </View>
    );
  }
}
