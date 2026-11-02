package handlers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const INDEX_HTML = "../index.html"
const FILE_NAME = "myFile"

func MainHandler(w http.ResponseWriter, req *http.Request) {
	file, err := os.ReadFile(INDEX_HTML)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}

func UploadHandler(w http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile(FILE_NAME)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer file.Close()

	var text string
	scaner := bufio.NewScanner(file)
	scaner.Scan()
	text = scaner.Text()

	filename := time.Now().UTC().String() + ".txt"
	resFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	res := service.Convert(text)
	resFile.WriteString(res)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
