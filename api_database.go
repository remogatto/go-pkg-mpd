// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import "os"
import "fmt"

func find(cmd *Command, c *Client) (err os.Error) {
	return c.requestList(
		"find \"%s\" \"%s\"",
		cmd.S("tag", "any"),
		cmd.S("term", ""),
	)
}

func list(cmd *Command, c *Client) (err os.Error) {
	var arg []Args
	var v string

	str := ""
	tag1 := cmd.S("tag1", "any")
	tag2 := cmd.S("tag2", "")

	if tag2 == "" {
		str = fmt.Sprintf("list \"%s\"", tag1)
	} else {
		var term string
		if term = cmd.S("term", ""); term == "" {
			return os.NewError("Missing parameter @term if parameter @tag2 has been supplied.")
		}
		str = fmt.Sprintf("list \"%s\" \"%s\" \"%s\"", tag1, tag2, term)
	}

	if arg, err = c.requestListArgs(str); err != nil {
		return
	}

	for _, a := range arg {
		for _, v = range a {
			fmt.Printf("%s\n", v)
		}
	}
	return
}

func listall(cmd *Command, c *Client) (err os.Error) {
	var arg []Args
	var v string

	str := "listall"
	path := cmd.S("path", "")
	if path != "" {
		str = fmt.Sprintf("listall \"%s\"", path)
	}

	if arg, err = c.requestListArgs(str); err != nil {
		return
	}

	for _, a := range arg {
		for _, v = range a {
			fmt.Printf("%s\n", v)
		}
	}
	return
}

func listallinfo(cmd *Command, c *Client) (err os.Error) {
	path := cmd.S("path", "")
	if path == "" {
		return c.requestList("listallinfo")
	}
	return c.requestList("listallinfo \"%s\"", path)
}

func lsinfo(cmd *Command, c *Client) (err os.Error) {
	var arg []Args
	var k, v string

	str := "lsinfo"
	path := cmd.S("path", "")
	if path != "" {
		str = fmt.Sprintf("lsinfo \"%s\"", path)
	}

	if arg, err = c.requestListArgs(str); err != nil {
		return
	}

	for _, a := range arg {
		for k, v = range a {
			if k != "directory" && k != "file" {
				continue
			}
			fmt.Printf("%s\n", v)
		}
	}
	return
}

func search(cmd *Command, c *Client) (err os.Error) {
	return c.requestList(
		"search \"%s\" \"%s\"",
		cmd.S("tag", "any"),
		cmd.S("term", ""),
	)
}

func count(cmd *Command, c *Client) (err os.Error) {
	return c.requestList(
		"count \"%s\" \"%s\"",
		cmd.S("tag", "any"),
		cmd.S("term", ""),
	)
}
