package conduit

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// StartServer starts the API server
func StartServer() {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatal("no dev.env file found")
	}
	port := os.Getenv("HTTP_PORT")
	router := mux.NewRouter()
	createRoutes(router)
	log.Printf("Running Server on Port: %s", port)

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
