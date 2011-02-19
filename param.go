// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import (
	"regexp"
	"strconv"
	"fmt"
)

type Param struct {
	Name     string
	Desc     string
	Pattern  *regexp.Regexp
	Optional bool
	Value    string
}

func newParam(name, desc string, pattern *regexp.Regexp, optional bool) *Param {
	return &Param{name, desc, pattern, optional, ""}
}

func (this *Param) IsValid(val string) bool {
	return this.Pattern.MatchString(val)
}

func (this *Param) String() string {
	ret := fmt.Sprintf("<%s>", this.Name)
	if this.Optional {
		ret = fmt.Sprintf("[%s]", ret)
	}
	return ret
}

func (this *Param) Int() int {
	if i, err := strconv.Atoi(this.Value); err == nil {
		return i
	}
	return 0
}

func (this *Param) Int64() int64 {
	if i, err := strconv.Atoi64(this.Value); err == nil {
		return i
	}
	return 0
}

func (this *Param) Byte() byte {
	return byte(this.Int())
}
