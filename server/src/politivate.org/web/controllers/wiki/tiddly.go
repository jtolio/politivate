// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wiki

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
	"gopkg.in/webhelp.v1/whcompat"

	"politivate.org/web/controllers/static"
)

var (
	mux     *http.ServeMux = http.NewServeMux()
	Handler http.Handler   = mux
)

func init() {
	mux.HandleFunc("/", main)
	mux.HandleFunc("/auth", auth)
	mux.HandleFunc("/status", status)
	mux.HandleFunc("/recipes/all/tiddlers/", tiddler)
	mux.HandleFunc("/recipes/all/tiddlers.json", tiddlerList)
	mux.HandleFunc("/bags/bag/tiddlers/", deleteTiddler)
}

type Tiddler struct {
	Rev  int    `datastore:"Rev,noindex"`
	Meta string `datastore:"Meta,noindex"`
	Text string `datastore:"Text,noindex"`
}

func main(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "bad method", 405)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "not found", 404)
		return
	}

	http.ServeContent(w, r, "index.html", time.Time{},
		bytes.NewReader(static.MustAsset("static/tiddly.html")))
}

func auth(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	u := user.Current(ctx)
	name := "GUEST"
	if u != nil {
		name = u.String()
	}
	fmt.Fprintf(w, "<html>\nYou are logged in as %s.\n\n<a href=\"/\">Main page</a>.\n", name)
}

func status(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "bad method", 405)
		return
	}
	ctx := whcompat.Context(r)
	w.Header().Set("Content-Type", "application/json")
	u := user.Current(ctx)
	name := "GUEST"
	if u != nil {
		name = u.String()
	}
	w.Write([]byte(`{"username": "` + name + `", "space": {"recipe": "all"}}`))
}

func tiddlerList(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	q := datastore.NewQuery("Tiddler")
	// Only need Meta, but get no results if we do this.
	if false {
		q = q.Project("Meta")
	}
	it := q.Run(ctx)
	var buf bytes.Buffer
	sep := ""
	buf.WriteString("[")
	for {
		var t Tiddler
		_, err := it.Next(&t)
		if err != nil {
			if err == datastore.Done {
				break
			}
			println("ERR", err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
		if len(t.Meta) == 0 {
			continue
		}
		meta := t.Meta

		// Tiddlers containing macros don't take effect until
		// they are loaded. Force them to be loaded by including
		// their bodies in the skinny tiddler list.
		// Might need to expand this to other kinds of tiddlers
		// in the future as we discover them.
		if strings.Contains(meta, `"$:/tags/Macro"`) {
			var js map[string]interface{}
			err := json.Unmarshal([]byte(meta), &js)
			if err != nil {
				continue
			}
			js["text"] = string(t.Text)
			data, err := json.Marshal(js)
			if err != nil {
				continue
			}
			meta = string(data)
		}

		buf.WriteString(sep)
		sep = ","
		buf.WriteString(meta)
	}
	buf.WriteString("]")
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf.Bytes())
}

func tiddler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTiddler(w, r)
	case "PUT":
		putTiddler(w, r)
	default:
		http.Error(w, "bad method", 405)
	}
}

func getTiddler(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	title := strings.TrimPrefix(r.URL.Path, "/recipes/all/tiddlers/")
	key := datastore.NewKey(ctx, "Tiddler", title, 0, nil)
	var t Tiddler
	if err := datastore.Get(ctx, key, &t); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var js map[string]interface{}
	err := json.Unmarshal([]byte(t.Meta), &js)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	js["text"] = string(t.Text)
	data, err := json.Marshal(js)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func putTiddler(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	title := strings.TrimPrefix(r.URL.Path, "/recipes/all/tiddlers/")
	key := datastore.NewKey(ctx, "Tiddler", title, 0, nil)
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cannot read data", 400)
		return
	}
	var js map[string]interface{}
	err = json.Unmarshal(data, &js)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	js["bag"] = "bag"

	rev := 1
	var old Tiddler
	if err := datastore.Get(ctx, key, &old); err == nil {
		rev = old.Rev + 1
	}
	js["revision"] = rev

	var t Tiddler
	text, ok := js["text"].(string)
	if ok {
		t.Text = text
	}
	delete(js, "text")
	t.Rev = rev
	meta, err := json.Marshal(js)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	t.Meta = string(meta)
	_, err = datastore.Put(ctx, key, &t)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	key2 := datastore.NewKey(ctx, "TiddlerHistory", title+"#"+fmt.Sprint(t.Rev), 0, nil)
	if _, err := datastore.Put(ctx, key2, &t); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	etag := fmt.Sprintf("\"bag/%s/%d:%x\"", url.QueryEscape(title), rev, md5.Sum(data))
	w.Header().Set("Etag", etag)
}

func deleteTiddler(w http.ResponseWriter, r *http.Request) {
	ctx := whcompat.Context(r)
	if r.Method != "DELETE" {
		http.Error(w, "bad method", 405)
		return
	}
	title := strings.TrimPrefix(r.URL.Path, "/bags/bag/tiddlers/")
	key := datastore.NewKey(ctx, "Tiddler", title, 0, nil)
	var t Tiddler
	if err := datastore.Get(ctx, key, &t); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	t.Rev++
	t.Meta = ""
	t.Text = ""
	if _, err := datastore.Put(ctx, key, &t); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	key2 := datastore.NewKey(ctx, "TiddlerHistory", title+"#"+fmt.Sprint(t.Rev), 0, nil)
	if _, err := datastore.Put(ctx, key2, &t); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
