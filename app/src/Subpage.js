/* @flow */
"use strict";

import React from 'react';
import { View, TouchableOpacity, Text } from 'react-native';
import { colors } from './common';
import Icon from 'react-native-vector-icons/Entypo';

export default class Subpage extends React.Component {
  render() {
    return (
      <View style={{
          flexDirection: "column", flex: 1}}>
        <View style={{
              flexDirection: "row",
              justifyContent: "center",
              borderBottomWidth: 1,
              borderColor: colors.primary.val}}>
          <Text style={{
              fontWeight: "bold",
              fontSize: 20, padding: 10, paddingLeft: 35, paddingRight: 35,
              color: colors.primary.val}}>
            {this.props.title}
          </Text>
          <TouchableOpacity onPress={this.props.appstate.backPress}
              style={{
                  position: "absolute",
                  top: 0, left: 0, bottom: 0,
                  flexDirection: "column", justifyContent: "center",
                  paddingLeft: 10, paddingRight: 10}}>
            <Icon name='chevron-thin-left' style={{
                fontSize: 30, padding: 0,
                color: colors.primary.val}}/>
          </TouchableOpacity>
        </View>
        <View style={{flex: 1}}>
          {this.props.children}
        </View>
      </View>
    );
  }
}
