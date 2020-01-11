package adapter

type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) drawGUI() string {
	return w.windowMachine.drawGUI()
}
