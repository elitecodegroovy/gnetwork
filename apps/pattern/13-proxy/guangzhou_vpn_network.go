package proxy

import "fmt"

const (
	HiResponseForGuangzhou = "Connect to Guangzhou VPN network!"
	HiResponseForShangHai  = "Connect to Shanghai VPN network!"
)

type GuangzhouVpnNetwork struct{}

func (s *GuangzhouVpnNetwork) connect() (string, error) {
	conInfo := fmt.Sprintf(HiResponseForGuangzhou)
	//TODO ...init network for shanghai vpn network
	return conInfo, nil
}
