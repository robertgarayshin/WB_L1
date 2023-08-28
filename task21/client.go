package task21

import "fmt"

type Client struct {
}

// Клиент, который хочет вставать Lightning в порт обычного компьютера
func (c *Client) InsertLightningPortIntoUSB(com Computer) {
	fmt.Println("Inserting")
	com.InsertIntoLightningPort()
}
