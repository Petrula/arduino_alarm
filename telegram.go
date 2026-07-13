package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendAlert(token, chatID, text string) error {
	URL := "https://api.telegram.org/bot" + token + "/sendMessage"
	parameters := map[string]string{
		"chat_id": chatID,
		"text":    text,
	}
	jsonBytes, err := json.Marshal(parameters)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(jsonBytes)
	resp, err := http.Post(URL, "application/json", reader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to send the request.")
	}
	return nil
}
