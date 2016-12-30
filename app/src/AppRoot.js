"use strict";

import React, { Component } from 'react';
import { Navigator, BackAndroid, AsyncStorage, Linking } from 'react-native';
import Tabs from './Tabs';
import LoginView from './LoginView';
import { LoadingView, ErrorView } from './common';

const REGISTERED_OTP_PREFIX = "politivate-org-app://www.politivate.org/api/v1/login/otp/";

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
      logged_in: false,
      error: null,
      token: null
    };
    this.renderScene = this.renderScene.bind(this);
    this.logout = this.logout.bind(this);
  }

  async componentDidMount() {
    try {
      let initialURL = await Linking.getInitialURL();
      if (initialURL && initialURL.startsWith(REGISTERED_OTP_PREFIX)) {
        let token = initialURL.slice(REGISTERED_OTP_PREFIX.length);
        if (token.indexOf("#") >= 0) {
          token = token.slice(0, token.indexOf("#"));
        }
        if (token.length > 0) {
          let req = new Request(
              "https://www.politivate.org/api/v1/login?otp=" + token);
          let resp = await fetch(req);
          if (resp.ok) {
            let json = await resp.json();
            let auth_token = json.resp.token;
            await AsyncStorage.setItem("@v1/auth/token", auth_token);
            this.setState({
              loading: false,
              logged_in: true,
              error: null,
              token: auth_token});
            return;
          }
        }
      }

      let auth_token = await AsyncStorage.getItem("@v1/auth/token");
      if (auth_token) {
        this.setState({
          loading: false,
          logged_in: true,
          error: null,
          token: auth_token});
        return;
      }

      this.setState({loading: false, logged_in: false});
    } catch(err) {
      this.setState({error: err, loading: false});
    }
  }

  renderScene(route, navigator) {
    let appstate = {
      logout: this.logout,
      authtoken: this.state.token};
    return (<BackHandler route={route} navigator={navigator}
                         appstate={appstate} />)
  }

  async logout() {
    try {
      this.setState({loading: true});
      await AsyncStorage.removeItem("@v1/auth/token");
      this.setState({
        loading: false,
        logged_in: false,
        error: null,
        token: null
        });
    } catch(err) {
      this.setState({loading: false, error: err});
    }
  }

  render() {
    if (this.state.loading) {
      return <LoadingView/>;
    }
    if (this.state.error) {
      return <ErrorView msg={this.state.error}/>;
    }
    if (!this.state.logged_in) {
      return <LoginView/>;
    }
    return (<Navigator initialRoute={{component: Tabs, _isRoot: true}}
                       renderScene={this.renderScene} />);
  }
}
