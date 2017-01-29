/* @flow */
"use strict";

import React, { Component } from 'react';
import {
  ScrollView, RefreshControl, Linking, View, Text, Image
} from 'react-native';
import { ErrorView, AppHeader, Button } from './common';

export default class LoginView extends Component {
  render() {
    return (
      <View style={{flex: 1, flexDirection: "column"}}>
        <AppHeader/>
        <View style={{flex: 1}}></View>
        <Button
          onPress={() => Linking.openURL("https://www.politivate.org/app/login").catch(err => {})}
          title="Login" />
      </View>
    );
  }
}
