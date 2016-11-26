"use strict";

import React, { Component } from 'react';
import { H2, ListItem, List, View } from 'native-base';
import { styles } from './common';

export default class ProfileTab extends Component {
  render() {
    return (
      <View tabLabel={this.props.tabLabel}>
        <View style={styles.tabheader}>
          <H2>Profile</H2>
        </View>
      </View>
    );
  }
}
