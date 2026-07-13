package main

import (
	"bufio"
	"fmt"

	"go.bug.st/serial"
)

func ListenPort(port serial.Port, config Config) {
	scanner := bufio.NewScanner(port)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "ALARM" {
			fmt.Println("Получено верное сообщение")
			err := SendAlert(config.TgToken, config.ChatID, line)
			if err != nil {
				fmt.Println("Пробуйте еще")
				port.Write([]byte("FAIL\n"))
			} else {
				fmt.Println("Отправлено в ТЕЛЕГРАММ")
				port.Write([]byte("OK\n"))
			}
		} else {
			fmt.Println("Получено не верное сообщение")
		}
	}
}
