package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	HUAWEI_PAD = "Huawei Pad"
	MI_PAD     = "Mi Pad"
)

type PadFactory struct{}

func (p *PadFactory) getElectronicDevice(name string) (Electron, error) {
	switch name {
	case HUAWEI_PAD:
		return new(HuaweiPad), nil
	case MI_PAD:
		return new(MiPad), nil
	default:
		return nil, errors.New(fmt.Sprintf("can't find any pad with the name %s!", name))
	}
}
