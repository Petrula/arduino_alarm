package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TgToken  string
	ChatID   string
	PortName string
}

func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}
	tgToken := os.Getenv("TELEGRAM_TOKEN")
	chatID := os.Getenv("CHAT_ID")
	port := os.Getenv("PORT_NAME")
	if tgToken == "" || chatID == "" || port == "" {
		return Config{}, errors.New("Environment variable not found")
	}
	config := Config{
		TgToken:  tgToken,
		ChatID:   chatID,
		PortName: port,
	}
	return config, nil
}
