/* @flow */
"use strict";

import React from 'react';
import { Text, Image, View, ScrollView, Button } from 'react-native';
import Subpage from './Subpage';
import LoadablePage from './LoadablePage';
import { Link, ErrorView, LoadingView, phonecall } from './common';
import FollowButton from './FollowButton';
import MapView from 'react-native-maps';

const EARTH_RADIUS = 6378137;

function square(val) { return val * val; }

// haversine returns distance between two gps coordinates in
// meters, assuming a spherical earth.
function haversine(lat1, lon1, lat2, lon2) {
  let latr1 = lat1 * Math.PI / 180;
  let latr2 = lat2 * Math.PI / 180;
  let dlatr = (lat2 - lat1) * Math.PI / 180;
  let dlonr = (lon2 - lon1) * Math.PI / 180;

  let a = square(Math.sin(dlatr/2)) +
          Math.cos(latr1) * Math.cos(latr2) *
          square(Math.sin(dlonr/2));
  return EARTH_RADIUS * 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1-a));
}

function euclidean(lat1, lon1, lat2, lon2) {
  let latr1 = lat1 * Math.PI / 180;
  let latr2 = lat2 * Math.PI / 180;
  let lonr1 = lon1 * Math.PI / 180;
  let lonr2 = lon2 * Math.PI / 180;
  var x = (lonr2 - lonr1) * Math.cos((latr1 + latr2)/2);
  var y = (latr2 - latr1);
  return Math.sqrt(square(x) + square(y)) * EARTH_RADIUS;
}

class ChallengeLocationMap extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      position: null,
      error: null,
      success_calls: 0,
      failure_calls: 0,
    };
    this.watch_id = null;
    this.updateSuccess = this.updateSuccess.bind(this);
    this.updateFailure = this.updateFailure.bind(this);
  }

  componentDidMount() {
    if (this.watch_id !== null) { return; }
    this.watch_id = navigator.geolocation.watchPosition(
        this.updateSuccess, this.updateFailure, {
      enableHighAccuracy: true,
      maximumAge: 0,
      distanceFilter: 1,
    });
  }

  componentWillUnmount() {
    if (this.watch_id !== null) {
      navigator.geolocation.clearWatch(this.watch_id);
      this.watch_id = null;
    }
  }

  updateSuccess(position) {
    if (this.watch_id === null) { return; }
    // makes debugging harder but also cheating
    if (position.mocked) { return; }
    this.setState((state) => ({
      success_calls: state.success_calls+1,
      position,
    }));
  }

  updateFailure(error) {
    if (this.watch_id === null) { return; }
    this.setState((state) => ({
      failure_calls: state.failure_calls+1,
      error,
    }));
  }

  render() {
    let chal = this.props.challenge;
    if (chal.database != "direct") {
      return <ErrorView msg="Unknown challenge type!"/>;
    }
    let pos = this.state.position;
    if (!pos) {
      return <LoadingView/>;
    }
    let distance = haversine(
        pos.coords.latitude, pos.coords.longitude,
        chal.direct_latitude, chal.direct_longitude);
    let in_range = distance <= chal.direct_radius;

    let latitudeDelta = Math.abs(chal.direct_latitude-pos.coords.latitude);
    let longitudeDelta = Math.abs(chal.direct_longitude-pos.coords.longitude);
    let initialRegion = {
      latitude: (chal.direct_latitude + pos.coords.latitude)/2 +
                latitudeDelta * 0.15,
      longitude: (chal.direct_longitude + pos.coords.longitude)/2,
      latitudeDelta: latitudeDelta * 1.5,
      longitudeDelta: longitudeDelta * 1.5,
    };


    let now = Date.now();
    let after_event_start = !(chal.event_start && now < chal.event_start);
    let before_event_end = !(chal.event_end && now > chal.event_end);

    var button;
    if (after_event_start) {
      if (before_event_end) {
        if (in_range) {
          button = <Button title="Check in" onPress={() => {}}/>;
        } else {
          button = <Button disabled title="Not in range" onPress={() => {}}/>;
        }
      } else {
        button = <Button disabled title="Event is over" onPress={() => {}}/>;
      }
    } else {
      button = <Button disabled title="Event hasn't started" onPress={() => {}}/>;
    }

    return (
      <View>
        <MapView style={{height: 200}} showUserLocation={false}
            rotateEnabled={false} pitchEnabled={false} loadingEnabled={true}
            initialRegion={initialRegion}>
          <MapView.Marker coordinate={{
              latitude: chal.direct_latitude,
              longitude: chal.direct_longitude,
            }}/>
          <MapView.Marker coordinate={{
              latitude: pos.coords.latitude,
              longitude: pos.coords.longitude,
            }} image={require("../images/person.png")}/>
          <MapView.Circle radius={chal.direct_radius} center={{
              latitude: chal.direct_latitude,
              longitude: chal.direct_longitude,
            }} strokeWidth={0} fillColor="#ff000033" geodesic={true}/>
        </MapView>
        <View style={{paddingTop: 10}}/>
        {button}
      </View>
    );
  }
}

class ChallengeLocationAction extends React.Component {
  render() {
    let chal = this.props.challenge;
    if (chal.database != "direct") {
      return <ErrorView msg="Unknown challenge type!"/>;
    }
    return (
      <View>
        { chal.event_start ? (
          <View style={{flexDirection: "row"}}>
            <Text style={{fontWeight: "bold", paddingRight: 10, width: 100}}>Start:</Text>
            <Text>{(new Date(chal.event_start)).toLocaleString()}</Text>
          </View>
        ) : null }
        { chal.event_end ? (
          <View style={{flexDirection: "row"}}>
            <Text style={{fontWeight: "bold", paddingRight: 10, width: 100}}>End:</Text>
            <Text>{(new Date(chal.event_end)).toLocaleString()}</Text>
          </View>
        ) : null }
        <View style={{flexDirection: "row", paddingBottom: 20}}>
          <Text style={{fontWeight: "bold", paddingRight: 10, width: 100}}>Address:</Text>
          <Text>{chal.direct_address.replace(", ", "\n").replace(", ", "\n")}</Text>
        </View>
        <ChallengeLocationMap challenge={chal}/>
      </View>
    );
  }
}

class ChallengePhonecallAction extends React.Component {
  render() {
    let who = this.props.who;
    let phone = this.props.phone;
    let message = "Call " + phone;
    if (who) {
      message = who + ": " + message;
    }
    return (
      <Button title={message}
          onPress={() => { phonecall(who, phone) }}/>
    );
  }
}

class ChallengePhonecallActions extends React.Component {
  render() {
    let chal = this.props.challenge;
    let results = [];
    if (chal.database == "direct") {
      results.push(
        <View style={{paddingTop: 10}} key="view-direct"/>
      );
      results.push(
        <ChallengePhonecallAction phone={chal.direct_phone}
            key="button-direct"/>
      );
    } else {
      for (var legislator of chal.legislators) {
        results.push(
          <View style={{paddingTop: 10}}
                key={"view-" + legislator.votesmart_id}/>
        );
        let title = {"senate": "Sen.", "house": "Rep."}[legislator.chamber];
        let name = legislator.first_name + " " + legislator.last_name;
        results.push(
          <ChallengePhonecallAction key={"button-" + legislator.votesmart_id}
                who={title + " " + name} phone={legislator.phone}/>
        );
      }
    }

    return (
      <View style={{paddingTop: 10}}>
        {results}
      </View>
    );
  }
}


export default class ChallengePage extends React.Component {
  resourceURL() {
    return "/v1/cause/" + this.props.challenge.cause_id +
           "/challenge/" + this.props.challenge.id;
  }

  renderLoaded(chal) {
    return (
      <View style={{
          padding: 20
        }}>
        <View style={{
            flexDirection: "row",
            alignItems: "center",
            paddingBottom: 10}}>
          <Image
            source={{uri: this.props.cause.icon_url}}
            style={{width: 50, height: 50, borderRadius: 10}}/>
          <View style={{paddingLeft: 10, flex: 1}}>
            <Text style={{fontWeight: "bold"}}>{this.props.cause.name}</Text>
            <Link url={this.props.cause.url}>
              {this.props.cause.url}
            </Link>
          </View>
          <FollowButton cause={this.props.cause}
                appstate={this.props.appstate} />
        </View>
        <Text>{chal.description}</Text>
        { chal.type == "phonecall" ?
          <ChallengePhonecallActions challenge={chal}/> :
          chal.type == "location" ?
          <ChallengeLocationAction challenge={chal}/> :
          null }
      </View>
    );
  }

  render() {
    return (
      <Subpage appstate={this.props.appstate}
               title={this.props.challenge.title}>
        <LoadablePage renderLoaded={this.renderLoaded.bind(this)}
                      resourceURL={this.resourceURL()}
                      appstate={this.props.appstate}/>
      </Subpage>
    );
  }
}
