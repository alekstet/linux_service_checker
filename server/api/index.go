package api

import "net/http"

func InitRouter(store *store) {
	http.HandleFunc("/collect", store.Get)
	http.HandleFunc("/make", store.Make)
}
