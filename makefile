# Copyright (c) 2010, Jim Teeuwen. All rights reserved.
# This code is subject to a 1-clause BSD license.
# See the LICENSE file for its contents.

include $(GOROOT)/src/Make.inc

TARG = github.com/jteeuwen/go-pkg-mpd
GOFILES = config.go patterns.go command.go param.go api_admin.go api_info.go \
	api_database.go api_playlist.go api_playback.go client.go args.go http.go \
	misc.go

include $(GOROOT)/src/Make.pkg
