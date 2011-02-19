// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import "regexp"

var (
	PatAny     = regexp.MustCompile(`^.+$`)
	PatInteger = regexp.MustCompile(`^[0-9]+$`)
	PatType    = regexp.MustCompile(`^any|artist|album|title|track|name|genre|date|composer|performer|comment|disc|filename$`)
	PatOnOff   = regexp.MustCompile(`^on|off$`)
	PatSign    = regexp.MustCompile(`^+|-$`) // this doesn't actually work as intended.
)
