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
	if req.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("Сервер не поддерживает %s запросы", req.Method), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, req, INDEX_HTML)
}

func UploadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Сервер не поддерживает %s запросы", req.Method), http.StatusMethodNotAllowed)
		return
	}

	file, _, err := req.FormFile(FILE_NAME)
	if err != nil {
		http.Error(w, fmt.Sprint("Ошибка при получении файла:", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	var text string
	scaner := bufio.NewScanner(file)
	scaner.Scan()
	text = scaner.Text()

	filename := time.Now().UTC().String() + ".txt"
	resFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, fmt.Sprint("Ошибка при создании файла:", err), http.StatusInternalServerError)
		return
	}
	defer resFile.Close()

	res := service.Convert(text)
	_, err = resFile.WriteString(res)
	if err != nil {
		http.Error(w, fmt.Sprint("Ошибка при написании ответа:", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
