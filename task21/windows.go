package task21

import "fmt"

type Windows struct{}

func (w *Windows) InsertIntoUsbPort() {
	fmt.Println("USB plugged into windows machine")
}
