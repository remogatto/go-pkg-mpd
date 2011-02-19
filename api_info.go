// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import (
	"os"
	"fmt"
)

func status(cmd *Command, c *Client) (err os.Error) {
	return c.request("status")
}

func simplestatus(cmd *Command, c *Client) (err os.Error) {
	var a Args
	if a, err = c.requestArgs("status"); err != nil {
		return
	}

	fmt.Printf(
		"[%s] vol: %s%%, repeat %s, single %s, random %s, consume %s\n",
		a["state"], a["volume"], onoff(a["repeat"]), onoff(a["single"]),
		onoff(a["random"]), onoff(a["consume"]),
	)
	return
}

func stats(cmd *Command, c *Client) (err os.Error) {
	return c.request("stats")
}

func outputs(cmd *Command, c *Client) (err os.Error) {
	return c.requestList("outputs")
}

func commands(cmd *Command, c *Client) (err os.Error) {
	return c.requestList("commands")
}

func notcommands(cmd *Command, c *Client) (err os.Error) {
	return c.requestList("notcommands")
}

func tagtypes(cmd *Command, c *Client) (err os.Error) {
	return c.requestList("tagtypes")
}

func urlhandlers(cmd *Command, c *Client) (err os.Error) {
	return c.requestList("urlhandlers")
}
