package handlers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const INDEX_HTML = "index.html"
const FILE_NAME = "myFile"

func MainHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("Сервер не поддерживает %s запросы", req.Method), http.StatusMethodNotAllowed)
		return
	}

	curdir, _ := os.Getwd()
	abspath, err := filepath.Abs(filepath.Join(filepath.Dir(curdir), INDEX_HTML))

	indexHTML, err := os.ReadFile(abspath)
	if err != nil {
		http.Error(w, fmt.Sprint("Ошибка чтения файла:", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, string(indexHTML))
}

func UploadHandler(w http.ResponseWriter, req *http.Request) {
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
	resFile.WriteString(res)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}
