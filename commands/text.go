package commands

import (
	"strings"

	"github.com/LightningDev1/LB-Selfbot-Free/utils"
	"github.com/LightningDev1/dgc"
	"github.com/common-nighthawk/go-figure"
)

func (c *_Commands) RegisterTextCommands() {
	c.Router.StartCategory("Text", "Text commands")

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "empty",
		Description: "Sends an empty message",
		Usage:       "[p]empty",
		Handler:     empty,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "reverse",
		Description: "Reverses the text and sends it",
		Usage:       "[p]reverse <text>",
		Handler:     reverse,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "ascii",
		Description: "Sends the text as ASCII art",
		Usage:       "[p]ascii <text>",
		Handler:     ascii,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "bold",
		Description: "Sends the text in bold",
		Usage:       "[p]bold <text>",
		Handler:     bold,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "strike",
		Description: "Add strikethrough effect to text",
		Usage:       "[p]strike <text>",
		Handler:     strike,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "italic",
		Description: "Sends the text in italicized font",
		Usage:       "[p]italic <text>",
		Handler:     italic,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "spoiler",
		Description: "Sends the text as a spoiler",
		Usage:       "[p]spoiler <text>",
		Handler:     spoiler,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "shrug",
		Description: "Sends a shrug",
		Usage:       "[p]shrug <text>",
		Handler:     shrug,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "lenny",
		Description: "Sends a lenny",
		Usage:       "[p]lenny",
		Handler:     lenny,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "tableflip",
		Description: "Sends a tableflip",
		Usage:       "[p]tableflip",
		Handler:     tableflip,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "unflip",
		Description: "Sends an unflip",
		Usage:       "[p]unflip",
		Handler:     unflip,
	})
	c.Router.RegisterCmd(&dgc.Command{
		Name:        "clap",
		Description: "Replaces spaces in the text into clap emojis",
		Usage:       "[p]clap <text>",
		Handler:     clap,
	})
}

func empty(ctx *dgc.Ctx) {
	_ = ctx.RespondText("_ _")
}

func reverse(ctx *dgc.Ctx) {
	_ = ctx.RespondText(utils.String.Reverse(ctx.Arguments.Raw()))
}

func ascii(ctx *dgc.Ctx) {
	ascii := figure.NewFigure(ctx.Arguments.Raw(), "", false).String()
	_ = ctx.RespondText("```" + ascii + "```")
}

func bold(ctx *dgc.Ctx) {
	_ = ctx.RespondText("**" + ctx.Arguments.Raw() + "**")
}

func strike(ctx *dgc.Ctx) {
	_ = ctx.RespondText("~~" + ctx.Arguments.Raw() + "~~")
}

func italic(ctx *dgc.Ctx) {
	_ = ctx.RespondText("*" + ctx.Arguments.Raw() + "*")
}

func spoiler(ctx *dgc.Ctx) {
	_ = ctx.RespondText("||" + ctx.Arguments.Raw() + "||")
}

func shrug(ctx *dgc.Ctx) {
	_ = ctx.RespondText("¯\\_(ツ)_/¯")
}

func lenny(ctx *dgc.Ctx) {
	_ = ctx.RespondText("( ͡° ͜ʖ ͡°)")
}

func tableflip(ctx *dgc.Ctx) {
	_ = ctx.RespondText("(╯°□°）╯︵ ┻━┻")
}

func unflip(ctx *dgc.Ctx) {
	_ = ctx.RespondText("┬─┬ ノ( ゜-゜ノ)")
}

func clap(ctx *dgc.Ctx) {
	text := ctx.Arguments.Raw()
	clapChars := [][]string{
		{" ", "👏"},
		{"\n", "👏\n"},
	}
	for _, element := range clapChars {
		text = strings.ReplaceAll(text, element[0], element[1])
	}
	_ = ctx.RespondText(text)
}
