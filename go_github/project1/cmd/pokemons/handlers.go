package main

import (
	"AliZhumataev_go_project/pkg/pokemons"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Pokemons []pokemons.Pokemon `json:"pokemons"`
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	func() {
		response, err := json.Marshal(interface{}(
			map[string]string{
				"404 status": "ok",
				"info":       pokemons.Info(),
			},
		))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
			return
		}
		w.Header().Set("Content", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}()
}

func PokemonsShow(w http.ResponseWriter, r *http.Request) {
	poki := pokemons.GetPokemons()
	respondWithJSON(w, http.StatusOK, poki)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Greetings Page! Available paths: \n /health-check \n /pokemons \n /pokemons/name")
}

func PokemonShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	poke, err := pokemons.GetPokemon(name)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Sorry, there is no such Pokemon name")
		return
	}

	respondWithJSON(w, http.StatusOK, poke)
}
