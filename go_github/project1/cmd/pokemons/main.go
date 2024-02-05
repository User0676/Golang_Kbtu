package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":8080"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health-check", healthCheck)
	r.HandleFunc("/pokemons", PokemonsShow)
	r.HandleFunc("/pokemons/{name}", PokemonShow)
	r.HandleFunc("/", HomePage)

	log.Printf("Starting server on http://localhost%s\n", PORT)
	err := http.ListenAndServe(PORT, r)
	log.Fatal(err)

}
