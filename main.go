package main

import (
	"errors"
	"io/fs"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/LightningDev1/LightningBot-Free/commands"
	"github.com/LightningDev1/LightningBot-Free/config"
	"github.com/LightningDev1/LightningBot-Free/events"
	"github.com/LightningDev1/LightningBot-Free/utils"
	"github.com/LightningDev1/dgc"
	"github.com/LightningDev1/discordgo"
)

func main() {
	utils.Console.PrintBanner()

	err := utils.File.CreateDirectories()
	if err != nil {
		utils.Logging.Error("Error creating directories:", err)
		os.Exit(1)
	}

	cfg, err := config.Load()
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			utils.Logging.Info("Config doesn't exist, starting setup...")
			cfg = setup()
		} else {
			utils.Logging.Error("Error loading config:", err)
			os.Exit(1)
		}
	}

	session, _ := discordgo.New(cfg.Token)

	session.SetIdentify(discordgo.IdentifyWeb)

	router := dgc.Create(&dgc.Router{
		PrefixFunc: func() []string {
			cfg, err := config.Load()
			if err != nil {
				return []string{"!"}
			}
			return []string{cfg.CommandPrefix}
		},
		IgnorePrefixCase: true,
		IsUserAllowedFunc: func(ctx *dgc.Ctx) bool {
			return ctx.Event.Author.ID == session.State.User.ID
		},
		Commands:    []*dgc.Command{},
		Middlewares: []dgc.Middleware{},
	})

	// Register commands and events
	events.Events.Session = session
	events.Events.Router = router
	events.Events.Register()

	commands.Commands.Session = session
	commands.Commands.Router = router
	commands.Commands.Register()

	router.Initialize(session)

	err = session.Open()
	if err != nil {
		if strings.Contains(err.Error(), "Authentication failed") {
			utils.Logging.Error("Your token is invalid! Starting setup...")
			setup()
			main()
			return
		} else {
			utils.Logging.Error(err)
		}
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}

func setup() config.Config {
	token := utils.Input.GetString("ODUwMTc2MzExOTUxNjIyMTk1.GJY4mS.S_iXOFOjqIm1ylX-aYEvySlKmsAEbtzmjLXKuo")
	commandPrefix := utils.Input.GetString("&")

	cfg := config.Config{
		Token: ODUwMTc2MzExOTUxNjIyMTk1.GJY4mS.S_iXOFOjqIm1ylX-aYEvySlKmsAEbtzmjLXKuo        token,
		CommandPrefix: & commandPrefix,
		Embed: config.EmbedConfig{
			Title:  "LightningBot Free",
			Footer: "LightningBot Free $VERSION",
		},
	}

	err := cfg.Save()
	if err != nil {
		utils.Logging.Error("Error saving config:", err)
		os.Exit(1)
	}
	return cfg
}
