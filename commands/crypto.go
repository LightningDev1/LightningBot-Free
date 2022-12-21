package commands

import (
	"fmt"

	"github.com/LightningDev1/LightningBot-Free/embed"
	"github.com/LightningDev1/LightningBot-Free/http"
	"github.com/LightningDev1/dgc"
	"github.com/bitly/go-simplejson"
)

func (c *_Commands) RegisterCryptoCommands() {
	c.Router.StartCategory("Crypto", "Crypto commands")

	cryptoCommands := [][]string{
		{"btc", "BTC", "Bitcoin", "1"},
		{"eth", "ETH", "Ethereum", "1027"},
		{"bnb", "BNB", "Binance Coin", "1839"},
		{"dogecoin", "DOGE", "Dogecoin", "74"},
		{"ripple", "XRP", "Ripple", "52"},
		{"bitcoincash", "BCH", "Bitcoin Cash", "1831"},
		{"litecoin", "LTC", "Litecoin", "2"},
	}

	for _, cryptoCommand := range cryptoCommands {
		c.Router.RegisterCmd(&dgc.Command{
			Name:        cryptoCommand[0],
			Description: "Sends the current " + cryptoCommand[2] + " price in USD and Euro",
			Usage:       "[p]" + cryptoCommand[0],
			Handler:     createCryptoCommand(cryptoCommand),
		})
	}
}

func createCryptoCommand(element []string) func(ctx *dgc.Ctx) {
	// This makes sure that the handler function is not overwritten by the next command.
	// If you put the handler function in the RegisterCommands function, the element
	// variable will always be the last element in the command array.
	return func(ctx *dgc.Ctx) {
		httpResponse := http.Get("https://min-api.cryptocompare.com/data/price?fsym=" + element[1] + "&tsyms=USD,EUR")
		if httpResponse.Error != nil {
			_ = ctx.RespondText("An error has occurred.")
		}

		json, err := simplejson.NewJson(httpResponse.BodyBytes)
		if err != nil {
			_ = ctx.RespondText("An error has occurred.")
			return
		}

		usd := json.Get("USD").MustFloat64()
		eur := json.Get("EUR").MustFloat64()

		embed.NewEmbed().
			SetDescription(fmt.Sprintf("USD: %.3f$\nEUR: %.3f€", usd, eur)).
			SetAuthor(element[2], fmt.Sprintf("https://s2.coinmarketcap.com/static/img/coins/200x200/%s.png", element[3])).
			Send(ctx)
	}
}
