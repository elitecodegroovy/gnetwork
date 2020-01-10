package proxy

import "fmt"

type ShanghaiVpnNetwork struct{}

func (s *ShanghaiVpnNetwork) connect() (string, error) {
	conInfo := fmt.Sprintf("Connect to Shanghai VPN network! ")
	//TODO ...init network for shanghai vpn network
	return conInfo, nil
}
