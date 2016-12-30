"use strict";

import React, { Component } from 'react';
import { ScrollView, RefreshControl } from 'react-native';
import { H2, View, Text } from 'native-base';
import { styles, ErrorView } from './common';

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
      this.setState({loading: true});
      let req = new Request("https://www.politivate.org/api/v1/profile",
          {headers: {"X-Auth-Token": this.props.appstate.authtoken}});
      let json = await (await fetch(req)).json();
      this.setState({
        loading: false,
        profile: json.resp,
      });
    } catch(error) {
      this.setState({
        loading: false,
        error: error,
      });
    }
  }

  renderLoaded() {
    return (
      <View>
        <Text>Id: {this.state.profile.id}</Text>
        <Text>Name: {this.state.profile.name}</Text>
      </View>
    );
  }

  render() {
    if (this.state.error) {
      return <ErrorView msg={this.state.error}/>;
    }
    return (
      <View>
        <View style={styles.tabheader}>
          <H2>Profile</H2>
        </View>
        <ScrollView refreshControl={
            <RefreshControl refreshing={this.state.loading}
                            onRefresh={this.update}/>}>
          { this.state.loading ? null : this.renderLoaded() }
        </ScrollView>
      </View>
    );
  }
}
