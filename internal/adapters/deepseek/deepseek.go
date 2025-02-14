package deepseek

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Methods interface {
	Chat(messages []Message) error
}

type deepSeek struct {
	cfg config
}

func New() Methods {
	return deepSeek{
		cfg: getConfig(),
	}
}

func (ds deepSeek) Chat(messages []Message) error {
	requestBody := buildRequestBody(messages, "deepseek-chat", 2048, 0.2)

	body, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/chat/completions", ds.cfg.url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ds.cfg.apiKey)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
