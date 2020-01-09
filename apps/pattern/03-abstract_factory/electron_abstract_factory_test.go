package abstract_factory

import (
	"strings"
	"testing"
)

func TestGetElectronFactory(t *testing.T) {
	electronFactory, err := GetElectronFactory(1)
	if err != nil {
		t.Fatal("msg :" + err.Error())
	}
	electron, err1 := electronFactory.getElectronicDevice("HUAWEI")
	if err1 != nil {
		t.Fatal("elctron msg" + err1.Error())
	}
	result := electron.Read()
	if !strings.Contains(result, "Huawei SmartCellphone") {
		t.Fatal("Huawei SmartCellphone doesn't work!")
	}

	huaweiPhone, ok := electron.(CellPhone)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	callMsg := huaweiPhone.Call("15914313549")
	if !strings.Contains(callMsg, "Huawei Cellphone") {
		t.Fatal("Huawei SmartCellphone calling doesn't work!")
	}
	videoContent := huaweiPhone.CallWithVideo("Let's talk with video!")
	if !strings.Contains(videoContent, "Huawei SmartCellphone call  'Let's talk with video!'") {
		t.Fatal("Huawei SmartCellphone CallWithVideo doesn't work!")
	}

	msg := huaweiPhone.Send("Hi!")
	if !strings.Contains(msg, "Huawei Cellphone sent message Hi!") {
		t.Fatal("Huawei SmartCellphone Sending doesn't work!")
	}

}
