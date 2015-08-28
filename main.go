package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	file, header, err := r.FormFile("file")

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	defer file.Close()
	log.Println("Incoming: ", header.Filename)

	out, err := os.Create(path.Join("files", header.Filename))
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintf(w, "File uploaded successfully: ")
	fmt.Fprintf(w, header.Filename)
	log.Println("Received: ", header.Filename)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./form.html")
		return
	})
	http.Handle("/uploads/", http.StripPrefix("/files/", http.FileServer(http.Dir("uploads"))))
	http.HandleFunc("/post", uploadHandler)
	http.ListenAndServe(":8080", nil)
}
