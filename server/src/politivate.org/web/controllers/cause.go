package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whmux"

	"politivate.org/web/controllers/auth"
	"politivate.org/web/models"
)

var (
	causeId = whmux.NewIntArg()
)

func init() {
	mux["cause"] = causeId.Shift(whmux.Dir{
		"": whmux.Exact(http.HandlerFunc(cause)),
		"challenges": whmux.Dir{
			"new": whmux.Method{
				"GET":  http.HandlerFunc(newChallengeForm),
				"POST": http.HandlerFunc(newChallengeCreate),
			},
		},
	})
}

func cause(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := auth.User(r)
	c := models.GetCause(ctx, causeId.MustGet(ctx))
	isAdministrating := false
	if u != nil {
		isAdministrating = u.IsAdministrating(ctx, c)
	}
	Render(w, r, "cause", map[string]interface{}{
		"IsAdministrating": isAdministrating,
		"Cause":            c,
	})
}

func administerCause(r *http.Request) *models.Cause {
	ctx := whcompat.Context(r)
	u := auth.User(r)
	c := models.GetCause(ctx, causeId.MustGet(ctx))
	if u == nil || !u.IsAdministrating(ctx, c) {
		whfatal.Redirect(fmt.Sprintf("/cause/%d", c.Id))
		return nil
	}
	return c
}

func newChallengeForm(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "new_challenge", map[string]interface{}{
		"Cause": administerCause(r),
		"Form":  map[string]interface{}{},
	})
}

func newChallengeCreate(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	c := administerCause(r)

	chal := c.NewChallenge(ctx)
	chal.Title = r.FormValue("title")
	chal.Description = r.FormValue("description")
	points, err := strconv.Atoi(r.FormValue("points"))
	if err != nil {
		whfatal.Error(wherr.BadRequest.Wrap(err))
	}
	chal.Points = points
	chal.Type = r.FormValue("type")
	switch chal.Type {
	case "phonecall":
		chal.Database = r.FormValue("phoneDatabase")
		chal.DirectPhone = r.FormValue("directphone")
	case "location":
		chal.Database = r.FormValue("locationDatabase")
		chal.DirectAddress = r.FormValue("directaddr")
	default:
		whfatal.Error(wherr.BadRequest.New("bad challenge type: %s", chal.Type))
	}
	restrictions, err := strconv.Atoi(r.FormValue("restrictionLength"))
	if err != nil {
		whfatal.Error(wherr.BadRequest.Wrap(err))
	}
	chal.Restrictions = make([]models.ChallengeRestriction, 0, restrictions)
	for i := 0; i < restrictions; i++ {
		chal.Restrictions = append(chal.Restrictions,
			models.ChallengeRestriction{
				Type:  r.FormValue(fmt.Sprintf("restrictionType[%d]", i)),
				Value: r.FormValue(fmt.Sprintf("restrictionValue[%d]", i))})
	}

	if r.FormValue("deadlineEnabled") != "" {
		tz, err := time.Parse("2006-01-02", r.FormValue("deadline"))
		if err != nil {
			whfatal.Error(wherr.BadRequest.Wrap(err))
		}
		chal.Deadline = tz
	}

	if r.FormValue("startdateEnabled") != "" {
		tz, err := time.Parse("2006-01-02", r.FormValue("startdate"))
		if err != nil {
			whfatal.Error(wherr.BadRequest.Wrap(err))
		}
		chal.Start = tz
	}

	if chal.Title == "" || chal.Description == "" || chal.Points < 0 || chal.Database == "" ||
		(chal.Type == "phonecall" && chal.Database == "direct" && chal.DirectPhone == "") ||
		(chal.Type == "location" && chal.Database == "direct" && chal.DirectAddress == "") {
		formVals := map[string]string{}
		for _, name := range []string{"title", "description", "points", "type",
			"phoneDatabase", "locationDatabase", "directphone", "directaddr",
			"deadlineEnabled", "deadline", "startdateEnabled", "startdate"} {
			formVals[name] = r.FormValue(name)
		}
		Render(w, r, "new_challenge", map[string]interface{}{
			"Cause":        c,
			"Error":        "Required field missing",
			"Form":         formVals,
			"Restrictions": chal.Restrictions,
		})
		return
	}

	chal.Save(ctx)
	whfatal.Redirect(fmt.Sprintf("/cause/%d", c.Id))
}
