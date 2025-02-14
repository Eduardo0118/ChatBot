package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Eduardo0118/ChatBot/internal/adapters/deepseek"
	"github.com/Eduardo0118/ChatBot/internal/model"
	"github.com/Eduardo0118/ChatBot/internal/service"
)

func Chat(service service.Methods) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqBody []model.Message
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			return
		}

		messages := parseToDeepSeekMessage(reqBody)
		if err := service.Chat(messages); err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func parseToDeepSeekMessage(message []model.Message) []deepseek.Message {
	var messages []deepseek.Message

	for _, m := range message {
		messages = append(messages, deepseek.Message{
			Role:    m.Role,
			Content: m.Content,
		})
	}

	return messages
}
