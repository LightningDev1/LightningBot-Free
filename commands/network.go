package commands

import (
	"fmt"

	"github.com/LightningDev1/LightningBot-Free/api"
	"github.com/LightningDev1/LightningBot-Free/embed"
	"github.com/LightningDev1/LightningBot-Free/http"
	"github.com/LightningDev1/LightningBot-Free/utils"
	"github.com/LightningDev1/dgc"
	"github.com/bitly/go-simplejson"
)

func (c *_Commands) RegisterNetworkCommands() {
	c.Router.StartCategory("Network", "Network commands")

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "ping",
		Description: "Check the latency of the bot",
		Usage:       "[p]ping",
		Handler:     ping,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "pingweb",
		Description: "Check the latency to a website",
		Usage:       "[p]pingweb <host>",
		Handler:     pingweb,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "ipinfo",
		Description: "Get info about an IP Address",
		Usage:       "[p]ipinfo <ip>",
		Handler:     ipinfo,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "discordstatus",
		Description: "Check the Discord status",
		Usage:       "[p]discordstatus",
		Handler:     discordstatus,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "open",
		Description: "Open a URL in the browser",
		Usage:       "[p]open <url>",
		Handler:     open,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "interface",
		Description: "See your public/local ip and MAC address",
		Usage:       "[p]interface",
		Handler:     _interface,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "paste",
		Description: "Create a paste",
		Usage:       "[p]pastebin <text>",
		Handler:     pastebin,
	})
}

func ping(ctx *dgc.Ctx) {
	emb := embed.NewEmbed()
	time, err := api.Ping("lightning-bot.com")
	if err != nil {
		emb.AddField("Ping to LightningBot", "Error")
	} else {
		emb.AddField("Ping to LightningBot", fmt.Sprintf("%dms", time.Milliseconds()))
	}
	time, err = api.Ping("google.com")
	if err != nil {
		emb.AddField("Ping to Google", "Error")
	} else {
		emb.AddField("Ping to Google", fmt.Sprintf("%dms", time.Milliseconds()))
	}
	emb.
		AddField("Ping to Discord", fmt.Sprintf("%dms", ctx.Session.HeartbeatLatency().Milliseconds())).
		Send(ctx)
}

func pingweb(ctx *dgc.Ctx) {
	host := ctx.Arguments.Raw()
	time, err := api.Ping(host)
	emb := embed.NewEmbed()
	if err != nil {
		emb.AddField("Ping to "+host, "Error")
	} else {
		emb.AddField("Ping to "+host, fmt.Sprintf("%dms", time.Milliseconds()))
	}
	emb.Send(ctx)
}

func ipinfo(ctx *dgc.Ctx) {
	ip := ctx.Arguments.Raw()
	httpResponse := http.Get(fmt.Sprintf("https://ipinfo.io/%s/json", ip))
	if httpResponse.Error != nil {
		_ = ctx.RespondText("An error has occurred.")
		return
	}

	ipinfo, err := simplejson.NewJson(httpResponse.BodyBytes)
	if err != nil {
		_ = ctx.RespondText("An error has occurred.")
		return
	}

	if ipinfo.Get("hostname").MustInt(0) == 404 {
		_ = ctx.RespondText("No results found.")
		return
	}

	embed.NewEmbed().
		AddField("IP", ip).
		AddField("Hostname", ipinfo.Get("hostname").MustString()).
		AddField("City", ipinfo.Get("city").MustString()).
		AddField("Region", ipinfo.Get("region").MustString()).
		AddField("Country", ipinfo.Get("country").MustString()).
		AddField("Location", ipinfo.Get("loc").MustString()).
		Send(ctx)
}

func discordstatus(ctx *dgc.Ctx) {
	httpResponse := http.Get("https://srhpyqt94yxb.statuspage.io/api/v2/components.json")
	if httpResponse.Error != nil {
		_ = ctx.RespondText("An error has occurred.")
		return
	}

	status, err := simplejson.NewJson(httpResponse.BodyBytes)
	if err != nil {
		_ = ctx.RespondText("An error has occurred.")
		return
	}

	emb := embed.NewEmbed()
	components := status.Get("components")
	for i := 0; i < len(components.MustArray()); i++ {
		component := components.GetIndex(i)
		emb.AddField(component.Get("name").MustString(), component.Get("status").MustString())
	}
	emb.Send(ctx)
}

func open(ctx *dgc.Ctx) {
	_ = utils.Misc.OpenURL(ctx.Arguments.Raw())
}

func _interface(ctx *dgc.Ctx) {
	embed.NewEmbed().
		AddField("Public IP", api.GetPublicIP()).
		AddField("Local IP", api.GetLocalIP()).
		AddField("MAC Address", api.GetMacAddress()).
		Send(ctx)
}

func pastebin(ctx *dgc.Ctx) {
	paste, err := api.CreatePaste(ctx.Arguments.Raw())
	if err != nil {
		_ = ctx.RespondText("An error has occurred.")
		return
	}
	embed.NewEmbed().
		AddField("Paste URL", paste).
		Send(ctx)
}
