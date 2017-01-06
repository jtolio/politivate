package gov

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/webhelp.v1/whcompat"
	"gopkg.in/webhelp.v1/wherr"
	"gopkg.in/webhelp.v1/whfatal"
	"gopkg.in/webhelp.v1/whjson"
	"gopkg.in/webhelp.v1/whmux"
)

func init() {

	mux["district"] = whmux.Dir{
		"locate": whmux.Exact(http.HandlerFunc(districtLocate)),
		"reps":   whmux.Exact(http.HandlerFunc(repsByDistrict)),
	}
}

func contains(vals map[string]string, name string) bool {
	_, exists := vals[name]
	return exists
}

func districtLocate(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	vals := map[string]string{}
	for _, name := range []string{"zip", "latitude", "longitude"} {
		val := r.FormValue(name)
		if val != "" {
			vals[name] = val
		}
	}

	if !contains(vals, "zip") &&
		(!contains(vals, "latitude") || !contains(vals, "longitude")) {
		whfatal.Error(wherr.BadRequest.New("missing argument"))
	}

	resp, err := SunlightAPIReq(ctx, "/districts/locate", vals)
	if err != nil {
		whfatal.Error(wherr.InternalServerError.Wrap(err))
	}
	defer resp.Body.Close()

	var districts struct {
		Results []struct {
			State    string `json:"state"`
			District int    `json:"district"`
		} `json:"results"`
	}
	err = json.NewDecoder(resp.Body).Decode(&districts)
	if err != nil {
		whfatal.Error(wherr.InternalServerError.Wrap(err))
	}

	whjson.Render(w, r, districts.Results)
}

func repsByDistrict(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	vals := map[string]string{"chamber": "house"}
	for _, name := range []string{"state", "district"} {
		val := r.FormValue(name)
		if val != "" {
			vals[name] = val
		}
	}

	if !contains(vals, "state") || !contains(vals, "district") {
		whfatal.Error(wherr.BadRequest.New("missing argument"))
	}

	type legislatorResult struct {
		BioguideId     string   `json:"bioguide_id"`
		Birthday       string   `json:"birthday"`
		Chamber        string   `json:"chamber"`
		ContactForm    string   `json:"contact_form"`
		CRPId          string   `json:"crp_id"`
		District       int      `json:"district"`
		FacebookId     string   `json:"facebook_id"`
		Fax            string   `json:"fax"`
		FECIds         []string `json:"fec_ids"`
		FirstName      string   `json:"first_name"`
		Gender         string   `json:"gender"`
		GovTrackId     string   `json:"govtrack_id"`
		ICPSRId        int64    `json:"icpsr_id"`
		InOffice       bool     `json:"in_office"`
		LastName       string   `json:"last_name"`
		LeadershipRole string   `json:"leadership_role"`
		LISMemberId    string   `json:"lis_id"`
		MiddleName     string   `json:"middle_name"`
		NameSuffix     string   `json:"name_suffix"`
		NickName       string   `json:"nickname"`
		OCDId          string   `json:"ocd_id"`
		OCEmail        string   `json:"oc_email"`
		Office         string   `json:"office"`
		Party          string   `json:"party"`
		Phone          string   `json:"phone"`
		SenateClass    int      `json:"senate_class"`
		State          string   `json:"state"`
		StateName      string   `json:"state_name"`
		StateRank      string   `json:"state_rank"`
		TermEnd        string   `json:"term_end"`
		TermStart      string   `json:"term_start"`
		ThomasId       string   `json:"thomas_id"`
		Title          string   `json:"title"`
		TwitterId      string   `json:"twitter_id"`
		VotesmartId    int64    `json:"votesmart_id"`
		Website        string   `json:"website"`
		YoutubeId      string   `json:"youtube_id"`
	}

	type legislatorResults struct {
		Count int `json:"count"`
		Page  struct {
			Count   int `json:"count"`
			PerPage int `json:"per_page"`
			Page    int `json:page"`
		}
		Results []legislatorResult `json:"results"`
	}

	var legislators []legislatorResult
	page := 0

	for {
		page += 1
		vals["page"] = fmt.Sprint(page)
		resp, err := SunlightAPIReq(ctx, "/legislators", vals)
		if err != nil {
			whfatal.Error(wherr.InternalServerError.Wrap(err))
		}
		defer resp.Body.Close()

		result := legislatorResults{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			whfatal.Error(wherr.InternalServerError.Wrap(err))
		}

		legislators = append(legislators, result.Results...)

		if result.Page.Page*result.Page.PerPage >= result.Count {
			if vals["chamber"] == "senate" {
				break
			}
			page = 0
			vals["chamber"] = "senate"
			delete(vals, "district")
		}
	}

	whjson.Render(w, r, legislators)
}
