package api

import "net/http"

func InitRouter(store *Store) {
	html := http.FileServer(http.Dir("./frontend/dist"))
	http.HandleFunc("/collect", store.Collect)
	http.HandleFunc("/make", store.Make)
	http.Handle("/", html)
}
