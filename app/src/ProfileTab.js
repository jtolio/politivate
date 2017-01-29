"use strict";

import React, { Component } from 'react';
import { ScrollView, RefreshControl, View, Text, Image } from 'react-native';
import { ErrorView, TabHeader } from './common';

export default class ProfileTab extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: true,
      profile: null,
      error: null
    };
    this.update = this.update.bind(this);
    this.renderLoaded = this.renderLoaded.bind(this);
  }

  componentDidMount() {
    this.update();
  }

  async update() {
    try {
      this.setState({loading: true, error: null});
      let profile = await this.props.appstate.request("GET", "/v1/profile");
      this.setState({loading: false, profile});
    } catch(error) {
      this.setState({loading: false, error});
    }
  }

  renderLoaded() {
    return (
      <View style={{
          padding: 20,
          paddingTop: 5,
          paddingBottom: 5
        }}>
        <View style={{
            flexDirection: "row",
            alignItems: "center",
            paddingBottom: 10}}>
          <Image
            source={{uri: this.state.profile.avatar_url}}
            style={{width: 50, height: 50, borderRadius: 10}}/>
          <View style={{paddingLeft: 10}}>
            <Text style={{fontWeight: "bold"}}>{this.state.profile.name}</Text>
          </View>
        </View>
        <Text>Profile info</Text>
      </View>
    );
  }

  render() {
    if (this.state.error) {
      return <ErrorView msg={this.state.error}/>;
    }
    return (
      <View style={{flex: 1}}>
        <TabHeader>Profile</TabHeader>
        <ScrollView refreshControl={
            <RefreshControl refreshing={this.state.loading}
                            onRefresh={this.update}/>}>
          { this.state.loading ? null : this.renderLoaded() }
        </ScrollView>
      </View>
    );
  }
}
