package cause

import (
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/controllers/api/gov"
	"politivate.org/web/controllers/auth"
	"politivate.org/web/models"
)

var (
	challengeId = whmux.NewIntArg()
)

func init() {
	mux["challenge"] = challengeId.Shift(
		whmux.Dir{
			"": whmux.RequireGet(http.HandlerFunc(serveChallenge)),
			"complete": whmux.ExactPath(whmux.RequireMethod("POST",
				http.HandlerFunc(completeChallenge))),
		})
}

func legislators(ctx context.Context, u *models.User, chal *models.Challenge) []*gov.Legislator {
	rv := make([]*gov.Legislator, 0)
	for _, district := range gov.DistrictLocateByGPS(
		ctx, u.Latitude, u.Longitude) {
		switch chal.Data.Database {
		case "us", "ushouse":
			rv = append(rv, gov.HouseRepsByDistrict(ctx, district)...)
		}
		switch chal.Data.Database {
		case "us", "ussenate":
			rv = append(rv, gov.SenatorsByDistrict(ctx, district)...)
		}
	}
	return rv
}

func serveChallenge(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	challenge := mustGetCause(ctx).GetChallenge(ctx, challengeId.MustGet(ctx))
	m := challenge.JSON()
	u := auth.User(r)
	if challenge.Data.Database != "direct" {
		m["legislators"] = legislators(ctx, u, challenge)
	}
	m["actions"] = challenge.Completed(ctx, u)

	whjson.Render(w, r, m)
}

func completeChallenge(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(r)
	challenge := mustGetCause(ctx).GetChallenge(ctx, challengeId.MustGet(ctx))

	now := time.Now()
	if !challenge.Info.EventStart.Time.IsZero() &&
		now.Before(challenge.Info.EventStart.Time) {
		whfatal.Error(wherr.BadRequest.New("event hasn't started"))
	}
	if !challenge.Info.EventEnd.Time.IsZero() &&
		now.After(challenge.Info.EventEnd.Time) {
		whfatal.Error(wherr.BadRequest.New("event is over"))
	}

	completed := challenge.Completed(ctx, u)

AddChallengeSwitch:
	switch challenge.Info.Type {
	case "location":
		if len(completed) > 0 {
			break AddChallengeSwitch
		}

		if challenge.Data.Database != "direct" {
			whfatal.Error(wherr.InternalServerError.New(
				"unknown location challenge database type"))
		}

		vals := make(map[string]float64)
		for _, name := range []string{"challenge_latitude", "challenge_longitude",
			"user_latitude", "user_longitude"} {
			if r.FormValue(name) == "" {
				whfatal.Error(wherr.BadRequest.New("missing field: %s", name))
			}
			val, err := strconv.ParseFloat(r.FormValue(name), 64)
			if err != nil {
				whfatal.Error(wherr.BadRequest.New("bad field %s: %s", name, err))
			}
			vals[name] = val
		}

		if challenge.Data.DirectLatitude != vals["challenge_latitude"] ||
			challenge.Data.DirectLongitude != vals["challenge_longitude"] {
			whfatal.Error(wherr.BadRequest.New(
				"challenge latitude/longitude mismatch"))
		}

		if challenge.Data.DirectRadius < Haversine(
			challenge.Data.DirectLatitude, challenge.Data.DirectLongitude,
			vals["user_latitude"], vals["user_longitude"]) {
			whfatal.Error(wherr.BadRequest.New("out of range"))
		}

		action := challenge.Action(ctx, u)
		action.Latitude = challenge.Data.DirectLatitude
		action.Longitude = challenge.Data.DirectLongitude
		action.Save(ctx)
		completed = append(completed, action)

	case "phonecall":
		phoneNumber := r.FormValue("phone_number")
		for _, action := range completed {
			if action.Phone == phoneNumber {
				break AddChallengeSwitch
			}
		}

		switch challenge.Data.Database {
		case "direct":
			if challenge.Data.DirectPhone != phoneNumber {
				whfatal.Error(wherr.BadRequest.New("invalid phone number"))
			}

		case "ushouse", "ussenate", "us":
			found := false
			for _, legislator := range legislators(ctx, u, challenge) {
				if phoneNumber == legislator.Phone {
					found = true
					break
				}
			}
			if !found {
				whfatal.Error(wherr.BadRequest.New("invalid phone number"))
			}

		default:
			whfatal.Error(wherr.InternalServerError.New(
				"unknown challenge database type"))
		}

		action := challenge.Action(ctx, u)
		action.Phone = phoneNumber
		action.Save(ctx)
		completed = append(completed, action)
	default:
		whfatal.Error(wherr.InternalServerError.New("unknown challenge type"))
	}

	whjson.Render(w, r, map[string]interface{}{
		"actions": completed,
	})
}
