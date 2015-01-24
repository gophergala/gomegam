package app

import (
	"encoding/json"
	"github.com/megamsys/libgo/action"
	"github.com/gophergala/gomegam/global"
	"fmt"
)

//
// this executes all actions for megam install
//
func AnalyticsProcess(app *global.App) error {
	actions := []*action.Action{&analyticsAction}

	pipeline := action.NewPipeline(actions...)
	err := pipeline.Execute(app)
	if err != nil {
		return err
	}
	return nil
}


