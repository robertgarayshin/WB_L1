package task21

import "fmt"

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts signal to USB")
	w.windowsMachine.InsertIntoUsbPort()
}
