package task21

import "fmt"

type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	// Логика нашего сервиса (позволяет устройству Apple подключиться только к устройству Apple)
	fmt.Println("Lightning connected into Mac machine")
}
