package web

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/jtolds/webhelp"
	"github.com/jtolds/webhelp/sessions"
	"github.com/markbates/goth"
	"github.com/spacemonkeygo/errors"
	"golang.org/x/net/context"
)

type ProviderGroup struct {
	handlers       []*ProviderHandler
	handlersByName map[string]*ProviderHandler
	mux            webhelp.DirMux
	urls           RedirectURLs
	groupBaseURL   string
}

func NewProviderGroup(sessionNamespace string, groupBaseURL string,
	urls RedirectURLs, providers ...goth.Provider) (*ProviderGroup, error) {

	groupBaseURL = strings.TrimRight(groupBaseURL, "/")

	g := &ProviderGroup{
		handlers:       make([]*ProviderHandler, 0, len(providers)),
		handlersByName: make(map[string]*ProviderHandler, len(providers)),
		urls:           urls,
		groupBaseURL:   groupBaseURL}

	g.mux = webhelp.DirMux{
		"all": webhelp.DirMux{"logout": webhelp.Exact(
			http.HandlerFunc(g.logoutAll))},
	}

	for _, provider := range providers {
		name := provider.Name()
		if name == "" {
			return nil, fmt.Errorf("empty provider name")
		}
		_, exists := g.mux[name]
		if exists {
			return nil, fmt.Errorf("two providers given with name %#v", name)
		}
		handler := NewProviderHandler(provider,
			fmt.Sprintf("%s-%s", sessionNamespace, name),
			fmt.Sprintf("%s/%s", groupBaseURL, name), urls)
		g.handlers = append(g.handlers, handler)
		g.handlersByName[name] = handler
		g.mux[name] = handler
	}

	return g, nil
}

func (g *ProviderGroup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.mux.ServeHTTP(w, r)
}

func (g *ProviderGroup) Routes(
	cb func(method, path string, annotations []string)) {
	webhelp.Routes(g.mux, cb)
}

var _ http.Handler = (*ProviderGroup)(nil)
var _ webhelp.RouteLister = (*ProviderGroup)(nil)

// Handler returns a specific ProviderHandler given the Provider name
func (g *ProviderGroup) Handler(providerName string) (rv *ProviderHandler,
	exists bool) {
	rv, exists = g.handlersByName[providerName]
	return rv, exists
}

// LoginURL returns the login URL for a given provider.
// redirectTo is the URL to navigate to after logging in.
func (g *ProviderGroup) LoginURL(providerName, redirectTo string) string {
	return g.handlersByName[providerName].LoginURL(redirectTo)
}

// LogoutURL returns the logout URL for a given provider.
// redirectTo is the URL to navigate to after logging out.
func (g *ProviderGroup) LogoutURL(providerName, redirectTo string) string {
	return g.handlersByName[providerName].LogoutURL(redirectTo)
}

// LogoutAllURL returns the logout URL for all providers.
// redirectTo is the URL to navigate to after logging out.
func (g *ProviderGroup) LogoutAllURL(redirectTo string) string {
	return g.groupBaseURL + "/all/logout?" + url.Values{
		"redirect_to": {redirectTo}}.Encode()
}

func (g *ProviderGroup) Providers() []*ProviderHandler {
	copy := make([]*ProviderHandler, 0, len(g.handlers))
	for _, handler := range g.handlers {
		copy = append(copy, handler)
	}
	return copy
}

func (g *ProviderGroup) User(ctx context.Context) (goth.User, error) {
	for _, handler := range g.handlers {
		if handler.LoggedIn(ctx) {
			return handler.User(ctx)
		}
	}
	return goth.User{}, webhelp.ErrBadRequest.New("no logged in user")
}

func (g *ProviderGroup) LoggedIn(ctx context.Context) bool {
	for _, handler := range g.handlers {
		if handler.LoggedIn(ctx) {
			return true
		}
	}
	return false
}

// LogoutAll will not return any HTTP response, but will simply prepare a
// response for logging a user out completely from all providers. If a user
// should log out of just a specific provider, use the Logout method
// on the associated ProviderHandler.
func (g *ProviderGroup) LogoutAll(ctx context.Context,
	w http.ResponseWriter) error {
	var errs errors.ErrorGroup
	for _, handler := range g.handlers {
		errs.Add(handler.Logout(ctx, w))
	}
	return errs.Finalize()
}

func (g *ProviderGroup) logoutAll(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	err := g.LogoutAll(ctx, w)
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}
	redirectTo := r.FormValue("redirect_to")
	if redirectTo == "" {
		redirectTo = g.urls.DefaultLogoutURL
	}
	webhelp.Redirect(w, r, redirectTo)
}

// LoginRequired is a middleware for redirecting users to a login page if
// they aren't logged in yet. loginRedirect should take the URL to redirect
// to after logging in and return a URL that will actually do the logging in.
// If you already know which provider a user should use, consider using
// (*ProviderHandler).LoginRequired instead, which doesn't require a
// loginRedirect URL.
func (g *ProviderGroup) LoginRequired(h http.Handler,
	loginRedirect func(redirectTo string) (url string)) http.Handler {
	return webhelp.RouteHandlerFunc(h,
		func(w http.ResponseWriter, r *http.Request) {
			ctx := webhelp.Context(r)
			if !g.LoggedIn(ctx) {
				webhelp.Redirect(w, r, loginRedirect(r.RequestURI))
			} else {
				h.ServeHTTP(w, r)
			}
		})
}

type ProviderHandler struct {
	provider         goth.Provider
	sessionNamespace string
	handlerBaseURL   string
	urls             RedirectURLs
	mux              webhelp.DirMux
}

// NewProviderHandler makes a provider handler. Requires a provider
// configuration, a session namespace, a base URL for the handler, and a
// collection of URLs for redirecting.
func NewProviderHandler(provider goth.Provider, sessionNamespace string,
	handlerBaseURL string, urls RedirectURLs) *ProviderHandler {
	if urls.DefaultLoginURL == "" {
		urls.DefaultLoginURL = "/"
	}
	if urls.DefaultLogoutURL == "" {
		urls.DefaultLogoutURL = "/"
	}
	h := &ProviderHandler{
		provider:         provider,
		sessionNamespace: sessionNamespace,
		handlerBaseURL:   strings.TrimRight(handlerBaseURL, "/"),
		urls:             urls}
	h.mux = webhelp.DirMux{
		"login":  webhelp.Exact(http.HandlerFunc(h.login)),
		"logout": webhelp.Exact(http.HandlerFunc(h.logout)),
		"_cb":    webhelp.Exact(http.HandlerFunc(h.cb))}
	return h
}

func (p *ProviderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.mux.ServeHTTP(w, r)
}

func (p *ProviderHandler) Name() string { return p.provider.Name() }

// Session returns a provider-specific authenticated session for the current
// user. This session is cleared whenever a user logs out.
func (o *ProviderHandler) Session(ctx context.Context) (*sessions.Session,
	error) {
	return sessions.Load(ctx, o.sessionNamespace)
}

// Logout prepares the request to log the user out of just this provider.
// If you're using a ProviderGroup you may be interested in LogoutAll.
func (o *ProviderHandler) Logout(ctx context.Context,
	w http.ResponseWriter) error {
	session, err := o.Session(ctx)
	if err != nil {
		return err
	}
	return session.Clear(w)
}

// LoginURL returns the login URL for this provider.
// redirectTo is the URL to navigate to after logging in.
func (o *ProviderHandler) LoginURL(redirectTo string) string {
	return o.handlerBaseURL + "/login?" + url.Values{
		"redirect_to": {redirectTo}}.Encode()
}

// LogoutURL returns the logout URL for this provider
// redirectTo is the URL to navigate to after logging out.
func (o *ProviderHandler) LogoutURL(redirectTo string) string {
	return o.handlerBaseURL + "/logout?" + url.Values{
		"redirect_to": {redirectTo}}.Encode()
}

func (o *ProviderHandler) login(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)

	redirectTo := r.FormValue("redirect_to")
	if redirectTo == "" {
		redirectTo = o.urls.DefaultLoginURL
	}

	session, err := o.Session(ctx)
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}

	if o.loggedIn(session) {
		webhelp.Redirect(w, r, redirectTo)
		return
	}

	state := newState()

	sess, err := o.provider.BeginAuth(state)
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}

	url, err := sess.GetAuthURL()
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}

	session.Values = map[interface{}]interface{}{
		"_state":       state,
		"_redirect_to": redirectTo,
		"_data":        sess.Marshal(),
		"_logged_in":   false,
	}
	err = session.Save(w)
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}

	webhelp.Redirect(w, r, url)
}

func (o *ProviderHandler) cb(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	session, err := o.Session(ctx)
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}

	val, exists := session.Values["_state"]
	existingState, correct := val.(string)
	if !exists || !correct {
		session.Clear(w)
		webhelp.HandleError(w, r, webhelp.ErrBadRequest.New("invalid session"))
		return
	}

	if existingState != r.FormValue("state") {
		webhelp.HandleError(w, r, webhelp.ErrBadRequest.New("csrf detected"))
		return
	}

	val, exists = session.Values["_redirect_to"]
	redirectTo, correct := val.(string)
	if !exists || !correct {
		redirectTo = o.urls.DefaultLoginURL
	}

	val, exists = session.Values["_data"]
	data, correct := val.(string)
	if !exists || !correct {
		session.Clear(w)
		webhelp.HandleError(w, r, webhelp.ErrBadRequest.New("invalid session"))
		return
	}

	sess, err := o.provider.UnmarshalSession(data)
	if err != nil {
		session.Clear(w)
		webhelp.HandleError(w, r, err)
		return
	}

	_, err = sess.Authorize(o.provider, r.URL.Query())
	if err != nil {
		session.Clear(w)
		webhelp.HandleError(w, r, err)
		return
	}

	session.Values["_data"] = sess.Marshal()
	session.Values["_logged_in"] = true
	err = session.Save(w)
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}

	webhelp.Redirect(w, r, redirectTo)
}

func (o *ProviderHandler) logout(w http.ResponseWriter, r *http.Request) {
	ctx := webhelp.Context(r)
	err := o.Logout(ctx, w)
	if err != nil {
		webhelp.HandleError(w, r, err)
		return
	}
	redirectTo := r.FormValue("redirect_to")
	if redirectTo == "" {
		redirectTo = o.urls.DefaultLogoutURL
	}
	webhelp.Redirect(w, r, redirectTo)
}

func (o *ProviderHandler) loggedIn(session *sessions.Session) bool {
	loggedIn, _ := session.Values["_logged_in"].(bool)
	return loggedIn
}

func (o *ProviderHandler) LoggedIn(ctx context.Context) bool {
	session, err := o.Session(ctx)
	if err != nil {
		return false
	}
	return o.loggedIn(session)
}

func (o *ProviderHandler) User(ctx context.Context) (goth.User, error) {
	session, err := o.Session(ctx)
	if err != nil {
		return goth.User{}, err
	}
	if !o.loggedIn(session) {
		return goth.User{}, webhelp.ErrBadRequest.New("not logged in")
	}

	val, exists := session.Values["_data"]
	data, correct := val.(string)
	if !exists || !correct {
		return goth.User{}, webhelp.ErrBadRequest.New("invalid session")
	}

	sess, err := o.provider.UnmarshalSession(data)
	if err != nil {
		return goth.User{}, err
	}

	return o.provider.FetchUser(sess)
}

// LoginRequired is a middleware for redirecting users to a login page if
// they aren't logged in yet. If you are using a ProviderGroup and don't know
// which provider a user should use, consider using
// (*ProviderGroup).LoginRequired instead
func (o *ProviderHandler) LoginRequired(h http.Handler) http.Handler {
	return webhelp.RouteHandlerFunc(h,
		func(w http.ResponseWriter, r *http.Request) {
			ctx := webhelp.Context(r)
			if o.LoggedIn(ctx) {
				h.ServeHTTP(w, r)
			} else {
				webhelp.Redirect(w, r, o.LoginURL(r.RequestURI))
			}
		})
}

type RedirectURLs struct {
	// If a login URL isn't provided to redirect to after successful login, use
	// this one.
	DefaultLoginURL string

	// If a logout URL isn't provided to redirect to after successful logout, use
	// this one.
	DefaultLogoutURL string
}

func newState() string {
	var p [16]byte
	_, err := rand.Read(p[:])
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(p[:])
}
