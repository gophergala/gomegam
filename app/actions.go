package app

import (
	"os"
	"github.com/megamsys/libgo/action"
	"github.com/gophergala/gomegam/global"
        "errors"
	log "code.google.com/p/log4go"
	"github.com/sendgrid/sendgrid-go"
	"github.com/tsuru/config"
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
		sendgrid()
	 	return &app, nil
	},
	Backward: func(ctx action.BWContext) {
		log.Info("[%s] Nothing to recover")
	},
	MinParams: 1,
}


func sendgrid() {
	
    sg := sendgrid.NewSendGridClient(os.Getenv("SD_UNAME"), os.Getenv("SD_KEY") )
    message := sendgrid.NewMail()
    email, kerr := config.GetString("addto:email")
	if kerr != nil {
		return "", kerr
	}
	
	
    message.AddTo(email)
    
    name, kerr := config.GetString("addto:name")
	if kerr != nil {
		return "", kerr
	}
    message.AddToName(name)
    message.SetSubject("GoMegam-IoT ")
    message.SetText("Welcome to Gomegam IoT Project")
    message.SetFrom("getyesh@megam.co.in")
    if r := sg.Send(message); r == nil {
        fmt.Println("Email sent!")
    } else {
        fmt.Println(r)
    }
}


