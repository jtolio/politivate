"use strict";

import React from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import { ErrorView, colors } from './common';
import Icon from 'react-native-vector-icons/Entypo';

export default class FollowButton extends React.Component {

  constructor(props) {
    super(props);
    this.state = {error: null, following: false, loading: true, followers: 0};
    this.update = this.update.bind(this);
    this.toggle = this.toggle.bind(this);
    this.request = this.request.bind(this);
  }

  componentDidMount() {
    this.update();
  }

  update() {
    this.request(this.props.appstate.resources.causeFollowers);
  }

  async request(method) {
    try {
      this.setState({loading: true, error: null});
      let resp = await method(this.props.cause.id);
      this.setState({
        loading: false,
        followers: resp.followers,
        following: resp.following});
    } catch(error) {
      this.setState({loading: false, error});
    }
  }

  toggle() {
    if (this.state.following) {
      this.request(this.props.appstate.resources.unfollowCause);
    } else {
      this.request(this.props.appstate.resources.followCause);
    }
  }

  render() {
    if (this.state.error) {
      return <ErrorView msg={this.state.error}/>;
    }
    return (<TouchableOpacity onPress={this.toggle}>
              { (this.state.loading ?
                (<Text>
                  <Icon name="hour-glass" size={30} />
                 </Text>) :
                (this.state.following ?
                <Text style={{color: colors.heart.val}}>
                  <Icon name="heart" size={30} style={{color: colors.heart.val}}/>
                </Text>
                : <Text>
                  <Icon name="heart-outlined" size={30}/>
                </Text>)) }
            </TouchableOpacity>);
  }
}
