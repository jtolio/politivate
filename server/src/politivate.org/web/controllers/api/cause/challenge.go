package cause

import (
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/controllers/api/gov"
	"politivate.org/web/controllers/auth"
)

var (
	challengeId = whmux.NewIntArg()
)

func init() {
	mux["challenge"] = challengeId.Shift(whmux.Exact(
		http.HandlerFunc(serveChallenge)))
}

func serveChallenge(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	challenge := mustGetCause(ctx).GetChallenge(ctx, challengeId.MustGet(ctx))
	resp := map[string]interface{}{
		"challenge": challenge,
	}
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
		resp["legislators"] = legislators
	}

	whjson.Render(w, r, resp)
}
