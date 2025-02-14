package main

import (
	"log"
	"net/http"

	"github.com/Eduardo0118/ChatBot/internal/adapters/deepseek"
	"github.com/Eduardo0118/ChatBot/internal/routes"
	"github.com/Eduardo0118/ChatBot/internal/service"
)

func main() {
	// Inicialização dos Adapters
	deepseekAdapter := deepseek.New()

	// Inicialização dos Services
	service := service.New(deepseekAdapter)

	r := routes.New(service)

	log.Println("Servidor iniciado na porta 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
