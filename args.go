// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import (
	"os"
	"fmt"
	"strconv"
)

type Args map[string]string

func (this Args) Bool(k string, d bool) bool {
	r := 0
	if d {
		r = 1
	}
	return this.Int(k, r) == 1
}

func (this Args) Byte(k string, d byte) byte {
	return byte(this.Int(k, int(d)))
}

func (this Args) Int(k string, d int) int {
	if v, e := strconv.Atoi(this.read(k)); e == nil {
		return v
	}
	return d
}

func (this Args) Int32(k string, d int32) int32 {
	return int32(this.Int(k, int(d)))
}

func (this Args) Int64(k string, d int64) int64 {
	if v, e := strconv.Atoi64(this.read(k)); e == nil {
		return v
	}
	return d
}

func (this Args) String(k, d string) string {
	if v := this.read(k); v != "" {
		return v
	}
	return d
}

func (this Args) read(key string) string {
	v, ok := this[key]
	if !ok {
		return ""
	}
	return v
}

func (this Args) Print() {
	for k, v := range this {
		fmt.Fprintf(os.Stdout, "%v : %v\n", k, v)
	}
}
