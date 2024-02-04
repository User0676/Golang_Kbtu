package pokemons

import (
	"errors"
)

type Pokemon struct {
	Id        string `json:"id"`
	Name      string `name:"title"`
	Region    string `json:"region"`
	Type      string `json:type`
	CanEvolve bool   `json:evolve`
}

var Pokemons = []Pokemon{
	{
		Id:        "1",
		Name:      "Charmander",
		Region:    "Volcano",
		Type:      "fire",
		CanEvolve: true,
	},
	{
		Id:        "2",
		Name:      "Bulbasaur",
		Region:    "Forest",
		Type:      "grass",
		CanEvolve: true,
	},
	{
		Id:        "3",
		Name:      "Squirtle",
		Region:    "river",
		Type:      "water",
		CanEvolve: true,
	},
	{
		Id:        "4",
		Name:      "Charizard",
		Region:    "Volcano",
		Type:      "fire",
		CanEvolve: false,
	},
	{
		Id:        "5",
		Name:      "Venusaur",
		Region:    "Forest",
		Type:      "grass",
		CanEvolve: false,
	},
	{
		Id:        "6",
		Name:      "Blastoise",
		Region:    "ocean",
		Type:      "water",
		CanEvolve: false,
	},
	{
		Id:        "7",
		Name:      "Pikachu",
		Region:    "mountains",
		Type:      "electric",
		CanEvolve: true,
	},
	{
		Id:        "8",
		Name:      "Tyranitar",
		Region:    "canyon",
		Type:      "rock+dark",
		CanEvolve: false,
	},
	{
		Id:        "10",
		Name:      "Mewtwo",
		Region:    "laboratory",
		Type:      "psychic",
		CanEvolve: false,
	},
	{
		Id:        "11",
		Name:      "Gyrados",
		Region:    "Volcano",
		Type:      "water",
		CanEvolve: false,
	},
	{
		Id:        "12",
		Name:      "Aipom",
		Region:    "Forest",
		Type:      "normal",
		CanEvolve: true,
	},
	{
		Id:        "13",
		Name:      "Ali",
		Region:    "Home",
		Type:      "normal",
		CanEvolve: true,
	},
	{
		Id:        "14",
		Name:      "Raichu",
		Region:    "Mountains",
		Type:      "electric",
		CanEvolve: true,
	},
	{
		Id:        "15",
		Name:      "Steelix",
		Region:    "canyon",
		Type:      "rock+steel",
		CanEvolve: true,
	},
}

func GetPokemons() []Pokemon {
	return Pokemons
}

func GetPokemon(name string) (*Pokemon, error) {
	for _, r := range Pokemons {
		if r.Name == name {
			return &r, nil
		}
	}
	return nil, errors.New("Not such pokemon name")
}
