package commands

import (
	"fmt"
	"math"
	"strings"

	"github.com/LightningDev1/LB-Selfbot-Free/constants"
	"github.com/LightningDev1/LB-Selfbot-Free/embed"
	"github.com/LightningDev1/LB-Selfbot-Free/utils"
	"github.com/LightningDev1/dgc"
)

func (c *_Commands) RegisterHelpCommand() {
	c.Router.StopCategory()

	c.Router.RegisterCmd(&dgc.Command{
		Name:    "help",
		Handler: generalHelpCommand,
	})
}

func generalHelpCommand(ctx *dgc.Ctx) {
	// Check if the user provided an argument
	if ctx.Arguments.Amount() > 0 {
		specificHelpCommand(ctx)
		return
	}

	// Send the general help embed
	renderDefaultGeneralHelpEmbed(ctx.Router).Send(ctx)
}

func specificHelpCommand(ctx *dgc.Ctx) {
	oldName := strings.ToLower(ctx.Arguments.Get(0).Raw())
	name := strings.ReplaceAll(oldName, "eventreactions", "Event Reactions")
	name = strings.ReplaceAll(name, "eventreaction", "Event Reactions")
	name = strings.ReplaceAll(name, "nsfw", "NSFW")
	page, _ := ctx.Arguments.Get(1).AsInt()
	prefix := ctx.Router.GetPrefixes()[0]

	if command := ctx.Router.GetCmd(name); command != nil {
		renderCommandHelp(ctx.Router, command).Send(ctx)
	} else {
		if category := ctx.Router.GetCategory(utils.String.Capitalize(name)); category != nil {
			renderCategoryHelp(ctx.Router, category, oldName, page).Send(ctx)
		} else {
			var closeCommands []string
			for _, command := range ctx.Router.Commands {
				if len(closeCommands) >= 5 {
					break
				}
				distance := utils.Misc.GetDistance(command.Name, name)
				if distance < 2 || strings.Contains(command.Name, name) {
					closeCommands = append(closeCommands, strings.ReplaceAll(command.Usage, "[p]", prefix)+" - "+command.Description)
				}
			}
			var closeCategories []string
			for _, category := range ctx.Router.Categories {
				if len(closeCategories) >= 5 {
					break
				}
				distance := utils.Misc.GetDistance(category.Name, utils.String.Capitalize(name))
				if distance < 2 || strings.Contains(category.Name, utils.String.Capitalize(name)) {
					closeCategories = append(closeCategories, category.Name+" - "+category.Description)
				}
			}
			description := fmt.Sprintf("`%s` is not a command or category!", name)
			if len(closeCommands) > 0 {
				description += "\n\nDid you mean one of these commands?\n\n" + strings.Join(closeCommands, "\n")
			}
			if len(closeCategories) > 0 {
				description += "\n\nDid you mean one of these categories?\n\n" + strings.Join(closeCategories, "\n")
			}
			embed.NewEmbed().SetDescription(description).Send(ctx)
		}
	}
}

func renderCommandHelp(router *dgc.Router, command *dgc.Command) *embed.Embed {
	// Try to get the category the command belongs in
	categoryName := ""
	for _, category := range router.Categories {
		category = router.GetCategory(category.Name)
		for _, cmd := range category.Commands {
			if cmd.Name == command.Name {
				categoryName = category.Name
				break
			}
		}
	}
	
	return embed.NewEmbed().
		SetDescription("Command Help: "+command.Name).
		AddField("Command Name", command.Name).
		AddField("Usage", "<> is required, [] is optional\n\n"+strings.ReplaceAll(command.Usage, "[p]", router.GetPrefixes()[0])).
		AddField("Description", command.Description).
		AddField("Category", categoryName)
}

func renderCategoryHelp(router *dgc.Router, category *dgc.Category, rawArgument string, page int) *embed.Embed {
	helpPages := map[int]string{}
	currentPage := 0
	for _, command := range category.Commands {
		splitUsage := strings.SplitN(command.Usage, " ", 2)
		usage := ""
		if len(splitUsage) > 1 {
			usage = " " + splitUsage[1]
		}
		helpText := fmt.Sprintf("%s%s%s - %s\n", router.GetPrefixes()[0], command.Name, usage, command.Description)
		if len(helpPages[currentPage])+len(helpText) > 1700 {
			currentPage++
		}
		helpPages[currentPage] += helpText
	}
	page = int(math.Max(0, float64(page-1)))

	emb := embed.NewEmbed().
		SetDescription(fmt.Sprintf("<> is required, [] is optional\n\nAmount of %s commands: %d", category.Name, len(category.Commands))).
		SetSubtext(fmt.Sprintf("Page: %d/%d", page+1, len(helpPages)))

	if page < len(helpPages) {
		emb.AddField("Commands", helpPages[page], true, true)
	} else {
		return embed.NewEmbed().SetDescription(fmt.Sprintf("That page does not exist! %s has %d pages.", category.Name, len(helpPages)))
	}
	if len(helpPages) > 1 {
		emb.SetSubtext(fmt.Sprintf("Page: %d/%d\n\nUse \x1b[34;1m%shelp %s <page>\x1b[0m to go to the next page", page+1, len(helpPages), router.GetPrefixes()[0], rawArgument))
	}
	return emb
}

func renderDefaultGeneralHelpEmbed(router *dgc.Router) *embed.Embed {
	categoriesString := ""
	for _, category := range router.Categories {
		// Get the commands in the category, will be empty otherwise
		category = router.GetCategory(category.Name)
		categoriesString += fmt.Sprintf("%s - %s\n", category.Name+strings.Repeat(" ", 15-len(category.Name)), category.Description)
	}
	prefix := router.GetPrefixes()[0]
	return embed.NewEmbed().
		SetDescription(fmt.Sprintf("LightningBot Free %s Help\nAmount of Commands: %d, <> is required, [] is optional\n", constants.VERSION, len(router.Commands))).
		AddField("Categories", categoriesString, true, true).
		SetSubtext(fmt.Sprintf(
			"Use \x1b[34;1m%shelp <category name>\x1b[0m or \x1b[34;1m%shelp <command name>\x1b[0m to get more info.",
			prefix,
			prefix,
		))
}
