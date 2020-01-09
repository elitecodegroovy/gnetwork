package abstract_factory

import "fmt"

type SamsungCellphone struct{}

func (m *SamsungCellphone) Call(number string) string {
	return fmt.Sprintf(" Samsung Cellphone call %s", number)
}

func (m *SamsungCellphone) Send(msg string) string {
	return fmt.Sprintf(" Samsung Cellphone sent message %s", msg)
}

func (m *SamsungCellphone) CallWithVideo(msg string) string {
	return fmt.Sprintf(" >>> Samsung SmartCellphone call  '%s' with video !", msg)
}

func (m *SamsungCellphone) Read() string {
	return fmt.Sprintf(" >>>  SamsungCellphone is using for reading .")
}
