"use strict";

import React, { Component } from 'react';
import { View } from 'react-native';
import {
  Container, Header, Button, Title, Content, Card, CardItem, Icon, Text,
  Thumbnail
} from 'native-base';
import { styles } from './common';

export default class Challenge extends Component {
  render() {
    let row = this.props.challenge;
    return (
      <Container>
        <Header>
          <Button transparent onPress={this.props.backPress}>
            <Icon name='ios-arrow-back' />
          </Button>
          <Title>{row.cause.name}</Title>
        </Header>
        <Content>
          <Card style={{flex:1}}>
            <CardItem header>
              <Thumbnail source={{uri: row.icon}} />
              <Text>{row.title}</Text>
            </CardItem>
            <CardItem>
              <Text>{row.short_desc}</Text>
            </CardItem>
          </Card>
        </Content>
      </Container>
    );
  }
}
