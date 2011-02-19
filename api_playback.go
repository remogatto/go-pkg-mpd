// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import (
	"os"
	"strconv"
)

func toggle(cmd *Command, c *Client) (err os.Error) {
	var arg Args
	if arg, err = c.requestArgs("status"); err != nil {
		return
	}

	if arg["state"] == "play" {
		return c.request("pause 1")
	}
	return c.request("play")
}

func crossfade(cmd *Command, c *Client) (err os.Error) {
	return c.request("crossfade %d", cmd.I("time", 0))
}

func next(cmd *Command, c *Client) (err os.Error) {
	return c.request("next")
}

func pause(cmd *Command, c *Client) (err os.Error) {
	t := cmd.S("toggle", "")
	v := 0
	if t == "on" {
		v = 1
	}
	return c.request("pause %d", v)
}

func play(cmd *Command, c *Client) (err os.Error) {
	return c.request("play %d", cmd.I("pos", 0))
}

func playid(cmd *Command, c *Client) (err os.Error) {
	return c.request("playid %d", cmd.I("id", 0))
}

func previous(cmd *Command, c *Client) (err os.Error) {
	return c.request("previous")
}

func random(cmd *Command, c *Client) (err os.Error) {
	t := cmd.S("toggle", "")
	v := 0
	if t == "on" {
		v = 1
	}
	return c.request("random %d", v)
}

func repeat(cmd *Command, c *Client) (err os.Error) {
	t := cmd.S("toggle", "")
	v := 0
	if t == "on" {
		v = 1
	}
	return c.request("repeat %d", v)
}

func seek(cmd *Command, c *Client) (err os.Error) {
	return c.request(
		"seek %d %d",
		cmd.I("pos", 0),
		cmd.I("time", 0),
	)
}

func seekid(cmd *Command, c *Client) (err os.Error) {
	return c.request(
		"seekid %d %d",
		cmd.I("id", 0),
		cmd.I("time", 0),
	)
}

func volume(cmd *Command, c *Client) (err os.Error) {
	v := cmd.I("value", -1)
	if v < 0 || v > 100 {
		return os.NewError("Volume must be in the range 0-100.")
	}

	if s := cmd.S("sign", ""); s != "" {
		var arg Args
		var cv int

		if s != "+" && s != "-" {
			// this should be caught by the patSign regex pattern, but the current
			// regexp implementation has some issues with escaping +/- signs. eg:
			// it isn't supported at all.
			return os.NewError("Invalid value for parameter @sign. Expected + or -")
		}

		if arg, err = c.requestArgs("status"); err != nil {
			return
		}

		if cv, err = strconv.Atoi(arg["volume"]); err != nil {
			return
		}

		if s == "+" {
			if v += cv; v > 100 {
				v = 100
			}
		} else {
			if v = cv - v; v < 0 {
				v = 0
			}
		}
	}

	return c.request("setvol %d", v)
}

func stop(cmd *Command, c *Client) (err os.Error) {
	return c.request("stop")
}
