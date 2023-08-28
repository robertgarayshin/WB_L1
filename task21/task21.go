package task21

// Суть работы паттерна - адаптация неизвестного сервиса под наш функционал

// Мы имеем код клиента, ожидающий определенных от объекта определенных качеств (порт Lightning),
// но также мы имеем другой объект под названием adaptee (ноутбук на Windows),
// который предоставляет тот же функционал, но через другой интерфейс (USB порт).

func Task21() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningPortIntoUSB(mac)

	windowsMachine := &Windows{}

	windowsMachineAdapter := &WindowsAdapter{windowsMachine: windowsMachine}

	client.InsertLightningPortIntoUSB(windowsMachineAdapter)
}
