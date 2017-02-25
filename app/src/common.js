"use strict";

import React, { Component } from 'react';
import {
  View, Text, Button, Image, TouchableOpacity, NativeModules, Linking, Alert
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

// confirm returns a promise that says true or false if the user clicks okay
// or cancel, given the alert message.
function confirm(title, message) {
  return new Promise(function(resolve, reject) {
    Alert.alert(title, message, [
        {text: "Cancel", onPress: () => { resolve(false); }, style: "cancel"},
        {text: "OK", onPress: () => { resolve(true); }},
      ],
      {cancelable: false});
  });
}

// phonecall will prompt the user if they want to make a call to 'number' and
// return true if the user did follow through and initiate the call. 'who' is
// used to show messages to the user about who is getting called. after the
// user hangs up, focus should be given back to the app (TODO: test ios)
//
// background: we want to give users points for calling, but we don't want to
// give users points if they cheat and don't call. the best case scenario would
// be to be able to look at the call log after initiating the call and confirm
// that the user spent x minutes in the call. unfortunately, this is hard on
// android and impossible on ios.
//
// instead, we're going to give users credit for calling if they actually
// follow through and really make the phone call. to cut down on false
// positives, we're going to give the user as many outs as possible to
// understand they're about to make a call and can cancel. ultimately, we need
// to know if the user goes forward and makes the call.
//
// on android, we're going to pop up a dialog and if the user okays this, then
// we'll start the call for them. we're not going to use the system dial dialog
// because then we don't know if the call actually got started.
//
// on ios, i don't know exactly what we're going to do yet but from what i've
// read it looks like using the undocumented "telprompt:" scheme will not only
// ask the user if they want to call, but throw an error in react native if
// they decide not to call, returning to the app. So in ios, I don't think we
// need the popup.
// TODO: test ios.
//
// either way, this function returns a promise that evaluates to true or false
// on our best estimate of if the user actually really started a phone call.
// in the future, if we can be more accurate and keep track of timing, etc,
// this is the function to put that logic in.
async function phonecall(who, number) {
  if (NativeModules.CallAndroid) {
    let msg = "Call " + who + " (" + number + ")?";
    if (!await confirm("Are you sure?", msg)) {
      return false;
    }
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
