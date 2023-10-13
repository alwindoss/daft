package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:ui
var uiFS embed.FS

func main() {
	staticFS := fs.FS(uiFS)
	htmlContent, err := fs.Sub(staticFS, "ui")
	if err != nil {
		log.Fatal(err)
	}
	renderFS := http.FileServer(http.FS(htmlContent))
	http.Handle("/", renderFS)

	http.ListenAndServe(":8080", nil)
}
