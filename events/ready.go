package events

import (
	"github.com/LightningDev1/LB-Selfbot-Free/utils"
	"github.com/LightningDev1/discordgo"
)

func (e *_Events) OnReady(_ *discordgo.Session, event *discordgo.Ready) {
	utils.Logging.Info("Ready!", event.User.Email)
}
