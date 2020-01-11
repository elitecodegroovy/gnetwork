package adapter

import "fmt"

type windows struct{}

func (w *windows) drawGUI() string {
	return fmt.Sprintf("WINDOWS Insert In Circle Port!")
}
