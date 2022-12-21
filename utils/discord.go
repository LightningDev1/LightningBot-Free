package utils

import (
	"bytes"
	"time"

	"github.com/LightningDev1/LB-Selfbot-Free/http"
	"github.com/LightningDev1/dgc"
	"github.com/LightningDev1/discordgo"
)

type DiscordUtils struct{}

func (DiscordUtils) GetUserFromMention(session *discordgo.Session, mention *dgc.Argument) (*discordgo.User, bool) {
	userID := Misc.Or(mention.AsUserMentionID(), mention.Raw())

	// Try to get user from cache
	user, err := session.State.GetUser(userID)
	if err == nil {
		return user, true
	}

	// Try to get user from API
	user, err = session.User(userID)
	if err == nil {
		return user, true
	}

	return nil, false
}

func (DiscordUtils) SendFileFromURL(ctx *dgc.Ctx, url string, fileName string, headers ...map[string]string) error {
	httpResponse := http.Get(url, headers...)
	if httpResponse.Error != nil {
		return httpResponse.Error
	}

	_, err := ctx.Session.ChannelFileSend(
		ctx.Event.ChannelID,
		fileName,
		bytes.NewReader(httpResponse.BodyBytes),
	)
	return err
}

func (DiscordUtils) GetChannel(session *discordgo.Session, channelID string) (*discordgo.Channel, error) {
	// Try to get channel from cache
	channel, err := session.State.Channel(channelID)
	if err == nil {
		return channel, nil
	}

	// Try to get channel from API
	channel, err = session.Channel(channelID)
	if err == nil {
		return channel, nil
	}

	return nil, err
}

func (DiscordUtils) GetChannelFromMention(ctx *dgc.Ctx, mention *dgc.Argument) (*discordgo.Channel, bool) {
	channelID := Misc.Or(mention.AsChannelMentionID(), mention.Raw())

	channel, err := Discord.GetChannel(ctx.Session, channelID)
	if err != nil {
		return nil, false
	}

	return channel, true
}

func (DiscordUtils) SleepWithTyping(session *discordgo.Session, channelID string, duration time.Duration) {
	// Send ChannelTyping every 5 seconds
	iterations := int(duration.Seconds() / 5)

	for i := 0; i < int(duration.Seconds()/5); i++ {
		_ = session.ChannelTyping(channelID)
		time.Sleep(5 * time.Second)
	}

	// Sleep for the remaining time
	interationsTime := time.Duration(iterations*5) * time.Second

	time.Sleep(duration - interationsTime)
}

var Discord = &DiscordUtils{}
