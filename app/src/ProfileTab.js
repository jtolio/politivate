"use strict";

import React, { Component } from 'react';
import {
  ScrollView, RefreshControl, View, Text, Image, TouchableOpacity
} from 'react-native';
import { ErrorView, TabHeader, Link, colors } from './common';

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
            <Text>Profile text!</Text>
          </View>
        </View>
        <View style={{paddingTop: 20}}/>
        <View style={{borderWidth: 1, borderColor: colors.primary.val, borderRadius: 10, padding: 10, paddingTop: 6}}>
          <Text style={{fontWeight: "bold", fontSize: 30, color: colors.primary.val}}>Points this month</Text>
          <View style={{borderBottomWidth: 1, borderColor: colors.primary.val}}/>
          <View style={{justifyContent: "space-between", flexDirection: "row"}}>
            <Text style={{fontWeight: "bold", fontSize: 20}}>Cause 1</Text>
            <Text style={{fontWeight: "bold", fontSize: 20}}>70</Text>
          </View>
          <View style={{borderBottomWidth: 1, borderColor: colors.primary.val}}/>
          <View style={{justifyContent: "space-between", flexDirection: "row"}}>
            <Text style={{fontWeight: "bold", fontSize: 20}}>Cause 2</Text>
            <Text style={{fontWeight: "bold", fontSize: 20}}>110</Text>
          </View>
        </View>
        <View style={{paddingTop: 20}}/>
        <View style={{borderWidth: 1, borderColor: colors.primary.val, borderRadius: 10, padding: 10, paddingTop: 6}}>
          <Text style={{fontWeight: "bold", fontSize: 30, color: colors.primary.val}}>Achievements</Text>
          <View style={{borderBottomWidth: 1, borderColor: colors.primary.val}}/>
          <View style={{justifyContent: "space-between", flexDirection: "row"}}>
            <Text style={{fontWeight: "bold", fontSize: 20}}>Longest streak</Text>
            <Text style={{fontWeight: "bold", fontSize: 20}}>8 days</Text>
          </View>
          <View style={{borderBottomWidth: 1, borderColor: colors.primary.val}}/>
          <View style={{justifyContent: "space-between", flexDirection: "row"}}>
            <Text style={{fontWeight: "bold", fontSize: 20}}>Phone calls</Text>
            <Text style={{fontWeight: "bold", fontSize: 20}}>18</Text>
          </View>
          <View style={{borderBottomWidth: 1, borderColor: colors.primary.val}}/>
          <View style={{justifyContent: "space-between", flexDirection: "row"}}>
            <Text style={{fontWeight: "bold", fontSize: 20}}>Top supporter</Text>
            <Text/>
          </View>
          <View style={{justifyContent: "space-between", flexDirection: "row"}}>
            <Text/>
            <Text style={{fontWeight: "bold", fontSize: 20}}>Cause 1, City, Month 1</Text>
          </View>
          <View style={{justifyContent: "space-between", flexDirection: "row"}}>
            <Text/>
            <Text style={{fontWeight: "bold", fontSize: 20}}>Cause 2, City, Month 2</Text>
          </View>
          <View style={{borderBottomWidth: 1, borderColor: colors.primary.val}}/>
          <View style={{justifyContent: "space-between", flexDirection: "row"}}>
            <Text style={{fontWeight: "bold", fontSize: 20}}>Rallies attended</Text>
            <Text style={{fontWeight: "bold", fontSize: 20}}>5</Text>
          </View>
          <View style={{borderBottomWidth: 1, borderColor: colors.primary.val}}/>
          <View style={{justifyContent: "space-between", flexDirection: "row"}}>
            <Text style={{fontWeight: "bold", fontSize: 20}}>Active days</Text>
            <Text style={{fontWeight: "bold", fontSize: 20}}>37</Text>
          </View>
        </View>
        <View style={{paddingTop: 20}}/>
        <View style={{borderWidth: 1, borderColor: colors.primary.val, borderRadius: 10, padding: 10, paddingTop: 6}}>
          <Text style={{fontWeight: "bold", fontSize: 30, color: colors.primary.val}}>Total Points</Text>
          <View style={{borderBottomWidth: 1, borderColor: colors.primary.val}}/>
          <View style={{justifyContent: "space-between", flexDirection: "row"}}>
            <Text style={{fontWeight: "bold", fontSize: 20}}>Cause 1</Text>
            <Text style={{fontWeight: "bold", fontSize: 20}}>350</Text>
          </View>
          <View style={{borderBottomWidth: 1, borderColor: colors.primary.val}}/>
          <View style={{justifyContent: "space-between", flexDirection: "row"}}>
            <Text style={{fontWeight: "bold", fontSize: 20}}>Cause 2</Text>
            <Text style={{fontWeight: "bold", fontSize: 20}}>620</Text>
          </View>
          <View style={{borderBottomWidth: 1, borderColor: colors.primary.val}}/>
          <View style={{justifyContent: "space-between", flexDirection: "row"}}>
            <Text style={{fontWeight: "bold", fontSize: 20}}>Cause 3</Text>
            <Text style={{fontWeight: "bold", fontSize: 20}}>70</Text>
          </View>
        </View>
        <View style={{paddingTop: 20}}/>
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
