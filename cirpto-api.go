package main

import (
	"cripto-util/aes-utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/encrypt", Encrypt).Methods("POST")
	router.HandleFunc("/api/v1/decrypt", Decrypt).Methods("POST")

	log.Printf("server start..")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func Decrypt(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("error read file: %s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	fileName := uuid.New().String()
	tempFileName := uuid.New().String()

	defer file.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("error read file: %s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("File was correct load")
	tempFile, err := os.CreateTemp("./", fileName)
	if err != nil {
		log.Printf("error create temp dir")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer tempFile.Close()
	log.Println("Temp file success create")
	tempFile.Write(fileBytes)
	aes.DecryptLargeFiles(tempFile.Name(), tempFileName, []byte("12345678912345678912345678912345"))
	returnFileBytes, err := os.ReadFile(tempFileName)
	if err != nil {
		log.Printf("error crypto file")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(returnFileBytes)
	defer os.Remove(tempFileName)
	defer os.Remove(tempFile.Name())
}

func Encrypt(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("error read file: %s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
	fileName := uuid.New().String()
	tempFileName := uuid.New().String()

	defer file.Close()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("error read file: %s", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("File was correct load")
	tempFile, err := os.CreateTemp("./", fileName)
	if err != nil {
		log.Printf("error create temp dir")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer tempFile.Close()
	log.Println("Temp file success create")
	tempFile.Write(fileBytes)
	aes.EncryptLargeFiles(tempFile.Name(), tempFileName, []byte("12345678912345678912345678912345"))
	returnFileBytes, err := os.ReadFile(tempFileName)
	if err != nil {
		log.Printf("error crypto file")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(returnFileBytes)
	defer os.Remove(tempFileName)
	defer os.Remove(tempFile.Name())
	defer os.Remove(fileName)
}
