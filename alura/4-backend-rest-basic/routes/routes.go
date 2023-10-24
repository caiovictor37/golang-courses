package routes

import (
	"go-rest-api/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.SelectAllPersonalities)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
