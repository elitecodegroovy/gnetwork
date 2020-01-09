package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	HUA_WEI = "HUAWEI"
	MI      = "MI"
	SAMSUNG = "SANSUNG"
)

type CellphoneFactory struct{}

func getCellphone(name string) (Electron, error) {
	switch name {
	case HUA_WEI:
		return new(HuaweiCellphone), nil
	case MI:
		return new(MiCellphone), nil
	case SAMSUNG:
		return new(SamsungCellphone), nil
	default:
		return nil, errors.New(fmt.Sprintf("can't find a cellphone by name %s!", name))
	}
}

func (c *CellphoneFactory) getElectronicDevice(name string) (Electron, error) {
	return getCellphone(name)
}
