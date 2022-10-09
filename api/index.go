package api

import "net/http"

func InitRouter(store *Store) {
	html := http.FileServer(http.Dir("./frontend/dist"))
	http.HandleFunc("/datas", store.Datas)
	http.HandleFunc("/action", store.Action)
	http.Handle("/", html)
}
