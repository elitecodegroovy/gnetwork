package adapter

import (
	"strings"
	"testing"
)

func TestAdapter(t *testing.T) {
	client := &client{}
	mac := &mac{}
	macResult := client.drawGUIInComputer(mac)
	if !strings.Contains(macResult, "MAC") {
		t.Fatal("mac calling failed!")
	}
	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}
	windowsResult := client.drawGUIInComputer(windowsMachineAdapter)
	if !strings.Contains(windowsResult, "WINDOWS") {
		t.Fatal("windows calling failed!")
	}
}
