package gov

import (
	"net/http"
	"net/url"

	"github.com/spacemonkeygo/spacelog"
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
	"gopkg.in/webhelp.v1/whmux"
)

var (
	Handler http.Handler = mux
	mux                  = whmux.Dir{}

	logger = spacelog.GetLogger()
)

func SunlightAPIReq(ctx context.Context, path string, vals map[string]string) (
	*http.Response, error) {
	query := url.Values{}
	for name, val := range vals {
		query.Add(name, val)
	}
	req, err := http.NewRequest("GET",
		"https://congress.api.sunlightfoundation.com"+path+"?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	return urlfetch.Client(ctx).Do(req)
}
