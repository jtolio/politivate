"use strict";

import React from 'react';
import { Text, Image, TouchableOpacity, View } from 'react-native';
import ChallengePage from './ChallengePage';
import ListTab from './ListTab';

export default class ChallengesTab extends React.Component {
  constructor(props) {
    super(props);
    this.renderRow = this.renderRow.bind(this);
  }

  renderRow(row) {
    return (
      <TouchableOpacity onPress={() => this.props.navigator.push({
              component: ChallengePage, passProps: {challenge: row}})}>
        <View>
          {(row.icon_url != "") ?
            <Image source={{uri: row.icon_url}} /> : null}
          <Text>{row.title}</Text>
          <Text>{row.description}</Text>
        </View>
      </TouchableOpacity>
    );
  }

  render() {
    return (
      <ListTab url="https://www.politivate.org/api/v1/challenges/"
        header="Challenges" renderRow={this.renderRow}
        appstate={this.props.appstate} />
    );
  }
}
