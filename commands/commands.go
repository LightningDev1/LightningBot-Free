package commands

import (
	"github.com/LightningDev1/dgc"
	"github.com/LightningDev1/discordgo"
)

type _Commands struct {
	Session *discordgo.Session
	Router  *dgc.Router
}

func (c *_Commands) Register() {
	c.RegisterHelpCommand()
	c.RegisterTextCommands()
	c.RegisterFunCommands()
	c.RegisterImageCommands()
	c.RegisterCryptoCommands()
	c.RegisterTrollingCommands()
	c.RegisterUtilityCommands()
	c.RegisterNetworkCommands()
	c.RegisterConfigCommands()
}

var Commands = &_Commands{}
