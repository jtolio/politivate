"use strict";

import React, { Component } from 'react';
import { H2, ListItem, List, View } from 'native-base';
import { styles, auth } from './common';

export default class CausesTab extends Component {
  componentDidMount() {
    auth.authorize('google')
      .then(resp => console.log(resp))
      .catch(err => console.log(err));
  }

  render() {
    return (
      <View tabLabel={this.props.tabLabel}>
        <View style={styles.tabheader}>
          <H2>Causes</H2>
        </View>
      </View>
    );
  }
}
