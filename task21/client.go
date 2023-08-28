package task21

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningPortIntoUSB(com Computer) {
	fmt.Println("Inserting")
	com.InsertIntoLightningPort()
}
