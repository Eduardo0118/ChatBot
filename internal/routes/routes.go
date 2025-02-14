package routes

import (
	"net/http"

	"github.com/Eduardo0118/ChatBot/internal/controller"
	"github.com/Eduardo0118/ChatBot/internal/service"
)

func New(service service.Methods) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.Chat(service))

	return mux
}
