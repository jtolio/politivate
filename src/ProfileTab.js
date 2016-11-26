"use strict";

import React, { Component } from 'react';
import { H1, ListItem, List, View } from 'native-base';
import { styles } from './common';

export default class ProfileTab extends Component {
  render() {
    return (
      <View tabLabel={this.props.tabLabel}>
        <View style={styles.header}>
          <H1>Profile</H1>
        </View>
      </View>
    );
  }
}
