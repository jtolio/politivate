/* @flow */
"use strict";

import React from 'react';
import {
  Text, Image, View, ScrollView, TouchableOpacity, Button
} from 'react-native';
import Subpage from './Subpage';
import { Link } from './common';
import FollowButton from './FollowButton';

export default class ChallengePage extends React.Component {
  render() {
    let row = this.props.challenge;
    return (
      <Subpage appstate={this.props.appstate} title={row.title}>
        <ScrollView>
          <View style={{
              padding: 20
            }}>
            <View style={{
                flexDirection: "row",
                alignItems: "center",
                paddingBottom: 10}}>
              <Image
                source={{uri: this.props.cause.icon_url}}
                style={{width: 50, height: 50, borderRadius: 10}}/>
              <View style={{paddingLeft: 10, flex: 1}}>
                <Text style={{fontWeight: "bold"}}>{this.props.cause.name}</Text>
                <TouchableOpacity onPress={() => Linking.openURL(
                    this.props.cause.url).catch(err => {})}>
                  <Link>
                    {this.props.cause.url}
                  </Link>
                </TouchableOpacity>
              </View>
              <FollowButton cause={this.props.cause}
                    appstate={this.props.appstate} />
            </View>
            <Text>{row.description}</Text>
            <View style={{paddingTop: 20}}/>
            <Button
              title={{"phonecall": "Call", "location": "Check In"}[this.props.challenge.type]}
              onPress={() => {}}/>
          </View>
        </ScrollView>
      </Subpage>
    );
  }
}
