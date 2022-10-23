package api

import "net/http"

func InitRouter(store *store) {
	http.HandleFunc("/make", store.Make)
}
