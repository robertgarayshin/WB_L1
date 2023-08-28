package task21

import "fmt"

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	// Адаптер, который позволяет подключить устройство Apple (наш сервис) к устройству Windows (неизвестный сервис)
	fmt.Println("Adapter converts signal to USB")
	w.windowsMachine.InsertIntoUsbPort()
}
