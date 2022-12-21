package commands

import (
	"github.com/LightningDev1/LB-Selfbot-Free/config"
	"github.com/LightningDev1/dgc"
)

func (c *_Commands) RegisterConfigCommands() {
	c.Router.StartCategory("Config", "Config commands")

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "prefix",
		Description: "Change the command prefix",
		Usage:       "[p]prefix <text>",
		Handler:     prefix,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "title",
		Description: "Change the title",
		Usage:       "[p]title <text>",
		Handler:     title,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "footer",
		Description: "Change the footer",
		Usage:       "[p]footer <text>",
		Handler:     footer,
	})
}

func prefix(ctx *dgc.Ctx) {
	cfg, err := config.Load()
	if err != nil {
		ctx.RespondText("Error loading config")
		return
	}

	cfg.CommandPrefix = ctx.Arguments.Raw()

	err = cfg.Save()
	if err != nil {
		ctx.RespondText("Error saving config")
		return
	}
}

func title(ctx *dgc.Ctx) {
	cfg, err := config.Load()
	if err != nil {
		ctx.RespondText("Error loading config")
		return
	}

	cfg.Embed.Title = ctx.Arguments.Raw()

	err = cfg.Save()
	if err != nil {
		ctx.RespondText("Error saving config")
		return
	}
}

func footer(ctx *dgc.Ctx) {
	cfg, err := config.Load()
	if err != nil {
		ctx.RespondText("Error loading config")
		return
	}

	cfg.Embed.Footer = ctx.Arguments.Raw()

	err = cfg.Save()
	if err != nil {
		ctx.RespondText("Error saving config")
		return
	}
}
