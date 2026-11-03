package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Server *http.Server
	Logger *log.Logger
}

func CreateRouter(logger *log.Logger) Server {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.MainHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)
	server := http.Server{
		Addr:         "localhost:8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 15,
	}
	return Server{
		Server: &server,
		Logger: logger,
	}
}
