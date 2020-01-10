package adapter

type client struct {
}

func (c *client) drawGUIInComputer(com computer) string {
	return com.drawGUI()
}
