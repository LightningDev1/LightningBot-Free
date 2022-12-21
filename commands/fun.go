package commands

import (
	"fmt"
	"strings"

	"github.com/LightningDev1/LB-Selfbot-Free/embed"
	"github.com/LightningDev1/LB-Selfbot-Free/http"
	"github.com/LightningDev1/LB-Selfbot-Free/utils"
	"github.com/LightningDev1/dgc"
	"github.com/bitly/go-simplejson"
)

func (c *_Commands) RegisterFunCommands() {
	c.Router.StartCategory("Fun", "Fun commands")

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "eightball",
		Description: "Asks the 8-Ball a question",
		Usage:       "[p]eightball <question>",
		Handler:     eightball,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "howgay",
		Description: "See how gay a person is",
		Usage:       "[p]howgay <user>",
		Handler:     howgay,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "penis",
		Description: "See how big a person's penis is",
		Usage:       "[p]penis <user>",
		Handler:     penis,
	})

	c.Router.RegisterCmd(&dgc.Command{
		Name:        "trivia",
		Description: "Send a random trivia question",
		Usage:       "[p]trivia",
		Handler:     trivia,
	})
}

func eightball(ctx *dgc.Ctx) {
	possibleAnswers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it, yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	answer := utils.RandomChoice(possibleAnswers)
	embed.NewEmbed().
		AddField("Question", ctx.Arguments.Raw()).
		AddField("Answer", answer).
		Send(ctx)
}

func howgay(ctx *dgc.Ctx) {
	user, success := utils.Discord.GetUserFromMention(ctx.Session, ctx.Arguments.Get(0))
	if !success {
		user = ctx.Event.Author
	}

	embed.NewEmbed().
		AddField("User", user.String()).
		AddField("Gay Percentage", utils.RandomPercentage()).
		Send(ctx)
}

func penis(ctx *dgc.Ctx) {
	user, success := utils.Discord.GetUserFromMention(ctx.Session, ctx.Arguments.Get(0))
	if !success {
		user = ctx.Event.Author
	}

	penisLength := utils.RandomNumber(1, 20)
	penis := "8"
	for i := 0; i < penisLength; i++ {
		penis += "="
	}
	penis += "Đ"

	embed.NewEmbed().
		AddField("User", user.String()).
		AddField("Penis", penis).
		AddField("Penis Length", fmt.Sprintf("%d cm", penisLength)).
		Send(ctx)
}

func trivia(ctx *dgc.Ctx) {
	httpResponse := http.Get("https://opentdb.com/api.php?amount=1")
	if httpResponse.Error != nil {
		_ = ctx.RespondText("An error has occurred: " + httpResponse.Error.Error())
		return
	}

	rawJson := strings.Replace(httpResponse.Body, "&quot;", "\"", -1)
	rawJson = strings.Replace(rawJson, "&#039;", "'", -1)

	trivia, err := simplejson.NewJson([]byte(rawJson))
	if err != nil {
		_ = ctx.RespondText("An error has occurred.")
		return
	}

	trivia = trivia.Get("results").GetIndex(0)

	answers := trivia.Get("incorrect_answers").MustStringArray()
	answers = append(answers, trivia.Get("correct_answer").MustString())
	answers = utils.RandomShuffle(answers)

	embed.NewEmbed().
		AddField("Question", trivia.Get("question").MustString()).
		AddField("Category", trivia.Get("category").MustString()).
		AddField("Difficulty", utils.String.Capitalize(trivia.Get("difficulty").MustString())).
		AddField("Answers", strings.Join(answers, ", ")).
		AddField("Correct Answer", fmt.Sprintf("||%s||", trivia.Get("correct_answer").MustString())).
		Send(ctx)
}
