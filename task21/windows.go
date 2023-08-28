package task21

import "fmt"

type Windows struct{}

func (w *Windows) InsertIntoUsbPort() {
	// Неизвестный сервис, который подключает что-то в USB-порт
	fmt.Println("USB plugged into windows machine")
}
