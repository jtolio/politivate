/* @flow */
"use strict";

import React from 'react';
import { View, TouchableOpacity, Text } from 'react-native';
import Icon from 'react-native-vector-icons/Ionicons';

export default class Subpage extends React.Component {
  render() {
    return (
      <View style={{flexDirection: "column", flex: 1}}>
        <View style={{flexDirection: "row", justifyContent: "center"}}>
          <TouchableOpacity onPress={this.props.backPress}
              style={{position: "absolute", top: 0, left: 0}}>
            <Icon name='ios-arrow-back' style={{fontSize: 30, padding: 10}}/>
          </TouchableOpacity>
          <Text style={{fontSize: 20, padding: 10}}>{this.props.title}</Text>
        </View>
        <View style={{flex: 1}}>
          {this.props.children}
        </View>
      </View>
    );
  }
}
