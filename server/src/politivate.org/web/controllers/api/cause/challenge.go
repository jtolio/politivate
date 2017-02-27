package cause

import (
	"net/http"
	"strconv"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/controllers/api/gov"
	"politivate.org/web/controllers/auth"
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

func serveChallenge(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	challenge := mustGetCause(ctx).GetChallenge(ctx, challengeId.MustGet(ctx))
	m := challenge.JSON()
	if challenge.Data.Database != "direct" {
		u := auth.User(r)
		legislators := make([]*gov.Legislator, 0)
		for _, district := range gov.DistrictLocateByGPS(
			ctx, u.Latitude, u.Longitude) {
			switch challenge.Data.Database {
			case "us", "ushouse":
				legislators = append(legislators,
					gov.HouseRepsByDistrict(ctx, district)...)
			}
			switch challenge.Data.Database {
			case "us", "ussenate":
				legislators = append(legislators,
					gov.SenatorsByDistrict(ctx, district)...)
			}
		}
		m["legislators"] = legislators
	}

	whjson.Render(w, r, m)
}

func completeChallenge(w http.ResponseWriter, r *http.Request) {
	latitude, err := strconv.ParseFloat(r.FormValue("latitude"), 64)
	if err != nil {
		whfatal.Error(wherr.BadRequest.Wrap(err))
	}
	longitude, err := strconv.ParseFloat(r.FormValue("longitude"), 64)
	if err != nil {
		whfatal.Error(wherr.BadRequest.Wrap(err))
	}
	whjson.Render(w, r, map[string]interface{}{
		"lat": latitude,
		"lon": longitude,
	})
}
