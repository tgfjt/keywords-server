package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/tgfjt/keywords-server/src"
	"google.golang.org/appengine"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.Handle("/playground", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(keywords_server.NewExecutableSchema(keywords_server.Config{Resolvers: &keywords_server.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	appengine.Main()
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{Status: "ok", Message: "Hello world."})
}
