package abstract_factory

import "fmt"

type HuaweiPad struct{}

func (h *HuaweiPad) ReadWithBigScreen(msg string) string {
	return fmt.Sprintf(" >>> Big Screen HuaweiPad is using for reading .")
}

func (h *HuaweiPad) Read() string {
	return fmt.Sprintf(" >>> Big Screen HuaweiPad reading .")
}
