package abstract_factory

import "fmt"

type HuaweiCellphone struct{}

func (m *HuaweiCellphone) Call(number string) string {
	return fmt.Sprintf(" Huawei Cellphone call %s", number)
}

func (m *HuaweiCellphone) Send(msg string) string {
	return fmt.Sprintf(" Huawei Cellphone sent message %s", msg)
}

func (m *HuaweiCellphone) CallWithVideo(msg string) string {
	return fmt.Sprintf(" >>> Huawei SmartCellphone call  '%s' with video", msg)
}

func (m *HuaweiCellphone) Read() string {
	return fmt.Sprintf(" >>> Huawei SmartCellphone is using for reading .")
}
