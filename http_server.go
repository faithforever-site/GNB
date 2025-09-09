package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func startHTTPServer() {
	os.MkdirAll("upload", os.ModePerm)

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/download/", downloadHandler)

	log.Println("HTTP server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "upload.html")
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	dst, _ := os.Create(filepath.Join("upload", header.Filename))
	defer dst.Close()

	_, _ = dst.ReadFrom(file)
	w.Write([]byte("Upload successful"))
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[len("/download/"):]
	http.ServeFile(w, r, filepath.Join("upload", filename))
}
