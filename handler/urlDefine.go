package handler

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func Route() http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(setJSONContentType)

	base := "/api/v1/"

	r.HandleFunc(base+"a/customer", customerAdd).Methods(http.MethodPost)
	r.HandleFunc(base+"u/customer", customerUpdate).Methods(http.MethodPost, http.MethodPut)
	r.HandleFunc(base+"d/customer", customerDelete).Methods(http.MethodPost, http.MethodDelete)
	r.HandleFunc(base+"g/customer", customerGet).Methods(http.MethodPost)
	r.HandleFunc(base+"l/customer", customerList).Methods(http.MethodGet)

	r.HandleFunc(base+"a/admin", adminAdd).Methods(http.MethodPost)
	r.HandleFunc(base+"u/admin", adminUpdate).Methods(http.MethodPost, http.MethodPut)
	r.HandleFunc(base+"d/admin", adminDelete).Methods(http.MethodPost, http.MethodDelete)
	r.HandleFunc(base+"g/admin", adminGet).Methods(http.MethodPost)
	r.HandleFunc(base+"l/admin", adminList).Methods(http.MethodGet)

	r.HandleFunc(base+"a/advert", advertAdd).Methods(http.MethodPost)
	r.HandleFunc(base+"u/advert", advertUpdate).Methods(http.MethodPost, http.MethodPut)
	r.HandleFunc(base+"d/advert", advertDelete).Methods(http.MethodPost, http.MethodDelete)
	r.HandleFunc(base+"g/advert", advertGet).Methods(http.MethodPost)
	r.HandleFunc(base+"l/advert", advertList).Methods(http.MethodGet)

	r.HandleFunc(base+"a/order", orderAdd).Methods(http.MethodPost)
	r.HandleFunc(base+"u/order", orderUpdate).Methods(http.MethodPost, http.MethodPut)
	r.HandleFunc(base+"d/order", orderDelete).Methods(http.MethodPost, http.MethodDelete)
	r.HandleFunc(base+"g/order", orderGet).Methods(http.MethodPost)
	r.HandleFunc(base+"l/order", orderList).Methods(http.MethodPost)

	r.HandleFunc(base+"a/log", logAdd).Methods(http.MethodPost)
	r.HandleFunc(base+"l/log", logList).Methods(http.MethodGet)

	r.HandleFunc(base+"l/history", historyList).Methods(http.MethodGet)

	r.HandleFunc(base+"login", login).Methods(http.MethodPost)
	r.HandleFunc(base+"admin/login", adminLogin).Methods(http.MethodPost)

	handler := cors.AllowAll().Handler(r)
	return handler
}

// todo: url'de id gibi bir bilginin taşınması güvenlik zafiyetine neden olabilir!
// todo: düzelt
