package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willsTavares/go-rest-api/controllers"
	"github.com/willsTavares/go-rest-api/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.PersonalityById).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonalityById).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonalityById).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
