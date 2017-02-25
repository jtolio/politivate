"use strict";

import React, { Component } from 'react';
import {
  View, Text, Button, Image, TouchableOpacity, NativeModules, Linking
} from 'react-native';
import Icon from 'react-native-vector-icons/Entypo';

class Color {
  constructor(hue, saturation, lightness) {
    this.hue = hue;
    this.saturation = saturation;
    this.lightness = lightness;
  }

  alpha(amount) {
    return "hsla(" + this.hue + ", " +
        (100 * this.saturation) + "%, " +
        (100 * this.lightness) + "%, " +
        amount + ")";
  }

  get val() {
    return this.alpha(1.0);
  }

  get faint() {
    return this.alpha(0.1);
  }
}

async function phonecall(number) {
  if (NativeModules.CallAndroid) {
    NativeModules.CallAndroid.call(number);
    return true;
  }
  // TODO: this is all just a guess about iOS. actually test it.
  let url = "telprompt:" + number;
  if (!await Linking.canOpenURL(url)) {
    url = "tel:" + number;
  }
  if (await Linking.canOpenURL(url)) {
    return false;
  }
  try {
    await Linking.openURL(url)
  } catch(err) {
    if (url.startsWith("telprompt:")) {
      return false;
    }
    throw err;
  }
  return true;
}

var palette = {
  red: new Color(5, 1.0, 0.594),
  white: new Color(0, 0.0, 1.0),
  blue: new Color(212, 1.0, 0.5),
};

var colors = {
  link: palette.blue,
  heart: palette.red,
  background: palette.white,
  primary: palette.blue,
  secondary: palette.red,
}

class TextLoadingView extends Component {
  constructor(props) {
    super(props);
    this.state = {counter: 0};
    this.tick = this.tick.bind(this);
  }

  componentDidMount() {
    this.timer = setInterval(this.tick, 300);
  }

  componentWillUnmount() {
    clearInterval(this.timer);
  }

  tick() {
    this.setState((prevState, props) => {
      return {counter: prevState.counter + 1};
    });
  }

  render() {
    let dots = "";
    for (let i = 0; i < this.state.counter % 4; i++) {
      dots += ".";
    }
    for (let i = dots.length; i < 4; i++) {
      dots += " ";
    }
    return (
      <View style={{
          flex: 1,
          alignItems: "center",
          justifyContent: "center"}}>
        <Text>Loading{dots}</Text>
      </View>
    );
  }
}

class LoadingView extends Component {
  render() {
    return (
      <View style={{
          flex: 1,
          alignItems: "center",
          justifyContent: "center"}}>
        <Icon name="hour-glass" size={30} />
      </View>
    );
  }
}

class ErrorView extends Component {
  render() {
    return <View><Text>Error: {this.props.msg.toString()}</Text></View>;
  }
}

class Link extends Component {
  render() {
    return (
      <TouchableOpacity onPress={() =>
          Linking.openURL(this.props.url).catch(err => {})}>
        <Text style={{color: colors.link.val}}>
          {this.props.children}
        </Text>
      </TouchableOpacity>
    );
  }
}

class TabHeader extends Component {
  render() {
    return (
      <View>
        <Text style={{
          padding: 10,
          fontWeight: "bold",
          color: colors.primary.val,
          fontSize: 20
        }}>
          {this.props.children}
        </Text>
      </View>);
  }
}

class AppHeader extends Component {
  render() {
    return (
      <View style={{flexDirection: "row", padding: 10}}>
        <Image source={require("../images/header.png")} style={{
          resizeMode: "contain",
          flex: 1,
          width: null,
          height: 30,
          borderWidth: 0,
        }}/>
      </View>);
  }
}

class Separator extends Component {
  render() {
    return (
      <View style={{borderBottomWidth: 1,
                    borderColor: colors.primary.val}}/>);
  }
}

class StyledButton extends Component {
  render() {
    return (
      <Button onPress={this.props.onPress} title={this.props.title}
              color={colors.primary.val} />
    );
  }
}

module.exports = {
  AppHeader, LoadingView, ErrorView, Link, TabHeader, colors, Separator,
  phonecall,
  Button: StyledButton
}
