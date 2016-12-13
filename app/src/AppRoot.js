"use strict";

import React, { Component } from 'react';
import { Navigator, BackAndroid, AsyncStorage } from 'react-native';
import Tabs from './Tabs';
import { auth, LoadingView } from './common';

class BackHandler extends Component {
  constructor(props) {
    super(props);
    this.backPress = this.backPress.bind(this);
  }

  backPress() {
    if (this.props.route.hasOwnProperty("_isRoot")) {
      return false;
    }
    this.props.navigator.pop();
    return true;
  }

  componentDidMount() {
    BackAndroid.addEventListener('hardwareBackPress', this.backPress);
  }

  componentWillUnmount() {
    BackAndroid.removeEventListener('hardwareBackPress', this.backPress);
  }

  render() {
    return (<this.props.route.component navigator={this.props.navigator}
               appstate={this.props.appstate} backPress={this.backPress}
               {...this.props.route.passProps}/>);
  }
}

export default class AppRoot extends Component {
  constructor(props) {
    super(props)
    this.state = {
      loading: true,
      token: null
    };
    this.renderScene = this.renderScene.bind(this);
    AsyncStorage.getItem("@v1/auth/token")
      .then((result) => {
        if (result !== null) {
          console.log("already logged in!");
          this.setState({
            loading: false,
            token: result});
          return;
        }
        console.log("logging in...");
        auth.authorize('google')
          .then(resp => {
              if (resp.status !== "ok") {
                console.error(resp.status);
                return;
              }
              console.log("success logging in!");
              let token = resp.response.credentials.oauth_token;
              AsyncStorage.setItem("@v1/auth/token", token)
                .then(() => {
                    this.setState({
                      loading: false,
                      token: token});
                  })
                .catch(err => console.log(err));
            })
          .catch(err => console.log(err));
        })
      .catch((err) => console.error(err));
  }

  renderScene(route, navigator) {
    return (<BackHandler route={route} navigator={navigator}
                         appstate={{token: this.state.token}} />)
  }

  render() {
    if (this.state.loading) {
      return <LoadingView/>;
    }
    return (<Navigator initialRoute={{component: Tabs, _isRoot: true}}
                       renderScene={this.renderScene} />);
  }
}
