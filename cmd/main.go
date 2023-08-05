package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/arpithku/URL_Shortener_1/pkg/models"
	"github.com/arpithku/URL_Shortener_1/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello MAC")
	val := models.GetURL(context.Background(), "testKey_2")
	fmt.Println(val)
	r := mux.NewRouter()
	routes.SetUpRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
