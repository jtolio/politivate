package gov

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/spacemonkeygo/spacelog"
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"
)

var (
	Handler http.Handler = mux
	mux                  = whmux.Dir{
		"legislators": whmux.Dir{
			"locate": whmux.Exact(http.HandlerFunc(legislatorsByGPS)),
		},
	}

	logger = spacelog.GetLogger()
)

func apiReq(ctx context.Context, uri string, vals map[string]string) (
	*http.Response, error) {
	query := url.Values{}
	for name, val := range vals {
		query.Add(name, val)
	}
	req, err := http.NewRequest("GET", uri+"?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := urlfetch.Client(ctx).Do(req)
	if err != nil {
		return nil, wherr.InternalServerError.Wrap(err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		resp.Body.Close()
		return nil, wherr.InternalServerError.New(
			"Invalid api status code returned from %s: %s", uri, resp.Status)
	}
	return resp, nil
}

func contains(vals map[string]string, name string) bool {
	_, exists := vals[name]
	return exists
}

func titlePrefix(chamber string) string {
	switch chamber {
	case "senate", "upper":
		return "Sen. "
	case "house", "lower":
		return "Rep. "
	default:
		return ""
	}
}

type Office struct {
	Phone string `json:"phone"`
}

type Legislator struct {
	FullName string   `json:"full_name"` // includes title.
	Offices  []Office `json:"offices"`
	Website  string   `json:"website"`
}

func (l *Legislator) MarshalJSON() ([]byte, error) {
	var phone string
	if len(l.Offices) > 0 {
		phone = l.Offices[0].Phone
	}

	m := map[string]interface{}{
		"full_name": l.FullName,
		"offices":   l.Offices,
		"website":   l.Website,
		"comment": `the fields "chamber", "votesmart_id", "first_name",
                "last_name", and "phone" are all deprecated.`,
		"first_name":   "",
		"last_name":    "",
		"chamber":      "",
		"phone":        phone,
		"votesmart_id": l.FullName + ": " + phone,
	}

	fullName := l.FullName
	if strings.HasPrefix(fullName, "Sen. ") {
		m["chamber"] = "senate"
		fullName = fullName[5:]
	} else if strings.HasPrefix(fullName, "Rep. ") {
		m["chamber"] = "house"
		fullName = fullName[5:]
	}
	nameParts := strings.SplitN(fullName, " ", 2)
	m["first_name"] = nameParts[0]
	if len(nameParts) > 1 {
		m["last_name"] = nameParts[1]
	}

	return json.Marshal(m)
}

func LegislatorsByGPS(ctx context.Context, database string,
	latitude, longitude float64) (rv []*Legislator) {
	rv = make([]*Legislator, 0) // json not null
	var federalDistricts *[]FederalDistrict
	if database == "ushouse" || database == "us" || database == "usandstate" {
		fds := FederalDistrictLocateByGPS(ctx, latitude, longitude)
		federalDistricts = &fds
		for _, fd := range *federalDistricts {
			for _, l := range HouseRepsByFederalDistrict(ctx, fd) {
				rv = append(rv, l.Convert())
			}
		}
	}
	if database == "ussenate" || database == "us" || database == "usandstate" {
		if federalDistricts == nil {
			fds := FederalDistrictLocateByGPS(ctx, latitude, longitude)
			federalDistricts = &fds
		}
		for _, fd := range *federalDistricts {
			for _, l := range SenatorsByFederalDistrict(ctx, fd) {
				rv = append(rv, l.Convert())
			}
		}
	}
	var stateReps *[]*OpenStatesLegislator
	if database == "statehouse" || database == "state" || database == "usandstate" {
		srs := StateRepsLocateByGPS(ctx, latitude, longitude)
		stateReps = &srs
		for _, sr := range *stateReps {
			if sr.Chamber == "lower" {
				rv = append(rv, sr.Convert())
			}
		}
	}
	if database == "statesenate" || database == "state" || database == "usandstate" {
		if stateReps == nil {
			srs := StateRepsLocateByGPS(ctx, latitude, longitude)
			stateReps = &srs
		}
		for _, sr := range *stateReps {
			if sr.Chamber == "upper" {
				rv = append(rv, sr.Convert())
			}
		}
	}
	return rv
}

func legislatorsByGPS(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	latitude, err := strconv.ParseFloat(r.FormValue("latitude"), 64)
	if err != nil {
		whfatal.Error(wherr.BadRequest.New("missing argument"))
	}
	longitude, err := strconv.ParseFloat(r.FormValue("longitude"), 64)
	if err != nil {
		whfatal.Error(wherr.BadRequest.New("missing argument"))
	}
	whjson.Render(w, r, LegislatorsByGPS(ctx, r.FormValue("database"),
		latitude, longitude))
}
