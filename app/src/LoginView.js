/* @flow */
"use strict";

import React, { Component } from 'react';
import {
  ScrollView, RefreshControl, Linking, View, Text, Image
} from 'react-native';
import { ErrorView, TabHeader, styles, Button } from './common';

export default class LoginView extends Component {
  render() {
    return (
      <View style={{flex: 1, flexDirection: "column"}}>
        <View style={[styles.appheader, {flex: 1}]}>
          <Image source={require("../images/header.png")}
                 style={styles.applogo} />
        </View>
        <Button
          onPress={() => Linking.openURL("https://www.politivate.org/app/login").catch(err => {})}
          title="Login" />
      </View>
    );
  }
}
