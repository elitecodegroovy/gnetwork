package adapter

import "fmt"

type mac struct {
}

func (m *mac) drawGUI() string {
	return fmt.Sprintf("MAC insert In Square Port")
}
