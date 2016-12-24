"use strict";

import React from 'react';
import { Linking, ActivityIndicator } from 'react-native';
import { Card, CardItem, Text, Thumbnail, Button, View } from 'native-base';
import Subpage from './Subpage';
import { ErrorView, Link } from './common';
import Icon from 'react-native-vector-icons/Entypo';

class FollowButton extends React.Component {
  render() {
    return (<Button transparent onPress={this.props.onPress}>
              { (this.props.loading ?
                (<Text>
                  <Icon name="hour-glass" size={30} />
                  .
                 </Text>) :
                (this.props.following ?
                <Text>
                  <Icon name="heart" size={30} />
                  {this.props.followers}
                </Text>
                : <Text>
                  <Icon name="heart-outlined" size={30} />
                  {this.props.followers}
                </Text>)) }
            </Button>);
  }
}

export default class CauseHeader extends React.Component {
  constructor(props) {
    super(props);
    this.state = {error: null, following: false, loading: true, followers: 0};
    this.update = this.update.bind(this);
    this.toggle = this.toggle.bind(this);
  }

  componentDidMount() {
    this.update();
  }

  update() {
    this.request("GET");
  }

  request(method) {
    let req = new Request("https://www.politivate.org/api/v1/cause/" +
                          this.props.cause.id + "/followers",
                          {method});
    fetch(req)
      .then((resp) => resp.json())
      .then((json) => this.setState({
                  loading: false,
                  error: null,
                  followers: json.resp.followers,
                  following: json.resp.following}))
      .catch((error) => this.setState({
                  loading: false,
                  error: error}));
  }

  toggle() {
    if (this.state.following) {
      this.request("DELETE");
    } else {
      this.request("POST");
    }
  }

  render() {
    if (this.state.error) {
      return <ErrorView msg={this.state.error}/>;
    }
    let row = this.props.cause;
    return (
      <CardItem header {...this.props}>
        {(row.icon_url != "") ?
          <Thumbnail source={{uri: row.icon_url}} /> : null}
        <Text>{row.name}</Text>
        <FollowButton onPress={() => this.toggle()}
                  loading={this.state.loading}
                  following={this.state.following}
                  followers={this.state.followers} />
      </CardItem>);
  }
}
