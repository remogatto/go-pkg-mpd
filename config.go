// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import (
	"os"
	"strconv"
)

type Config struct {
	Address  string
	Port     int
	Password string
}

func NewConfig() *Config {
	c := new(Config)
	c.Address = "127.0.0.1"
	c.Port = 6600
	c.Password = ""

	if v := os.Getenv("MPD_HOST"); len(v) > 0 {
		c.Address = v
	}

	if v := os.Getenv("MPD_PORT"); len(v) > 0 {
		if p, err := strconv.Atoi(v); err != nil {
			c.Port = p
		}
	}

	if v := os.Getenv("MPD_PASSWORD"); len(v) > 0 {
		c.Password = v
	}

	return c
}
