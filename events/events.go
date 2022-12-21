package events

import "github.com/LightningDev1/discordgo"

type _Events struct {
	Session *discordgo.Session
}

func (e *_Events) Register() {
	e.Session.AddHandler(e.OnReady)
}

var Events = &_Events{}