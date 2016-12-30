"use strict";

import React from 'react';
import { H2, Text, Card, CardItem, Thumbnail } from 'native-base';
import ListTab from './ListTab';
import CausePage from './CausePage';
import FollowButton from './FollowButton';

export default class CausesTab extends React.Component {
  constructor(props) {
    super(props);
    this.renderRow = this.renderRow.bind(this);
  }

  renderRow(row) {
    let followButton = (
      <FollowButton cause={row} appstate={this.props.appstate} />
    );
    return (
      <Card>
        <CardItem header button onPress={() => this
          .props.navigator.push({component: CausePage, passProps: {
              cause: row, followButton: followButton}})}>
          {(row.icon_url != "") ?
            <Thumbnail source={{uri: row.icon_url}} /> : null}
          <Text>{row.name}</Text>
          {followButton}
        </CardItem>
      </Card>
    );
  }

  render() {
    return (
      <ListTab url="https://www.politivate.org/api/v1/causes/"
        header={<H2>Causes</H2>} renderRow={this.renderRow}
        appstate={this.props.appstate} />
    );
  }
}
