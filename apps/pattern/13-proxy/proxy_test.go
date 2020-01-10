package proxy

import (
	"strings"
	"testing"
)

func TestProxy(t *testing.T) {
	shangHaiClient := &Client{
		Network: new(GuangzhouVpnNetwork),
	}
	result, err := shangHaiClient.doRequest()
	if err != nil {
		t.Fatal("can't connect to Shanghai VPN network! ")
	}
	if !strings.Contains(result, HiResponseForGuangzhou) {
		t.Fatal("response is not correct! ")
	}

	vpnProxy1 := &VpnProxy{
		Type: 2,
	}

	proxy1 := &Client{
		Network: vpnProxy1,
	}
	result, err = proxy1.doRequest()
	if err != nil {
		t.Fatal("can't connect to VPN network! ")
	}
	if !strings.Contains(result, HiResponseForGuangzhou) {
		t.Fatal("response is not correct! ")
	}

	vpnProxy2 := &VpnProxy{
		Type: 1,
	}
	proxy2 := &Client{
		Network: vpnProxy2,
	}
	result, err = proxy2.doRequest()
	if err != nil {
		t.Fatal("can't connect to VPN network! ")
	}
	if !strings.Contains(result, HiResponseForShangHai) {
		t.Fatal("response is not correct! ")
	}
}
