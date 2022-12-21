package commands

import (
	"strconv"
	"strings"
	"time"

	"github.com/LightningDev1/LightningBot-Free/utils"
	"github.com/LightningDev1/dgc"
)

func (c *_Commands) RegisterTrollingCommands() {
	c.Router.StartCategory("Trolling", "Trolling commands")

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "purgehack",
		Description: "Purge a chat without permissions",
		Usage:       "[p]purgehack",
		Handler:     purgehack,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "ghostping",
		Description: "Ghostping a user",
		Usage:       "[p]ghostping <user>",
		Handler:     ghostping,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "triggertyping",
		Description: "Fake \"User is typing\" message",
		Usage:       "[p]triggertyping <seconds> [#channel]",
		Handler:     triggertyping,
	})
}

func purgehack(ctx *dgc.Ctx) {
	_ = ctx.RespondText("ﾠ" + strings.Repeat("\n", 1998) + "ﾠ")
}

func ghostping(ctx *dgc.Ctx) {
	_ = ctx.Session.ChannelMessageDelete(ctx.Event.ChannelID, ctx.Event.Message.ID)
}

func triggertyping(ctx *dgc.Ctx) {
	seconds, err := strconv.Atoi(ctx.Arguments.Get(0).Raw())
	if err != nil {
		_ = ctx.RespondText("Argument <seconds> must be a number.")
		return
	}

	channel, success := utils.Discord.GetChannelFromMention(ctx, ctx.Arguments.Get(1))
	if !success {
		channel, err = utils.Discord.GetChannel(ctx.Session, ctx.Event.ChannelID)
		if err != nil {
			_ = ctx.RespondText("Could not find channel")
			return
		}
	}

	utils.Discord.SleepWithTyping(ctx.Session, channel.ID, time.Duration(seconds)*time.Second)
}
