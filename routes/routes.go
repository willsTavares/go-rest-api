package routes

import (
	"log"
	"net/http"

	_ "github.com/willsTavares/go-rest-api/docs"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/willsTavares/go-rest-api/controllers"
	"github.com/willsTavares/go-rest-api/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	// Swagger
	r.PathPrefix("/").Handler(httpSwagger.WrapHandler)

	// API routes
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.PersonalityById).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonalityById).Methods("DELETE")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonalityById).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
