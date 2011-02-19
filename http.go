// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import (
	"http"
	"strings"
	"os"
	"fmt"
	"io/ioutil"
)

func httpGet(uri string, params map[string]interface{}) (response string, err os.Error) {
	var r *http.Response

	if r, _, err = http.Get(fmt.Sprintf("%s?%s", uri, makeQueryString(params))); err != nil {
		return
	}
	defer r.Body.Close()

	var b []byte
	if b, err = ioutil.ReadAll(r.Body); err != nil {
		return
	}

	return string(b), nil
}

func httpPost(uri string, params map[string]interface{}) (response string, err os.Error) {
	var r *http.Response

	body := strings.NewReader(makeQueryString(params))
	if r, err = http.Post(uri, "text/html; charset=utf-8", body); err != nil {
		return
	}
	defer r.Body.Close()

	var b []byte
	if b, err = ioutil.ReadAll(r.Body); err != nil {
		return
	}

	return string(b), nil
}

func makeQueryString(params map[string]interface{}) (qs string) {
	for k, v := range params {
		qs += fmt.Sprintf("%s=%v&", k, http.URLEscape(fmt.Sprintf("%v", v)))
	}
	if len(qs) > 1 {
		qs = qs[0 : len(qs)-1] // strip trailing &
	}
	return
}
