package main

import (
	"embed"
	"net/http"
)

//go:embed ui
var uiFS embed.FS

func main() {
	a := http.FileServer(http.FS(uiFS))
	http.Handle("/", a)
	http.Handle("/ui", http.StripPrefix("/ui", http.FileServer(http.FS(uiFS))))

	http.ListenAndServe(":8080", nil)
}
