// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import (
	"os"
	"fmt"
)

func add(cmd *Command, c *Client) (err os.Error) {
	return c.request("add \"%s\"", cmd.S("path", ""))
}

func addid(cmd *Command, c *Client) (err os.Error) {
	pos := cmd.I("pos", -1)
	if pos > -1 {
		return c.request("addid \"%s\" %d", cmd.S("path", ""), pos)
	}
	return c.request("addid \"%s\"")
}

func clear(cmd *Command, c *Client) (err os.Error) {
	return c.request("clear")
}

func current(cmd *Command, c *Client) (err os.Error) {
	var a, stats Args
	if a, err = c.requestArgs("currentsong"); err != nil {
		return
	}

	if stats, err = c.requestArgs("stats"); err != nil {
		return
	}

	fmt.Printf("[%s/%s] %s - %s - %s (%s)\n",
		a["Pos"], stats["songs"], a["Artist"], a["Album"], a["Title"],
		parseTime(a.Int("Time", 0)),
	)
	return
}

func delete(cmd *Command, c *Client) (err os.Error) {
	return c.request("delete %d", cmd.I("pos", 0))
}

func deleteid(cmd *Command, c *Client) (err os.Error) {
	return c.request("deleteid %d", cmd.I("id", 0))
}

func load(cmd *Command, c *Client) (err os.Error) {
	return c.request("load \"%s\"", cmd.S("name", ""))
}

func rename(cmd *Command, c *Client) (err os.Error) {
	return c.request("rename \"%s\" \"%s\"",
		cmd.S("oldname", ""),
		cmd.S("newname", ""),
	)
}

func move(cmd *Command, c *Client) (err os.Error) {
	return c.request("move %d %d",
		cmd.I("src", 0),
		cmd.I("dest", 0),
	)
}

func moveid(cmd *Command, c *Client) (err os.Error) {
	return c.request("moveid %d %d",
		cmd.I("src", 0),
		cmd.I("dest", 0),
	)
}

func plinfo(cmd *Command, c *Client) (err os.Error) {
	pos := cmd.I("pos", -1)
	if pos == -1 {
		return c.requestList("playlistinfo")
	}
	return c.requestList("playlistinfo %d", pos)
}

func plchanges(cmd *Command, c *Client) (err os.Error) {
	return c.requestList("plchanges %d", cmd.I("version", 0))
}

func plchangesid(cmd *Command, c *Client) (err os.Error) {
	var arg []Args
	if arg, err = c.requestListArgs("plchangesposid %d", cmd.I("version", 0)); err != nil {
		return
	}

	for _, a := range arg {
		fmt.Printf("id: %s, cpos: %s\n", a["Id"], a["cpos"])
	}
	return
}

func rm(cmd *Command, c *Client) (err os.Error) {
	return c.request("rm \"%s\"", cmd.S("name", ""))
}

func save(cmd *Command, c *Client) (err os.Error) {
	return c.request("save \"%s\"", cmd.S("name", ""))
}

func shuffle(cmd *Command, c *Client) (err os.Error) {
	return c.request("shuffle")
}

func swap(cmd *Command, c *Client) (err os.Error) {
	return c.request("swap %d %d",
		cmd.I("pos1", 0),
		cmd.I("pos2", 0),
	)
}

func swapid(cmd *Command, c *Client) (err os.Error) {
	return c.request("swapid %d %d",
		cmd.I("id1", 0),
		cmd.I("id2", 0),
	)
}

func listpl(cmd *Command, c *Client) (err os.Error) {
	var arg []Args
	if arg, err = c.requestListArgs("listplaylist \"%s\"", cmd.S("name", "")); err != nil {
		return
	}

	var v string
	for _, a := range arg {
		for _, v = range a {
			fmt.Printf("%s\n", v)
		}
	}
	return
}

func listplinfo(cmd *Command, c *Client) (err os.Error) {
	return c.requestList("listplaylistinfo \"%s\"", cmd.S("name", ""))
}

func pladd(cmd *Command, c *Client) (err os.Error) {
	return c.request("playlistadd \"%s\" \"%s\"",
		cmd.S("name", ""),
		cmd.S("path", ""),
	)
}

func plclear(cmd *Command, c *Client) (err os.Error) {
	return c.request("playlistclear \"%s\"", cmd.S("name", ""))
}

func pldelete(cmd *Command, c *Client) (err os.Error) {
	return c.request("playlistdelete \"%s\" %d",
		cmd.S("name", ""),
		cmd.I("id", 0),
	)
}

func plmove(cmd *Command, c *Client) (err os.Error) {
	return c.request("playlistmove \"%s\" %d %d",
		cmd.S("name", ""),
		cmd.I("id", 0),
		cmd.I("pos", 0),
	)
}

func plsearch(cmd *Command, c *Client) (err os.Error) {
	var arg []Args
	if arg, err = c.requestListArgs("playlistsearch \"%s\" \"%s\"",
		cmd.S("tag", ""),
		cmd.S("term", "")); err != nil {
		return
	}

	for _, a := range arg {
		fmt.Printf("[%6s:%6s] %s - %s - %s (%s)\n",
			a["Pos"], a["Id"], a["Artist"], a["Album"], a["Title"], parseTime(a.Int("Time", 0)),
		)
	}
	return
}
