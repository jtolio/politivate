"use strict";

import React from 'react';
import { Text } from 'react-native';
import { Card, CardItem, Thumbnail } from 'native-base';
import ChallengePage from './ChallengePage';
import ListTab from './ListTab';

export default class ChallengesTab extends React.Component {
  constructor(props) {
    super(props);
    this.renderRow = this.renderRow.bind(this);
  }

  renderRow(row) {
    return (
      <Card>
        <CardItem button header onPress={() => this.props.navigator.push({
              component: ChallengePage, passProps: {challenge: row}})}>
          {(row.icon_url != "") ?
            <Thumbnail source={{uri: row.icon_url}} /> : null}
          <Text>{row.title}</Text>
        </CardItem>
        <CardItem>
          <Text>{row.description}</Text>
        </CardItem>
      </Card>
    );
  }

  render() {
    return (
      <ListTab url="https://www.politivate.org/api/v1/challenges/"
        header={<Text>Challenges</Text>} renderRow={this.renderRow}
        appstate={this.props.appstate} />
    );
  }
}
