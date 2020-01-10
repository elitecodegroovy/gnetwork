package proxy

type VpnNetwork interface {
	connect() (string, error)
}
