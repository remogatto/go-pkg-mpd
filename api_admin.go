// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import "os"

func disableoutput(cmd *Command, c *Client) (err os.Error) {
	return c.request("disableoutput %d", cmd.I("id", 0))
}

func enableoutput(cmd *Command, c *Client) (err os.Error) {
	return c.request("enableoutput %d", cmd.I("id", 0))
}

func kill(cmd *Command, c *Client) (err os.Error) {
	return c.request("kill")
}

func update(cmd *Command, c *Client) (err os.Error) {
	var path string
	if path = cmd.S("path", ""); path == "" {
		return c.request("update")
	}
	return c.request("update \"%s\"", path)
}
