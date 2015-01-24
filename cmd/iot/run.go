package main

import (
	"github.com/megamsys/libgo/cmd"
	"launchpad.net/gnuflag"
)

type CIBStart struct {
	fs *gnuflag.FlagSet
}

func (g *CIBStart) Info() *cmd.Info {
	desc := `starts the cib base web daemon.
	
	
	`
	return &cmd.Info{
		Name:    "start",
		Usage:   `start`,
		Desc:    desc,
		MinArgs: 0,
	}
}

func (c *CIBStart) Run(context *cmd.Context, client *cmd.Client) error {
	RunWeb()
	return nil
}

func (c *CIBStart) Flags() *gnuflag.FlagSet {
	if c.fs == nil {
		c.fs = gnuflag.NewFlagSet("cib", gnuflag.ExitOnError)
	}
	return c.fs
}

