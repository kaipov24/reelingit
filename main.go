package main

import (
	"log"
	"net/http"

	"github.com/kaipov24/reelingit/handlers"
	"github.com/kaipov24/reelingit/logger"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger %v", err)
	}
	defer logInstance.Close()
	return logInstance
}

func main() {

	logInstance := initializeLogger()

	movieHandler := handlers.MovieHandler{}

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)

	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8080"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
		logInstance.Error("Server failed", err)
	}
}
