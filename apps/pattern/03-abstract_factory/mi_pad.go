package abstract_factory

import "fmt"

type MiPad struct{}

func (h *MiPad) ReadWithBigScreen(msg string) string {
	return fmt.Sprintf(" >>> Big Screen MiPad is using for reading .")
}

func (h *MiPad) Read() string {
	return fmt.Sprintf(" >>> Big Screen MiPad reading .")
}
