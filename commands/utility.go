package commands

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/LightningDev1/LB-Selfbot-Free/api"
	"github.com/LightningDev1/LB-Selfbot-Free/constants"
	"github.com/LightningDev1/LB-Selfbot-Free/embed"
	"github.com/LightningDev1/LB-Selfbot-Free/utils"
	"github.com/LightningDev1/dgc"
)

func (c *_Commands) RegisterUtilityCommands() {
	c.Router.StartCategory("Utility", "Utility commands")

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "tts",
		Description: "Send a Text-To-Speech message",
		Usage:       "[p]tts <language> <text>",
		Handler:     tts,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "translate",
		Description: "Translate text to a language",
		Usage:       "[p]translate <language> <text>",
		Handler:     translate,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "detectlang",
		Description: "Detect a texts language",
		Usage:       "[p]detectlang <text>",
		Handler:     detectlang,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "english",
		Description: "Translate text to English",
		Usage:       "[p]english <text>",
		Handler:     english,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "website",
		Description: "Open the LightningBot website",
		Usage:       "[p]website",
		Handler:     website,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "dashboard",
		Description: "Open the LightningBot dashboard",
		Usage:       "[p]dashboard",
		Handler:     dashboard,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "calc",
		Description: "Open the calculator",
		Usage:       "[p]calc",
		Handler:     calc,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "notepad",
		Description: "Open notepad",
		Usage:       "[p]notepad",
		Handler:     notepad,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "theme",
		Description: "Change your Discord theme",
		Usage:       "[p]theme <light/dark>",
		Handler:     theme,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "compact",
		Description: "Toggle Discord compact mode",
		Usage:       "[p]compact <on/off>",
		Handler:     compact,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "devmode",
		Description: "Toggle Discord developer mode",
		Usage:       "[p]devmode <on/off>",
		Handler:     devmode,
	})
}

func tts(ctx *dgc.Ctx) {
	ttsFile, err := api.TTS(utils.String.RemoveWords(ctx.Arguments.Raw(), 1), ctx.Arguments.Get(0).Raw())
	if err != nil {
		_ = ctx.RespondText("An error has occurred.")
		return
	}

	if strings.Contains(string(ttsFile), "Bad Request") {
		_ = ctx.RespondText("Invalid language. Please use an ISO 639-1 language code. See https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes for more information.")
		return
	}

	_, _ = ctx.Session.ChannelFileSend(ctx.Event.ChannelID, "tts.mp3", bytes.NewReader(ttsFile))
}

func translate(ctx *dgc.Ctx) {
	translateResult := api.Translate(utils.String.RemoveWords(ctx.Arguments.Raw(), 1), ctx.Arguments.Get(0).Raw())
	if translateResult.Error != nil {
		_ = ctx.RespondText("An error has occurred: " + translateResult.Error.Error())
		return
	}

	embed.NewEmbed().
		AddField("Source Language", utils.String.Capitalize(constants.LANGUAGES[translateResult.SourceLanguage])).
		AddField("Destination Language", utils.String.Capitalize(constants.LANGUAGES[translateResult.DestinationLanguage])).
		AddField("Original Text", translateResult.Original).
		AddField("Translated Text", translateResult.Result).
		Send(ctx)
}

func detectlang(ctx *dgc.Ctx) {
	translateResult := api.Translate(ctx.Arguments.Raw(), "en")
	if translateResult.Error != nil {
		_ = ctx.RespondText("An error has occurred: " + translateResult.Error.Error())
		return
	}

	embed.NewEmbed().
		AddField("Text", translateResult.Original).
		AddField("Detected Language", utils.String.Capitalize(constants.LANGUAGES[translateResult.SourceLanguage])).
		AddField("Confidence", fmt.Sprintf("%.2f", translateResult.Confidence)).
		Send(ctx)
}

func english(ctx *dgc.Ctx) {
	translateResult := api.Translate(ctx.Arguments.Raw(), "en")
	if translateResult.Error != nil {
		_ = ctx.RespondText("An error has occurred: " + translateResult.Error.Error())
		return
	}
	
	embed.NewEmbed().
		AddField("Source Language", utils.String.Capitalize(constants.LANGUAGES[translateResult.SourceLanguage])).
		AddField("Original Text", translateResult.Original).
		AddField("Translated Text", translateResult.Result).
		Send(ctx)
}

func website(*dgc.Ctx) {
	_ = utils.Misc.OpenURL("https://lightning-bot.com/")
}

func dashboard(*dgc.Ctx) {
	_ = utils.Misc.OpenURL("https://lightning-bot.com/dashboard")
}

func calc(ctx *dgc.Ctx) {
	if runtime.GOOS != "windows" {
		_ = ctx.RespondText("This command is only available on Windows")
	}
	_ = exec.Command("C:\\Windows\\System32\\calc.exe").Start()
}

func notepad(ctx *dgc.Ctx) {
	if runtime.GOOS != "windows" {
		_ = ctx.RespondText("This command is only available on Windows")
	}
	_ = exec.Command("C:\\Windows\\System32\\notepad.exe").Start()
}

func theme(ctx *dgc.Ctx) {
	theme := strings.ToLower(ctx.Arguments.Get(0).Raw())
	if theme != "light" && theme != "dark" {
		_ = ctx.RespondText("Error: <light/dark> must be either \"light\" or \"dark\"")
		return
	}

	_ = api.ChangeSettings(map[string]interface{}{"theme": theme}, true)
}

func compact(ctx *dgc.Ctx) {
	toggle := strings.ToLower(ctx.Arguments.Get(0).Raw())
	if toggle != "on" && toggle != "off" {
		_ = ctx.RespondText("Error: <on/off> must be either \"on\" or \"off\"")
		return
	}

	_ = api.ChangeSettings(map[string]interface{}{"message_display_compact": toggle == "on"}, true)
}

func devmode(ctx *dgc.Ctx) {
	toggle := strings.ToLower(ctx.Arguments.Get(0).Raw())
	if toggle != "on" && toggle != "off" {
		_ = ctx.RespondText("Error: <on/off> must be either \"on\" or \"off\"")
		return
	}

	_ = api.ChangeSettings(map[string]interface{}{"developer_mode": toggle == "on"}, true)
}
