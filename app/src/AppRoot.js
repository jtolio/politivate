"use strict";

import React, { Component } from 'react';
import {
  Navigator, BackAndroid, AsyncStorage, Linking, View, NativeModules
} from 'react-native';
import Tabs from './Tabs';
import Resources from './Resources';
import LoginView from './LoginView';
import { LoadingView, ErrorView, colors } from './common';

const REGISTERED_OTP_PREFIX = "politivate-org-app://www.politivate.org/api/v1/login/otp/";
const DEMO_MODE_TOKEN = "";
//const DEMO_MODE_TOKEN = "d52e33faf8695d6c2d7dab00db0a64c502ea7ceb4329272752317633e35ca040";

class BackHandler extends Component {
  constructor(props) {
    super(props);
    this.backPress = this.backPress.bind(this);
  }

  backPress() {
    if (this.props.route.hasOwnProperty("_isRoot")) {
      return false;
    }
    this.props.appstate.navigator.pop();
    return true;
  }

  componentDidMount() {
    BackAndroid.addEventListener('hardwareBackPress', this.backPress);
  }

  componentWillUnmount() {
    BackAndroid.removeEventListener('hardwareBackPress', this.backPress);
  }

  render() {
    return (
      <View style={{backgroundColor: colors.background.val, flex: 1}}>
        <this.props.route.component
            appstate={{backPress: this.backPress, ...this.props.appstate}}
            {...this.props.route.passProps}/>
      </View>
    );
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

  async otpLogin(url) {
    if (!url || !url.startsWith(REGISTERED_OTP_PREFIX)) {
      return false;
    }
    let otp = url.slice(REGISTERED_OTP_PREFIX.length);
    let fragmentIndex = otp.indexOf("#");
    if (fragmentIndex >= 0) {
      otp = otp.slice(0, fragmentIndex);
    }
    if (otp.length == 0) {
      return false;
    }
    let last_otp = await AsyncStorage.getItem("@v1/auth/last_otp");
    if (last_otp == otp) {
      return false;
    }
    let req = new Request(
        "https://www.politivate.org/api/v1/login?otp=" + otp);
    let resp = await fetch(req);
    if (!resp.ok) {
      return false;
    }
    await AsyncStorage.setItem("@v1/auth/last_otp", otp);
    let json = await resp.json();
    let auth_token = json.resp.token;
    await AsyncStorage.setItem("@v1/auth/token", auth_token);
    this.setState({
      loading: false,
      logged_in: true,
      error: null,
      token: auth_token});
    return true;
  }

  async componentDidMount() {
    Linking.addEventListener('url', event => {
      this.otpLogin(event.url)
    })

    try {
      if (await this.otpLogin(await Linking.getInitialURL())) {
        return;
      }

      let auth_token = DEMO_MODE_TOKEN.length > 0 ? DEMO_MODE_TOKEN : (
            await AsyncStorage.getItem("@v1/auth/token"));
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
      resources: new Resources(this.state.token, this.logout),
      navigator: navigator,
    };

    return (<BackHandler route={route} appstate={appstate} />)
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
