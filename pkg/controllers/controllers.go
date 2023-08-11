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
	var req models.RequestResponseBody
	timeoutContext, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	key := req.URL
	req.Shortened_URL = models.SetURL(timeoutContext, key)
	req.Shortened_URL = "http://localhost:8080/shorten/" + req.Shortened_URL
	w.Header().Set("Content-Type", "pkglication/json")
	output, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	//w.WriteHeader(http.StatusOK)
	w.Write(output)
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
	err := models.DeleteURL(timeoutContext, key)
	if err != nil {
		if err.Error() == "url does not exist" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("URL does not exist"))
			return
		}
		panic(err)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write([]byte(" http://localhost:8080/shorten/" + key + " Deleted Succesfully"))
}

func UpdateShortenURL(w http.ResponseWriter, r *http.Request) {

	var req models.RequestResponseBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	key := params["key"]
	url := req.URL
	timeoutContext, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	err = models.UpdateURL(timeoutContext, key, url)
	if err != nil {
		if err.Error() == "url does not exist" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("URL does not exist"))
			return
		}
		panic(err)
	}
	w.Header().Set("Content-Type", "pkglication/json")
	req.Shortened_URL = "http://localhost:8080/shorten/" + key
	w.WriteHeader(http.StatusOK)
	output, _ := json.Marshal(req)
	w.Write(output)
}
