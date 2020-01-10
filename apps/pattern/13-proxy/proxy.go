package proxy

import "errors"

const (
	ShangHaiNetwork   = 1
	GugangZhouNetwork = 2
)

type VpnProxy struct {
	Type int32
}

func (v *VpnProxy) connect() (string, error) {
	switch v.Type {
	case ShangHaiNetwork:
		shanghaiVpnNetwork := new(ShanghaiVpnNetwork)
		return shanghaiVpnNetwork.connect()
	case GugangZhouNetwork:
		GuangzhouVpnNetwork := new(GuangzhouVpnNetwork)
		return GuangzhouVpnNetwork.connect()
	default:
		return "", errors.New("can't find any vpn network")
	}
}
