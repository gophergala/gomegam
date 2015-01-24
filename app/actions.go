package app

import (
	"github.com/megamsys/libgo/action"
	"github.com/gophergala/gomegam/global"
        "errors"
	log "code.google.com/p/log4go"
)

var analyticsAction = action.Action{
	Name: "analyticsAction",
	Forward: func(ctx action.FWContext) (action.Result, error) {
		var app global.App
		switch ctx.Params[0].(type) {
		case global.App:
			app = ctx.Params[0].(global.App)
		case *global.App:
			app = *ctx.Params[0].(*global.App)
		default:
			return nil, errors.New("First parameter must be App or *global.App.")
		}
	 	return &app, nil
	},
	Backward: func(ctx action.BWContext) {
		log.Info("[%s] Nothing to recover")
	},
	MinParams: 1,
}


