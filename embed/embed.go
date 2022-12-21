package embed

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/LightningDev1/LB-Selfbot-Free/config"
	"github.com/LightningDev1/LB-Selfbot-Free/constants"
	"github.com/LightningDev1/dgc"
	"github.com/LightningDev1/discordgo"
)

// https://github.com/Clinet/discordgo-embed

// Embed ...
type Embed struct {
	*discordgo.MessageEmbed
	Subtext string
}

var (
	LimitTitle       = 256
	LimitDescription = 2048
	LimitFieldValue  = 1024
	LimitFieldName   = 256
	LimitField       = 25
	LimitFooter      = 2048
)

// NewEmbed returns a new embed object
func NewEmbed() *Embed {
	embed := &Embed{&discordgo.MessageEmbed{}, ""}
	cfg, err := config.Load()
	if err != nil {
		return embed
	}
	embed.SetTitle(strings.Replace(cfg.Embed.Title, "$VERSION", constants.VERSION, -1))
	embed.SetFooter(strings.Replace(cfg.Embed.Footer, "$VERSION", constants.VERSION, -1))
	return embed
}

// SetTitle ...
func (e *Embed) SetTitle(name string) *Embed {
	e.Title = name
	return e
}

// SetDescription [desc]
func (e *Embed) SetDescription(description string) *Embed {
	if len(description) > 2048 {
		description = description[:2048]
	}
	e.Description = description
	return e
}

// SetSubtext ...
func (e *Embed) SetSubtext(text string) *Embed {
	e.Subtext = text
	return e
}

// AddField [name] [value]
func (e *Embed) AddField(name, value string, flags ...bool) *Embed {
	inline := len(flags) > 0
	bypass := len(flags) > 1
	fields := make([]*discordgo.MessageEmbedField, 0)

	if len(name) > LimitFieldName {
		name = name[:LimitFieldName]
	}

	if len(value) > LimitFieldValue && !bypass {
		i := LimitFieldValue
		extended := false
		for i = LimitFieldValue; i < len(value); {
			if i != LimitFieldValue && !extended {
				name += " (extended)"
				extended = true
			}
			if value[i] == []byte(" ")[0] || value[i] == []byte("\n")[0] || value[i] == []byte("-")[0] {
				fields = append(fields, &discordgo.MessageEmbedField{
					Name:   name,
					Value:  value[i-LimitFieldValue : i],
					Inline: inline,
				})
			} else {
				fields = append(fields, &discordgo.MessageEmbedField{
					Name:   name,
					Value:  value[i-LimitFieldValue:i-1] + "-",
					Inline: inline,
				})
				i--
			}

			if (i + LimitFieldValue) < len(value) {
				i += LimitFieldValue
			} else {
				break
			}
		}
		if i < len(value) {
			name += " (extended)"
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:   name,
				Value:  value[i:],
				Inline: inline,
			})
		}
	} else {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   name,
			Value:  value,
			Inline: inline,
		})
	}

	e.Fields = append(e.Fields, fields...)

	return e
}

// SetFooter [Text] [iconURL]
func (e *Embed) SetFooter(args ...string) *Embed {
	iconURL := ""
	text := ""
	proxyURL := ""

	switch {
	case len(args) > 2:
		proxyURL = args[2]
		fallthrough
	case len(args) > 1:
		iconURL = args[1]
		fallthrough
	case len(args) > 0:
		text = args[0]
	case len(args) == 0:
		return e
	}

	e.Footer = &discordgo.MessageEmbedFooter{
		IconURL:      iconURL,
		Text:         text,
		ProxyIconURL: proxyURL,
	}

	return e
}

// SetImage ...
func (e *Embed) SetImage(args ...string) *Embed {
	var URL string
	var proxyURL string

	if len(args) == 0 {
		return e
	}
	if len(args) > 0 {
		URL = args[0]
	}
	if len(args) > 1 {
		proxyURL = args[1]
	}
	e.Image = &discordgo.MessageEmbedImage{
		URL:      URL,
		ProxyURL: proxyURL,
	}
	return e
}

// SetThumbnail ...
func (e *Embed) SetThumbnail(args ...string) *Embed {
	var URL string
	var proxyURL string

	if len(args) == 0 {
		return e
	}
	if len(args) > 0 {
		URL = args[0]
	}
	if len(args) > 1 {
		proxyURL = args[1]
	}
	e.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL:      URL,
		ProxyURL: proxyURL,
	}
	return e
}

// SetAuthor ...
func (e *Embed) SetAuthor(args ...string) *Embed {
	var (
		name     string
		iconURL  string
		URL      string
		proxyURL string
	)

	if len(args) == 0 {
		return e
	}
	if len(args) > 0 {
		name = args[0]
	}
	if len(args) > 1 {
		iconURL = args[1]
	}
	if len(args) > 2 {
		URL = args[2]
	}
	if len(args) > 3 {
		proxyURL = args[3]
	}

	e.Author = &discordgo.MessageEmbedAuthor{
		Name:         name,
		IconURL:      iconURL,
		URL:          URL,
		ProxyIconURL: proxyURL,
	}

	return e
}

// SetURL ...
func (e *Embed) SetURL(URL string) *Embed {
	e.URL = URL
	return e
}

// SetColor ...
func (e *Embed) SetColor(clr int) *Embed {
	e.Color = clr
	return e
}

// InlineAllFields sets all fields in the embed to be inline
func (e *Embed) InlineAllFields() *Embed {
	for _, v := range e.Fields {
		v.Inline = true
	}
	return e
}

// Truncate truncates any embed value over the character limit.
func (e *Embed) Truncate() *Embed {
	e.TruncateDescription()
	e.TruncateFields()
	e.TruncateFooter()
	e.TruncateTitle()
	return e
}

// MakeFieldInline adds last field as InLine
func (e *Embed) MakeFieldInline() *Embed {
	length := len(e.Fields) - 1
	e.Fields[length].Inline = true
	return e
}

// TruncateFields truncates fields that are too long
func (e *Embed) TruncateFields() *Embed {
	if len(e.Fields) > 25 {
		e.Fields = e.Fields[:LimitField]
	}

	for _, v := range e.Fields {

		if len(v.Name) > LimitFieldName {
			v.Name = v.Name[:LimitFieldName]
		}

		if len(v.Value) > LimitFieldValue {
			v.Value = v.Value[:LimitFieldValue]
		}

	}
	return e
}

// TruncateDescription ...
func (e *Embed) TruncateDescription() *Embed {
	if len(e.Description) > LimitDescription {
		e.Description = e.Description[:LimitDescription]
	}
	return e
}

// TruncateTitle ...
func (e *Embed) TruncateTitle() *Embed {
	if len(e.Title) > LimitTitle {
		e.Title = e.Title[:LimitTitle]
	}
	return e
}

// TruncateFooter ...
func (e *Embed) TruncateFooter() *Embed {
	if e.Footer != nil && len(e.Footer.Text) > LimitFooter {
		e.Footer.Text = e.Footer.Text[:LimitFooter]
	}
	return e
}

func (e *Embed) ToJson() string {
	jsonString, _ := json.Marshal(e)
	return string(jsonString)
}

func (e *Embed) ToCodeblock() []string {
	embedString := ""
	var embedList []string
	if e.Author != nil {
		embedString += fmt.Sprintf("```fix\n%s\n```\n", e.Author.Name)
	}
	if e.Title != "" {
		embedString += fmt.Sprintf("```ini\n[%s]\n```", e.Title)
	}
	if e.Description != "" {
		if len(embedString+e.Description) > 2000 {
			embedList = append(embedList, embedString)
			embedString = ""
		}
		embedString += fmt.Sprintf("```md\n%s\n```", e.Description)
	}
	if len(e.Fields) > 0 {
		mdAdded := false
		for _, v := range e.Fields {
			fieldText := fmt.Sprintf("# %s\n%s\n\n", v.Name, v.Value)
			if len(embedString+fieldText) > 2000 {
				if mdAdded {
					embedString += "```"
					mdAdded = false
				}
				embedList = append(embedList, embedString)
				embedString = ""
			}
			if !mdAdded {
				embedString += "```md\n"
				mdAdded = true
			}
			embedString += fieldText
		}
		embedString += "```"
	}
	if e.Subtext != "" {
		embedString += fmt.Sprintf("```ansi\n%s\n```", e.Subtext)
	}
	if e.Footer != nil {
		footerString := fmt.Sprintf("```css\n[%s]```", e.Footer.Text)
		if len(embedString+footerString) > 2000 {
			embedList = append(embedList, embedString)
			embedString = ""
		}
		embedString += footerString
	}
	if embedString != "" {
		embedList = append(embedList, embedString)
	}
	return embedList
}

func (e *Embed) ToText() []string {
	str := ""
	if e.Author != nil {
		str += fmt.Sprintf("%s\n\n", e.Author.Name)
	}
	if e.Title != "" {
		str += fmt.Sprintf("```ini\n[%s]```\n\n", e.Title)
	}
	if e.Description != "" {
		str += fmt.Sprintf("%s\n\n", e.Description)
	}
	if len(e.Fields) > 0 {
		for _, v := range e.Fields {
			if v.Inline {
				v.Value = "```" + v.Value + "```"
			}
			str += fmt.Sprintf("> **%s**\n%s\n\n", v.Name, v.Value)
		}
	}
	if e.Subtext != "" {
		str += fmt.Sprintf("```ansi\n%s\n```", e.Subtext)
	}
	if e.Footer != nil {
		str += fmt.Sprintf("\n```css\n[%s]```\n", e.Footer.Text)
	}
	var strings_ []string
	newStr := ""
	for _, line := range strings.Split(str, "\n") {
		if len(newStr+line)+1 > 2000 {
			strings_ = append(strings_, newStr)
			newStr = ""
		}
		newStr += line + "\n"
	}
	strings_ = append(strings_, newStr)
	return strings_
}

func (e *Embed) ToURLEmbed() []string {
	return []string{"Not implemented"}
}

func (e *Embed) ToString() []string {
	return e.ToText()
}

func (e *Embed) Send(ctx *dgc.Ctx) *discordgo.Message {
	var message *discordgo.Message
	for _, text := range e.ToString() {
		message, _ = ctx.Session.ChannelMessageSend(ctx.Event.ChannelID, text)
	}
	if e.Image != nil && e.Image.URL != "" {
		_, _ = ctx.Session.ChannelMessageSend(ctx.Event.ChannelID, e.Image.URL)
	}
	return message
}
