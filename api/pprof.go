// Copyright 2013 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/tsuru/tsuru/auth"
	"net/http"
	"net/http/pprof"
)

func indexHandler(w http.ResponseWriter, r *http.Request, t auth.Token) error {
	pprof.Index(w, r)
	return nil
}

func cmdlineHandler(w http.ResponseWriter, r *http.Request, t auth.Token) error {
	pprof.Cmdline(w, r)
	return nil
}

func profileHandler(w http.ResponseWriter, r *http.Request, t auth.Token) error {
	pprof.Profile(w, r)
	return nil
}

func symbolHandler(w http.ResponseWriter, r *http.Request, t auth.Token) error {
	pprof.Symbol(w, r)
	return nil
}
