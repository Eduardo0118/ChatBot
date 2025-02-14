package service

import "github.com/Eduardo0118/ChatBot/internal/adapters/deepseek"

type Methods interface {
	Chat(messages []deepseek.Message) error
}

type service struct {
	ds deepseek.Methods
}

func New(ds deepseek.Methods) Methods {
	return service{
		ds: ds,
	}
}

func (s service) Chat(messages []deepseek.Message) error {
	return s.ds.Chat(messages)
}
