package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	CellphoneType = 1
	PadType       = 2
)

type ElectronFactory interface {
	getElectronicDevice(name string) (Electron, error)
}

func GetElectronFactory(electronType int) (ElectronFactory, error) {
	switch electronType {
	case CellphoneType:
		return new(CellphoneFactory), nil
	case PadType:
		return new(PadFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf(" can't find any type facotry for the type value %d !", electronType))
	}
}
