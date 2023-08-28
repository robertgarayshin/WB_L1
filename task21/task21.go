package task21

func Task21() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningPortIntoUSB(mac)

	windowsMachine := &Windows{}

	windowsMachineAdapter := &WindowsAdapter{windowsMachine: windowsMachine}

	client.InsertLightningPortIntoUSB(windowsMachineAdapter)
}
