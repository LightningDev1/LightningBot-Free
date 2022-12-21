package events

import (
	"github.com/LightningDev1/dgc"
	"github.com/LightningDev1/discordgo"
)

type _Events struct {
	Session *discordgo.Session
	Router  *dgc.Router
}

func (e *_Events) Register() {
	e.Session.AddHandler(e.OnReady)
	e.Router.RegisterMiddleware(e.OnCommand)
}

var Events = &_Events{}
