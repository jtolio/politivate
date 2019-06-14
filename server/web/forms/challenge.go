package forms

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"

	"politivate.org/web/models"
)

func NewChallengeForm() *Form {
	return &Form{
		Template: "challenge_form",
		Form:     map[string]interface{}{},
	}
}

// TODO
var timeZone = time.FixedZone("EST", -18000)

func formatTime(t models.Time) string {
	return strings.TrimSuffix(t.Time.In(timeZone).
		Format("2006-01-02T15:04 MST"), " EST")
}

func parseTime(t string) (models.Time, error) {
	tz, err := time.ParseInLocation("2006-01-02T15:04", t, timeZone)
	if err != nil {
		return models.Time{}, err
	}
	return models.Time{Time: tz}, nil
}

func EditChallengeForm(chal *models.Challenge) *Form {
	vals := map[string]interface{}{
		"title":            chal.Info.Title,
		"description":      chal.Data.Description,
		"type":             chal.Info.Type,
		"locationDatabase": "direct",
		"eventStart":       formatTime(chal.Info.EventStart),
		"eventEnd":         formatTime(chal.Info.EventEnd),
		"restrictions":     chal.Info.Restrictions,
	}

	switch chal.Info.Type {
	default:
		vals["type"] = "phonecall"
	case "phonecall":
		vals["phoneDatabase"] = chal.Data.Database
		if chal.Data.Database == "direct" {
			vals["directphone"] = chal.Data.DirectPhone
		}
	case "location":
		vals["directaddr"] = chal.Data.DirectAddress
		vals["directlat"] = fmt.Sprint(chal.Data.DirectLatitude)
		vals["directlon"] = fmt.Sprint(chal.Data.DirectLongitude)
		vals["directradius"] = fmt.Sprint(chal.Data.DirectRadius)
	}

	if chal.Info.EventStart.Null() {
		if chal.Info.EventEnd.Null() {
			vals["dateType"] = "none"
		} else {
			vals["dateType"] = "deadline"
		}
	} else {
		vals["dateType"] = "event"
	}

	return &Form{
		Template: "challenge_form",
		Form:     vals,
	}
}

func ProcessChallengeForm(chal *models.Challenge, r *http.Request) (
	ok bool, f *Form) {

	chal.Info.Title = r.FormValue("title")
	chal.Data.Description = r.FormValue("description")
	chal.Info.Type = r.FormValue("type")

	switch chal.Info.Type {
	case "phonecall":
		chal.Data.Database = r.FormValue("phoneDatabase")
		chal.Data.DirectPhone = r.FormValue("directphone")

	case "location":
		chal.Data.Database = r.FormValue("locationDatabase")
		if chal.Data.Database != "direct" {
			whfatal.Error(wherr.BadRequest.New(
				"only direct locations currently supported"))
		}
		chal.Data.DirectAddress = r.FormValue("directaddr")
		latitude, err := strconv.ParseFloat(r.FormValue("directlat"), 64)
		if err != nil {
			whfatal.Error(wherr.BadRequest.Wrap(err))
		}
		longitude, err := strconv.ParseFloat(r.FormValue("directlon"), 64)
		if err != nil {
			whfatal.Error(wherr.BadRequest.Wrap(err))
		}
		radius, err := strconv.ParseFloat(r.FormValue("directradius"), 64)
		if err != nil {
			whfatal.Error(wherr.BadRequest.Wrap(err))
		}
		chal.Data.DirectLatitude = latitude
		chal.Data.DirectLongitude = longitude
		chal.Data.DirectRadius = radius

	default:
		whfatal.Error(wherr.BadRequest.New("bad challenge type: %s",
			chal.Info.Type))
	}

	switch chal.Data.Database {
	default:
		whfatal.Error(wherr.BadRequest.New("bad database type: %s",
			chal.Data.Database))
	case "direct", "us", "ushouse", "ussenate", "state", "statehouse",
		"statesenate", "usandstate":
	}

	restrictions, err := strconv.Atoi(r.FormValue("restrictionLength"))
	if err != nil {
		whfatal.Error(wherr.BadRequest.Wrap(err))
	}

	// reset
	chal.Info.Restrictions = make([]models.ChallengeRestriction, 0, restrictions)

	for i := 0; i < restrictions; i++ {
		rtype := r.FormValue(fmt.Sprintf("restrictionType[%d]", i))
		switch rtype {
		case "state":
		default:
			whfatal.Error(wherr.BadRequest.New(
				"unknown restriction type: %s", rtype))
		}
		chal.Info.Restrictions = append(chal.Info.Restrictions,
			models.ChallengeRestriction{
				Type:  rtype,
				Value: r.FormValue(fmt.Sprintf("restrictionValue[%d]", i))})
	}

	// reset
	chal.Info.EventStart = models.Time{}
	chal.Info.EventEnd = models.Time{}

	switch r.FormValue("dateType") {
	case "none":
	case "event":
		tz, err := parseTime(r.FormValue("eventStart"))
		if err != nil {
			whfatal.Error(wherr.BadRequest.Wrap(err))
		}
		chal.Info.EventStart = tz
		fallthrough
	case "deadline":
		tz, err := parseTime(r.FormValue("eventEnd"))
		if err != nil {
			whfatal.Error(wherr.BadRequest.Wrap(err))
		}
		chal.Info.EventEnd = tz
	default:
		whfatal.Error(wherr.BadRequest.New("bad date type: %s",
			r.FormValue("dateType")))
	}

	if chal.Info.Title == "" || chal.Data.Description == "" ||
		chal.Data.Database == "" ||
		(chal.Info.Type == "phonecall" && chal.Data.Database == "direct" &&
			chal.Data.DirectPhone == "") ||
		(chal.Info.Type == "location" && chal.Data.Database == "direct" &&
			(chal.Data.DirectAddress == "" || chal.Data.DirectLatitude == 0 ||
				chal.Data.DirectLongitude == 0)) {
		f := EditChallengeForm(chal)
		f.Error = "Required field missing"
		return false, f
	}

	return true, nil
}
