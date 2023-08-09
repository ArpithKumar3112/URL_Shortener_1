package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arpithku/URL_Shortener_1/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello MAC")
	//val := models.GetURL(context.Background(), "testKey_2")
	//fmt.Println(val)
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello from HandleFunc")
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal("Hello from HandleFunc")
		w.Write(res)
	}).Methods("POST")
	routes.SetUpRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
