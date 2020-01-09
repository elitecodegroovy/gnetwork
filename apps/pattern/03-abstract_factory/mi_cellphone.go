package abstract_factory

import "fmt"

type MiCellphone struct{}

func (m *MiCellphone) Call(number string) string {
	return fmt.Sprintf(" MI Cellphone call %s", number)
}

func (m *MiCellphone) Send(msg string) string {
	return fmt.Sprintf(" MI Cellphone sent message %s", msg)
}

func (m *MiCellphone) CallWithVideo(msg string) string {
	return fmt.Sprintf(" >>> Mi SmartCellphone call  '%s' with video", msg)
}

func (m *MiCellphone) Read() string {
	return fmt.Sprintf(" >>>  MiCellphone is using for reading .")
}
