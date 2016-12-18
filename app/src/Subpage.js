"use strict";

import React from 'react';
import { Container, Header, Button, Title, Content, Icon } from 'native-base';

export default class Subpage extends React.Component {
  render() {
    return (
      <Container>
        <Header>
          <Button transparent onPress={this.props.backPress}>
            <Icon name='ios-arrow-back' />
          </Button>
          <Title>{this.props.title}</Title>
        </Header>
        <Content>
          {this.props.children}
        </Content>
      </Container>
    );
  }
}
