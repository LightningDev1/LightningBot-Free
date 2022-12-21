package commands

func (c *_Commands) RegisterConfigCommands() {
	c.Router.StartCategory("Config", "Config commands")
}
