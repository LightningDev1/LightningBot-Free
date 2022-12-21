package commands

import (
	"fmt"
	"net/url"

	"github.com/LightningDev1/LB-Selfbot-Free/embed"
	"github.com/LightningDev1/LB-Selfbot-Free/http"
	"github.com/LightningDev1/LB-Selfbot-Free/utils"
	"github.com/LightningDev1/dgc"
	"github.com/bitly/go-simplejson"
)

func (c *_Commands) RegisterImageCommands() {
	c.Router.StartCategory("Image", "Image commands")

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "tweet",
		Description: "Tweet a message",
		Usage:       "[p]tweet <username> <text>",
		Handler:     tweet,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "meme",
		Description: "Send a random meme",
		Usage:       "[p]meme",
		Handler:     meme,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "shibe",
		Description: "Send an image of a shiba inu",
		Usage:       "[p]shibe",
		Handler:     shibe,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "catgif",
		Description: "Send a GIF of a cat",
		Usage:       "[p]catgif",
		Handler:     catgif,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "hub",
		Description: "Make a PornHub logo with custom text",
		Usage:       "[p]hub <left> <right>",
		Handler:     hub,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "pokemon",
		Description: "Send a random Pokémon image",
		Usage:       "[p]pokemon",
		Handler:     pokemon,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "wasted",
		Description: "Add a wasted overlay on the users avatar",
		Usage:       "[p]wasted [user]",
		Handler:     wasted,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "avatar",
		Description: "View a users avatar",
		Usage:       "[p]avatar [user]",
		Handler:     avatar,
	})
}

func tweet(ctx *dgc.Ctx) {
	httpResponse := http.Get(fmt.Sprintf(
		"https://nekobot.xyz/api/imagegen?type=tweet&text=%s&username=%s",
		url.QueryEscape(utils.String.RemoveWords(ctx.Arguments.Raw(), 1)),
		url.QueryEscape(ctx.Arguments.Get(0).Raw()),
	))
	if httpResponse.Error != nil {
		_ = ctx.RespondText("An error has occurred: " + httpResponse.Error.Error())
		return
	}

	tweet, err := simplejson.NewJson(httpResponse.BodyBytes)
	if err != nil {
		_ = ctx.RespondText("An error has occurred: " + err.Error())
		return
	}

	_ = ctx.RespondText(tweet.Get("message").MustString())
}

func meme(ctx *dgc.Ctx) {
	httpResponse := http.Get("https://api.lightning-bot.com/api/v3/misc/meme")
	if httpResponse.Error != nil {
		_ = ctx.RespondText("An error has occurred: " + httpResponse.Error.Error())
		return
	}

	jsonResult, err := simplejson.NewJson(httpResponse.BodyBytes)
	if err != nil {
		_ = ctx.RespondText("An error has occurred: " + err.Error())
		return
	}

	meme := jsonResult.Get("meme")

	embed.NewEmbed().
		SetTitle(meme.Get("title").MustString()).
		SetImage(meme.Get("imagelink").MustString()).
		AddField("URL", fmt.Sprintf("`%s`", meme.Get("permalink").MustString())).
		SetFooter(fmt.Sprintf("Subreddit: %s | Upvotes: %d | Posted by: u/%s", meme.Get("subreddit").MustString(), meme.Get("upvotes").MustInt(), meme.Get("author").MustString())).
		Send(ctx)
}

func shibe(ctx *dgc.Ctx) {
	httpResponse := http.Get("https://shibe.online/api/shibes?count=1&urls=true&httpsUrls=true")
	if httpResponse.Error != nil {
		_ = ctx.RespondText("An error has occurred: " + httpResponse.Error.Error())
		return
	}

	jsonResult, err := simplejson.NewJson(httpResponse.BodyBytes)
	if err != nil {
		_ = ctx.RespondText("An error has occurred: " + err.Error())
		return
	}

	_ = ctx.RespondText(jsonResult.GetIndex(0).MustString())
}

func catgif(ctx *dgc.Ctx) {
	err := utils.Discord.SendFileFromURL(ctx, "https://cataas.com/cat/gif", "cat.gif")
	if err != nil {
		_ = ctx.RespondText("An error has occurred: " + err.Error())
	}
}

func hub(ctx *dgc.Ctx) {
	_ = ctx.RespondText(fmt.Sprintf(
		"https://api.lightning-bot.com/api/v3/imagegen/pornhub?text1=%s&text2=%s",
		url.QueryEscape(ctx.Arguments.Get(0).Raw()),
		url.QueryEscape(ctx.Arguments.Get(1).Raw()),
	))
}

func pokemon(ctx *dgc.Ctx) {
	_ = ctx.RespondText(
		fmt.Sprintf(
			"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/%d.png",
			utils.RandomNumber(1, 898),
		),
	)
}

func wasted(ctx *dgc.Ctx) {
	user, success := utils.Discord.GetUserFromMention(ctx.Session, ctx.Arguments.Get(0))
	if !success {
		user = ctx.Event.Author
	}

	_ = ctx.RespondText(
		fmt.Sprintf(
			"https://some-random-api.ml/canvas/wasted?avatar=%s",
			url.QueryEscape(user.AvatarURL("")),
		),
	)
}

func avatar(ctx *dgc.Ctx) {
	user, success := utils.Discord.GetUserFromMention(ctx.Session, ctx.Arguments.Get(0))
	if !success {
		user = ctx.Event.Author
	}

	_ = ctx.RespondText(user.AvatarURL(""))
}
