package main

import (
	"fmt"
	"log"
	dbg "runtime/debug"

	"github.com/gorilla/mux"

	"movie-rating-api/db"
	movieHttp "movie-rating-api/http"

	"github.com/gorilla/handlers"
	"github.com/rs/cors"

	"net/http"
)

var r = mux.NewRouter()

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("ERROR: %v\n", err)
			dbg.PrintStack()
			log.Fatalf("FATAL: %v\n", err)
		}
	}()

	log.Print("******* MOVIE RATING API *******")
	const port = "8080"

	err := db.InitializeDB()
	if err != nil {
		log.Fatalln(fmt.Sprintf("failed to initialize db: %s\n", err.Error()))
	}

	client := db.NewDBCLient(nil)

	err = db.InitializeMovies(client)
	if err != nil {
		log.Fatalln(fmt.Sprintf("failed to initialize movies: %s\n", err.Error()))
	}

	movieHttp.ConfigureRouter(r)

	log.Printf("starting api on port %s\n", port)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})

	corsObj := handlers.AllowedOrigins([]string{"*"})
	handler := c.Handler(handlers.CORS(corsObj)(r))

	// start http server
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	if err != nil {
		fmt.Printf("failed to listen and service with err: %s\n", err.Error())
	}
}
