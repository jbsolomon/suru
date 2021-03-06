package cmd

import (
	"gopkg.in/suru.v0"
	"gopkg.in/suru.v0/cui"
	"gopkg.in/suru.v0/cui/renderer"
	"gopkg.in/suru.v0/cui/state"
	"gopkg.in/suru.v0/cui/view"

	"github.com/jroimartin/gocui"
	"github.com/pkg/errors"
)

type Live struct{ DB }

// Live wants a database connection.

func (Live) Cmd(c Context) error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return errors.Wrap(err, "creating live UI")
	}

	return (&cui.State{
		DB: c.DB,
		State: state.State{Root: &renderer.Frame{
			Stringer: view.Sprintf("Suru v%s", suru.Version),
		}},
	}).Live(g)
}

func (Live) Short() string { return "Enter live mode" }
func (Live) Help() string  { return "TODO" }
