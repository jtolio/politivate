package cause

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/models"
	"politivate.org/web/views"
)

func init() {
	mux["challenges"] = whmux.Dir{
		"new": whmux.ExactPath(whmux.Method{
			"GET":  http.HandlerFunc(newChallengeForm),
			"POST": http.HandlerFunc(newChallengeCreate),
		}),
	}
}

func newChallengeForm(w http.ResponseWriter, r *http.Request) {
	views.Render(w, r, "new_challenge", map[string]interface{}{
		"Cause": administerCause(r),
		"Form":  map[string]interface{}{},
	})
}

func newChallengeCreate(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	c := administerCause(r)

	chal := c.NewChallenge(ctx)
	chal.Info.Title = r.FormValue("title")
	chal.Data.Description = r.FormValue("description")
	points, err := strconv.Atoi(r.FormValue("points"))
	if err != nil {
		points = -1
	}
	chal.Info.Points = points
	chal.Info.Type = r.FormValue("type")
	switch chal.Info.Type {
	case "phonecall":
		chal.Data.Database = r.FormValue("phoneDatabase")
		chal.Data.DirectPhone = r.FormValue("directphone")
	case "location":
		chal.Data.Database = r.FormValue("locationDatabase")
		chal.Data.DirectAddress = r.FormValue("directaddr")
	default:
		whfatal.Error(wherr.BadRequest.New("bad challenge type: %s",
			chal.Info.Type))
	}
	switch chal.Data.Database {
	default:
		whfatal.Error(wherr.BadRequest.New("bad database type: %s",
			chal.Data.Database))
	case "direct", "us", "ushouse", "ussenate":
	}
	restrictions, err := strconv.Atoi(r.FormValue("restrictionLength"))
	if err != nil {
		whfatal.Error(wherr.BadRequest.Wrap(err))
	}
	chal.Info.Restrictions = make([]models.ChallengeRestriction, 0, restrictions)
	for i := 0; i < restrictions; i++ {
		chal.Info.Restrictions = append(chal.Info.Restrictions,
			models.ChallengeRestriction{
				Type:  r.FormValue(fmt.Sprintf("restrictionType[%d]", i)),
				Value: r.FormValue(fmt.Sprintf("restrictionValue[%d]", i))})
	}

	switch r.FormValue("dateType") {
	case "none":
	case "event":
		tz, err := time.Parse("2006-01-02T15:04 MST",
			r.FormValue("eventStart")+" EST")
		if err != nil {
			whfatal.Error(wherr.BadRequest.Wrap(err))
		}
		chal.Info.EventStart = models.NullableTime{Time: tz}
		fallthrough
	case "deadline":
		tz, err := time.Parse("2006-01-02T15:04 MST",
			r.FormValue("eventEnd")+" EST")
		if err != nil {
			whfatal.Error(wherr.BadRequest.Wrap(err))
		}
		chal.Info.EventEnd = models.NullableTime{Time: tz}
	default:
		whfatal.Error(wherr.BadRequest.New("bad date type: %s",
			r.FormValue("dateType")))
	}

	if chal.Info.Title == "" || chal.Data.Description == "" ||
		chal.Info.Points < 0 || chal.Data.Database == "" ||
		(chal.Info.Type == "phonecall" && chal.Data.Database == "direct" &&
			chal.Data.DirectPhone == "") ||
		(chal.Info.Type == "location" && chal.Data.Database == "direct" &&
			chal.Data.DirectAddress == "") {
		formVals := map[string]string{}
		for _, name := range []string{"title", "description", "points", "type",
			"phoneDatabase", "locationDatabase", "directphone", "directaddr",
			"dateType", "eventStart", "eventEnd"} {
			formVals[name] = r.FormValue(name)
		}
		views.Render(w, r, "new_challenge", map[string]interface{}{
			"Cause":        c,
			"Error":        "Required field missing",
			"Form":         formVals,
			"Restrictions": chal.Info.Restrictions,
		})
		return
	}

	chal.Save(ctx)
	whfatal.Redirect(fmt.Sprintf("/cause/%d", c.Id))
}
