package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/arpithku/URL_Shortener_1/pkg/models"
	"github.com/gorilla/mux"
)

func PostShortenURL(w http.ResponseWriter, r *http.Request) {
	timeoutContext, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	params := mux.Vars(r)
	key := params["key"]
	value := params["value"]
	models.SetURL(timeoutContext, key, value)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
}

func GetShortenURL(w http.ResponseWriter, r *http.Request) {
	timeoutContext, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	params := mux.Vars(r)
	key := params["key"]
	value := models.GetURL(timeoutContext, key)
	fmt.Println(value)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(value)
	w.Write(res)
}

func DeleteShortenURL(w http.ResponseWriter, r *http.Request) {
	timeoutContext, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	params := mux.Vars(r)
	key := params["key"]
	value := models.GetURL(timeoutContext, key)
	fmt.Println(value)

}

func UpdateShortenURL(w http.ResponseWriter, r *http.Request) {
	timeoutContext, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	params := mux.Vars(r)
	key := params["key"]
	value := params["value"]
	err := models.UpdateURL(timeoutContext, key, value)
	w.Header().Set("Content-Type", "pkglication/json")
	if err != nil {
		if err.Error() == "Key does not exist" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Key does not exist"))
			return
		}

	}
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(value)
	w.Write(res)
}
