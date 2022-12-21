package events

import (
	"fmt"

	"github.com/LightningDev1/LightningBot-Free/config"
	"github.com/LightningDev1/LightningBot-Free/constants"
	"github.com/LightningDev1/LightningBot-Free/utils"
	"github.com/LightningDev1/discordgo"
)

func (e *_Events) OnReady(_ *discordgo.Session, event *discordgo.Ready) {
	utils.Console.PrintBanner()

	cfg, err := config.Load()
	if err != nil {
		utils.Logging.Error("Error loading config:", err)
		return
	}

	friends := getFriendsAmount(event)

	bannerTextUpper := fmt.Sprintf("LightningBot Free %s | %s | Prefix: %s", constants.VERSION, event.User.String(), cfg.CommandPrefix)
	bannerTextLower := fmt.Sprintf("%d Commands | %d Servers | %d Friends", len(e.Router.Commands), len(event.Guilds), friends)

	message := utils.Console.Colored(
		utils.String.CenterSlices([]string{bannerTextUpper})+"\n"+utils.String.CenterSlices([]string{bannerTextLower}), 0, 203, 255, false,
	)

	fmt.Println(message)
}

func getFriendsAmount(event *discordgo.Ready) int {
	friends := 0

	for _, rel := range event.Relationships {
		if rel.Type == 1 {
			friends++
		}
	}

	return friends
}
