package main

import (
	"fmt"

	"go.bug.st/serial"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	mode := serial.Mode{BaudRate: 9600}
	port, err := serial.Open(config.PortName, &mode)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer port.Close()
	ListenPort(port, config)
}
