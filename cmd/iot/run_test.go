package main

import (
	"github.com/megamsys/libgo/cmd"
	"launchpad.net/gocheck"
)

func (s *S) TestCIBStartInfo(c *gocheck.C) {
	desc := `starts the cib base web daemon.
	
	
	`
	expected := &cmd.Info{
		Name:    "start",
		Usage:   `start`,
		Desc:    desc,
		MinArgs: 0,
	}
	command := CIBStart{}
	c.Assert(command.Info(), gocheck.DeepEquals, expected)
}