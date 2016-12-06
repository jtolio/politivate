// Copyright (C) 2016 JT Olds
// See LICENSE for copying information

package webhelp

import (
	"net/http"
	"sort"
	"strings"

	"golang.org/x/net/context"
)

// DirMux is a Handler that mimics a directory. It mutates an incoming
// request's URL.Path to properly namespace handlers. This way a handler can
// assume it has the root of its section. If you want the original URL, use
// req.RequestURI (but don't modify it).
type DirMux map[string]Handler

func (d DirMux) HandleHTTP(ctx context.Context, w ResponseWriter,
	r *http.Request) error {
	dir, left := Shift(r.URL.Path)
	handler, ok := d[dir]
	if !ok {
		return ErrNotFound.New("resource: %#v", dir)
	}
	r.URL.Path = left
	return handler.HandleHTTP(ctx, w, r)
}

func (d DirMux) Routes(cb func(method, path string, annotations []string)) {
	keys := make([]string, 0, len(d))
	for element := range d {
		keys = append(keys, element)
	}
	sort.Strings(keys)
	for _, element := range keys {
		Routes(d[element], func(method, path string, annotations []string) {
			if element == "" {
				cb(method, "/", annotations)
			} else {
				cb(method, "/"+element+path, annotations)
			}
		})
	}
}

var _ Handler = DirMux(nil)
var _ RouteLister = DirMux(nil)

// Shift pulls the first directory out of the path and returns the remainder.
func Shift(path string) (dir, left string) {
	// slice off the first "/"s if they exists
	path = strings.TrimLeft(path, "/")

	if len(path) == 0 {
		return "", ""
	}

	// find the first '/' after the initial one
	split := strings.Index(path, "/")
	if split == -1 {
		return path, ""
	}
	return path[:split], path[split:]
}

// MethodMux is a Handler muxer that keys off of the given HTTP request method
type MethodMux map[string]Handler

func (m MethodMux) HandleHTTP(ctx context.Context, w ResponseWriter,
	r *http.Request) error {
	if handler, found := m[r.Method]; found {
		return handler.HandleHTTP(ctx, w, r)
	}
	return ErrMethodNotAllowed.New("bad method: %#v", r.Method)
}

func (m MethodMux) Routes(cb func(method, path string, annotations []string)) {
	keys := make([]string, 0, len(m))
	for method := range m {
		keys = append(keys, method)
	}
	sort.Strings(keys)
	for _, method := range keys {
		handler := m[method]
		Routes(handler, func(_, path string, annotations []string) {
			cb(method, path, annotations)
		})
	}
}

var _ Handler = MethodMux(nil)
var _ RouteLister = MethodMux(nil)

// ExactPath takes a Handler that returns a new Handler that doesn't accept any
// more path elements
func ExactPath(h Handler) Handler {
	return DirMux{"": h}
}

// ExactMethod takes a Handler and returns a new Handler that only works with
// the given HTTP method.
func ExactMethod(method string, h Handler) Handler {
	return MethodMux{method: h}
}

// ExactGet is simply ExactMethod but called with "GET" as the first argument.
func ExactGet(h Handler) Handler {
	return ExactMethod("GET", h)
}

// Exact is simply ExactGet and ExactPath called together.
func Exact(h Handler) Handler {
	return ExactGet(ExactPath(h))
}

// OverlayMux is essentially a DirMux that you can put in front of another
// Handler. If the requested entry isn't in the overlay DirMux, the Fallback
// will be used. If no Fallback is specified this works exactly the same as
// a DirMux.
type OverlayMux struct {
	Fallback Handler
	Overlay  DirMux
}

func (o OverlayMux) HandleHTTP(ctx context.Context, w ResponseWriter,
	r *http.Request) error {
	dir, left := Shift(r.URL.Path)
	handler, ok := o.Overlay[dir]
	if !ok {
		if o.Fallback == nil {
			return ErrNotFound.New("resource: %#v", dir)
		}
		return o.Fallback.HandleHTTP(ctx, w, r)
	}
	r.URL.Path = left
	return handler.HandleHTTP(ctx, w, r)
}

func (o OverlayMux) Routes(cb func(method, path string, annotations []string)) {
	Routes(o.Overlay, cb)
	if o.Fallback != nil {
		Routes(o.Fallback, cb)
	}
}

var _ Handler = OverlayMux{}
var _ RouteLister = OverlayMux{}

type RedirectHandler string

func (target RedirectHandler) HandleHTTP(ctx context.Context, w ResponseWriter,
	r *http.Request) error {
	return Redirect(w, r, string(target))
}

func (target RedirectHandler) Routes(
	cb func(method, path string, annotations []string)) {
	cb(AllMethods, AllPaths, []string{"-> " + string(target)})
}

type RedirectHandlerFunc func(ctx context.Context, r *http.Request) string

func (f RedirectHandlerFunc) HandleHTTP(ctx context.Context, w ResponseWriter,
	r *http.Request) error {
	return Redirect(w, r, f(ctx, r))
}

func (f RedirectHandlerFunc) Routes(
	cb func(method, path string, annotations []string)) {
	cb(AllMethods, AllPaths, []string{"-> f(req)"})
}
