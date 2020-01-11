package proxy

type Client struct {
	Network VpnNetwork
}

func (c *Client) doRequest() (string, error) {
	return c.Network.connect()
}
