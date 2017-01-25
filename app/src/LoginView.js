"use strict";

import React, { Component } from 'react';
import { ScrollView, RefreshControl, Linking, View, Text } from 'react-native';
import { Button } from 'native-base';
import { styles, ErrorView } from './common';

export default class LoginView extends Component {
  render() {
    return (
      <View>
        <View style={styles.tabheader}>
          <Text>Politivate</Text>
        </View>
        <Button block onPress={() =>
            Linking.openURL("https://www.politivate.org/app/login").catch(err => {})
        }>Login</Button>
      </View>
    );
  }
}
