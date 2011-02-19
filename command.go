// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import (
	"fmt"
	"os"
	"bytes"
)

type Command struct {
	Name   string
	Desc   string
	Params []*Param
	Exec   func(cmd *Command, c *Client) os.Error
}

// Get parameter as string
func (this *Command) S(name, def string) string {
	for _, v := range this.Params {
		if v.Name == name {
			return v.Value
		}
	}
	return def
}

// Get parameter as int
func (this *Command) I(name string, def int) int {
	for _, v := range this.Params {
		if v.Name == name {
			return v.Int()
		}
	}
	return def
}

// Get parameter as int64
func (this *Command) I64(name string, def int64) int64 {
	for _, v := range this.Params {
		if v.Name == name {
			return v.Int64()
		}
	}
	return def
}

func (this *Command) String() string {
	var d []byte
	buf := bytes.NewBuffer(d)
	buf.WriteString(this.Name)

	for _, v := range this.Params {
		buf.WriteByte(' ')
		buf.WriteString(v.String())
	}

	return buf.String()
}

func (this *Command) Run(cfg *Config, data []string) (err os.Error) {
	rpcount := 0
	for _, v := range this.Params {
		if !v.Optional {
			rpcount++
		}
	}

	if len(data)-1 < rpcount {
		err = os.NewError(fmt.Sprintf("Missing parameters for command '%s'.", this.Name))
		return
	}

	for i := 1; i < len(data); i++ {
		if i > len(this.Params) {
			break
		}

		if !this.Params[i-1].IsValid(data[i]) {
			err = os.NewError(fmt.Sprintf(
				"Invalid value '%s' for parameter '%s'.",
				data[i], this.Params[i-1],
			))
			return
		}

		this.Params[i-1].Value = data[i]
	}

	client := newClient()
	if err = client.Open("tcp", fmt.Sprintf("%s:%d", cfg.Address, cfg.Port)); err != nil {
		return
	}

	defer client.Close()

	if len(cfg.Password) > 0 {
		if err = client.request("password \"%s\"", cfg.Password); err != nil {
			return
		}
	}

	return this.Exec(this, client)
}

func CommandList() []string {
	return []string{
		"disableoutput", "enableoutput", "kill", "update", "status", "simplestatus",
		"stats", "outputs", "commands", "notcommands", "tagtypes", "urlhandlers", "find",
		"list", "listall", "listallinfo", "lsinfo", "search", "count", "add", "addid",
		"clear", "current", "delete", "deleteid", "load", "rename", "move", "moveid",
		"plinfo", "plchanges", "plchangesid", "rm", "save", "shuffle", "swap", "swapid",
		"listpl", "listplinfo", "pladd", "plclear", "pldelete", "plmove", "plsearch",
		"crossfade", "next", "pause", "play", "playid", "previous", "random", "repeat",
		"seek", "seekid", "volume", "stop", "toggle",
	}
}

func CreateCommand(name string) *Command {
	cmd := &Command{Name: name, Params: make([]*Param, 0)}

	switch name {
	/* Admin commands */
	case "disableoutput":
		cmd.Desc = "Turns an audio-output source off."
		cmd.Params = []*Param{
			newParam("id", "Id of the output device. Use the 'outputs' command to find all valid Ids.", PatInteger, false),
		}
		cmd.Exec = disableoutput
	case "enableoutput":
		cmd.Desc = "Turns an audio-output source on."
		cmd.Params = []*Param{
			newParam("id", "Id of the output device. Use the 'outputs' command to find all valid Ids.", PatInteger, false),
		}
		cmd.Exec = enableoutput
	case "kill":
		cmd.Desc = "Stops MPD from running, in a safe way. Writes a state file if defined."
		cmd.Exec = kill

	/* Informational commands */
	case "update":
		cmd.Desc = "Scans the music directory as defined in the MPD configuration file's music_directory  setting. Adds new files and their metadata (if any) to the MPD database and removes files and metadata from the database that are no longer in the directory."
		cmd.Params = []*Param{
			newParam("path", "path is an optional argument that picks an exact directory or file to update, otherwise the root of the music_directory in your MPD configuration file is assumed.", PatAny, true),
		}
		cmd.Exec = update
	case "status":
		cmd.Desc = "Reports the current status of MPD, as well as the current settings of some playback options."
		cmd.Exec = status
	case "simplestatus":
		cmd.Desc = "Same as status, but only basic info in 'prettier' output.."
		cmd.Exec = simplestatus
	case "stats":
		cmd.Desc = "Reports database and playlist statistics."
		cmd.Exec = stats
	case "outputs":
		cmd.Desc = "Reports information about all known audio output devices."
		cmd.Exec = outputs
	case "commands":
		cmd.Desc = "Reports which commands the current user has access to."
		cmd.Exec = commands
	case "notcommands":
		cmd.Desc = "Reports which commands the current user has *no* access to."
		cmd.Exec = notcommands
	case "tagtypes":
		cmd.Desc = "Reports a list of available song metadata fields."
		cmd.Exec = tagtypes
	case "urlhandlers":
		cmd.Desc = "Reports a list of available URL handlers."
		cmd.Exec = urlhandlers

	/* Database commands */
	case "find":
		cmd.Desc = "Finds songs in the database with a case sensitive, exact match to @term."
		cmd.Params = []*Param{
			newParam("tag", "This is the type of metadata you wish to use to refine the search. Examples would be album, artist , title or any.", PatType, false),
			newParam("term", "This is the value that is being searched for in @tag.", PatAny, false),
		}
		cmd.Exec = find
	case "list":
		cmd.Desc = "Reports all metadata of @type1."
		cmd.Params = []*Param{
			newParam("tag1", "This lists all metadata of @tag1", PatType, false),
			newParam("tag2", "Only required if @term is present. This specifies to look for @tag2 in the list of @tag1", PatType, true),
			newParam("term", "Only required if @tag2 is present. This specifies to look for matches of @term of @tag2 in the list of @tag1", PatAny, true),
		}
		cmd.Exec = list
	case "listall":
		cmd.Desc = "Reports all directories and filenames in @path recursively."
		cmd.Params = []*Param{
			newParam("path", "An optional path or directory to act as the root of the list.", PatAny, true),
		}
		cmd.Exec = listall
	case "listallinfo":
		cmd.Desc = "Reports all information in database about all music files in <string path> recursively."
		cmd.Params = []*Param{
			newParam("path", "An optional path or directory to act as the root of the list.", PatAny, true),
		}
		cmd.Exec = listallinfo
	case "lsinfo":
		cmd.Desc = "Reports contents of @path, from the database."
		cmd.Params = []*Param{
			newParam("path", "An optional path or directory to act as the root of the list.", PatAny, true),
		}
		cmd.Exec = lsinfo
	case "search":
		cmd.Desc = "Finds songs in the database with a case insensitive match to @what."
		cmd.Params = []*Param{
			newParam("tag", "This is the type of metadata you wish to use to refine the search. Examples would be album, artist , title or any.", PatType, false),
			newParam("term", "This is the value that is being searched for in @type.", PatAny, false),
		}
		cmd.Exec = search
	case "count":
		cmd.Desc = "Reports the number of songs and their total playtime in the database matching @what."
		cmd.Params = []*Param{
			newParam("tag", "This is the type of metadata you wish to use to refine the search. Examples would be album, artist , title or any.", PatType, false),
			newParam("term", "This is the value that is being searched for in @type.", PatAny, false),
		}
		cmd.Exec = count

	/* Playlist commands */
	case "add":
		cmd.Desc = "Add a single file from the database to the playlist. This command increments the playlist version by 1 for each song added to the playlist."
		cmd.Params = []*Param{
			newParam("path", "A single directory or file. If this is a directory or path all files in directory or path are added recursively. Adding all files in the database is as simple as add /.", PatAny, false),
		}
		cmd.Exec = add
	case "addid":
		cmd.Desc = "Same as 'add', but this returns a playlistid and allows specifying a position at which to insert the file(s)."
		cmd.Params = []*Param{
			newParam("path", "A single directory or file. If this is a directory or path all files in directory or path are added recursively. Adding all files in the database is as simple as add /.", PatAny, false),
			newParam("pos", "Optional integer value specifying the location at which to insert the file(s) into the playlist.", PatInteger, true),
		}
		cmd.Exec = addid
	case "clear":
		cmd.Desc = "Clears the current playlist. Increments the playlist version by 1."
		cmd.Exec = clear
	case "current":
		cmd.Desc = "Reports the metadata of the currently playing song."
		cmd.Exec = current
	case "delete":
		cmd.Desc = "Deletes the specified song from the playlist. increments the playlist version by 1."
		cmd.Params = []*Param{
			newParam("pos", "Position of the song in the playlist.", PatInteger, false),
		}
		cmd.Exec = delete
	case "deleteid":
		cmd.Desc = "Deletes the specified song from the playlist. Increments the playlist version by 1."
		cmd.Params = []*Param{
			newParam("id", "Id of the song to delete.", PatInteger, false),
		}
		cmd.Exec = deleteid
	case "load":
		cmd.Desc = "Load the playlist @name from the playlist directory, Increments the playlist version by the number of songs added."
		cmd.Params = []*Param{
			newParam("name", "Name of the playlist file *without* the file extension.", PatAny, false),
		}
		cmd.Exec = load
	case "rename":
		cmd.Desc = "Renames a playlist from @oldname to @newname."
		cmd.Params = []*Param{
			newParam("oldname", "Current name of the playlist.", PatAny, false),
			newParam("newname", "New name of the playlist.", PatAny, false),
		}
		cmd.Exec = rename
	case "move":
		cmd.Desc = "Moves a song from position @src to position @dest."
		cmd.Params = []*Param{
			newParam("src", "Source position.", PatInteger, false),
			newParam("dest", "Target position.", PatInteger, false),
		}
		cmd.Exec = move
	case "moveid":
		cmd.Desc = "Moves a song with id @src to position @dest."
		cmd.Params = []*Param{
			newParam("src", "Song Id of the track you want to move.", PatInteger, false),
			newParam("dest", "Target position.", PatInteger, false),
		}
		cmd.Exec = moveid
	case "plinfo":
		cmd.Desc = "Reports metadata for songs in the playlist."
		cmd.Params = []*Param{
			newParam("pos", "An optional number that specifies a single song to display information for.", PatInteger, true),
		}
		cmd.Exec = plinfo
	case "plchanges":
		cmd.Desc = "Reports changed songs currently in the playlist since @version."
		cmd.Params = []*Param{
			newParam("version", "The number for the version to display changed songs of the playlist.", PatInteger, false),
		}
		cmd.Exec = plchanges
	case "plchangesid":
		cmd.Desc = "Same as plchanges, but returns only the songids."
		cmd.Params = []*Param{
			newParam("version", "The number for the version to display changed songs of the playlist.", PatInteger, false),
		}
		cmd.Exec = plchangesid
	case "rm":
		cmd.Desc = "Removes the playlist called @name from the playlist directory."
		cmd.Params = []*Param{
			newParam("name", "The name of the saved playlist to be removed from the playlist directory.", PatAny, false),
		}
		cmd.Exec = rm
	case "save":
		cmd.Desc = "Saves the current playlist to @name in the playlist directory."
		cmd.Params = []*Param{
			newParam("name", "The name for the saved playlist.", PatAny, false),
		}
		cmd.Exec = save
	case "shuffle":
		cmd.Desc = "Shuffles the current playlist, increments playlist version by 1."
		cmd.Exec = shuffle
	case "swap":
		cmd.Desc = "Swap positions of songs at positions @pos1 and @pos2. Increments playlist version by 1."
		cmd.Params = []*Param{
			newParam("pos1", "First song position", PatInteger, false),
			newParam("pos2", "Second song position", PatInteger, false),
		}
		cmd.Exec = swap
	case "swapid":
		cmd.Desc = "Swap positions of songs with id @pos1 and @pos2. Increments playlist version by 1."
		cmd.Params = []*Param{
			newParam("id1", "First song id", PatInteger, false),
			newParam("id2", "Second song id", PatInteger, false),
		}
		cmd.Exec = swapid
	case "listpl":
		cmd.Desc = "Reports files in playlist named @name."
		cmd.Params = []*Param{
			newParam("name", "Name of the playlist", PatAny, false),
		}
		cmd.Exec = listpl
	case "listplinfo":
		cmd.Desc = "Reports songs in playlist named @name."
		cmd.Params = []*Param{
			newParam("name", "Name of the playlist", PatAny, false),
		}
		cmd.Exec = listplinfo
	case "pladd":
		cmd.Desc = "Adds @path to playlist @name."
		cmd.Params = []*Param{
			newParam("name", "Name of playlist.", PatAny, false),
			newParam("path", "Path of file to add to playlist.", PatAny, false),
		}
		cmd.Exec = pladd
	case "plclear":
		cmd.Desc = "Clears playlist @name."
		cmd.Params = []*Param{
			newParam("name", "Name of playlist.", PatAny, false),
		}
		cmd.Exec = plclear
	case "pldelete":
		cmd.Desc = "Deletes song with given @id from playlist @name."
		cmd.Params = []*Param{
			newParam("name", "Name of playlist.", PatAny, false),
			newParam("id", "Id of song to delete.", PatInteger, false),
		}
		cmd.Exec = pldelete
	case "plmove":
		cmd.Desc = "Moves song with given @id in playlist @name to position @pos."
		cmd.Params = []*Param{
			newParam("name", "Name of playlist.", PatAny, false),
			newParam("id", "Id of song to move.", PatInteger, false),
			newParam("pos", "New song position.", PatInteger, false),
		}
		cmd.Exec = plmove
	case "plsearch":
		cmd.Desc = "Case-insensitive playlist search with 'pretty' output. Easier to use when looking for specific songs to play. Outputs a list of entries like: [#pos:#id] Artist - Album - Title (mm:ss). Listed #pos and #id can be used directly with the 'play' and 'playid' commands."
		cmd.Params = []*Param{
			newParam("tag", "Metadata field to search in.", PatType, false),
			newParam("term", "Term to search for.", PatAny, false),
		}
		cmd.Exec = plsearch

	/* Playback commands */
	case "crossfade":
		cmd.Desc = "Sets crossfading (mixing) between songs."
		cmd.Params = []*Param{
			newParam("time", "Crossfade time in seconds.", PatInteger, false),
		}
		cmd.Exec = crossfade
	case "next":
		cmd.Desc = "Skip to next song."
		cmd.Exec = next
	case "pause":
		cmd.Desc = "Toggle pause on/off"
		cmd.Params = []*Param{
			newParam("toggle", "Whether to pause or resume.", PatOnOff, false),
		}
		cmd.Exec = pause
	case "play":
		cmd.Desc = "Play the song at the specified position."
		cmd.Params = []*Param{
			newParam("pos", "Position of song to play.", PatInteger, false),
		}
		cmd.Exec = play
	case "playid":
		cmd.Desc = "Play the song with the specified id."
		cmd.Params = []*Param{
			newParam("id", "Id of the song to play.", PatInteger, false),
		}
		cmd.Exec = playid
	case "previous":
		cmd.Desc = "Go back to previous song."
		cmd.Exec = previous
	case "random":
		cmd.Desc = "Toggle random mode on/off"
		cmd.Params = []*Param{
			newParam("toggle", "Whether to play randomized or not.", PatOnOff, false),
		}
		cmd.Exec = random
	case "repeat":
		cmd.Desc = "Toggle repeat mode on/off"
		cmd.Params = []*Param{
			newParam("toggle", "Whether to repeat or not.", PatOnOff, false),
		}
		cmd.Exec = repeat
	case "seek":
		cmd.Desc = "Skip to specific point in time in song at position @pos."
		cmd.Params = []*Param{
			newParam("pos", "Position of song to skip to.", PatInteger, false),
			newParam("time", "Time in seconds to jump to.", PatInteger, false),
		}
		cmd.Exec = seek
	case "seekid":
		cmd.Desc = "Skip to specific point in time in song with @id."
		cmd.Params = []*Param{
			newParam("id", "Id of song to skip to.", PatInteger, false),
			newParam("time", "Time in seconds to jump to.", PatInteger, false),
		}
		cmd.Exec = seekid
	case "volume":
		cmd.Desc = "Volume adjustment. Allows setting of explicit volume value as well as a relative increase and decrease of current volume."
		cmd.Params = []*Param{
			newParam("value", "New volume value. Can be 0-100. Either used to set volume explicitely to specified value, or in conjunction with optional @sign parameter to increase/decrease current volume.", PatInteger, false),
			newParam("sign", "Optional + or - to indicate current volume should be adjusted by @value.", PatSign, true),
		}
		cmd.Exec = volume
	case "stop":
		cmd.Desc = "Stop the playback."
		cmd.Exec = stop
	case "toggle":
		cmd.Desc = "Toggles between play/pause"
		cmd.Exec = toggle

	default:
		return nil
	}

	return cmd
}
