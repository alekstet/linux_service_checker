package api

import "net/http"

const frontendPath = "./frontend/dist"

func InitRouter(store *store) {
	html := http.FileServer(http.Dir(frontendPath))
	http.HandleFunc("/make", store.Make)
	http.HandleFunc("/collect", store.Collect)
	http.Handle("/", html)
}
