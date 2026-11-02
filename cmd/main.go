package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("Ошибка открытия файла логирования:", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "logger: ", log.Lshortfile)
	server := server.CreateRouter(logger)

	err = server.Server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
