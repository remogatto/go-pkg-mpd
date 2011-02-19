// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import "fmt"

// Miscellaneous helper functions

// parses number of seconds into (hh:mm:ss) format.
func parseTime(seconds int) string {
	time := [3]int{0, 0, 0}

	time[0] = seconds / 3600
	seconds %= 3600

	time[1] = seconds / 60
	seconds %= 60

	time[2] = seconds

	if time[0] == 0 { // no need to list hours if they aren't there.
		return fmt.Sprintf("%02d:%02d", time[1], time[2])
	} else {
		return fmt.Sprintf("%02d:%02d:%02d", time[0], time[1], time[2])
	}

	return "00:00"
}

// simply converts 1 to 'on' and 0 to 'off'
func onoff(v string) string {
	if v == "1" {
		return "on"
	}
	return "off"
}
