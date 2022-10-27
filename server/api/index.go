package api

import "net/http"

func InitRouter(store *store) {
	http.HandleFunc("/collect", store.Collect)
	http.HandleFunc("/make", store.Make)
}
