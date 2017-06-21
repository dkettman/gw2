package gw2

// structs.go Holds all of the structs used in the application.

// Account holds the base account data
type Account struct {
	Characters map[string]Character
	Name       string `json:"name"`
	Wallet     []Currency
	WorldID    int `json:"world"`
	World      string
}

// Character struct holds data specific to each Character in an Account
type Character struct {
	Core CharacterCore
}

// CharacterCore holds data from /v2/characters/?/core
type CharacterCore struct {
	Name       string `json:"name"`
	Race       string `json:"race"`
	Gender     string `json:"gender"`
	Profession string `json:"profession"`
	Level      int    `json:"level"`
	Guild      string `json:"guild"`
	Age        int    `json:"age"`
	Created    string `json:"created"`
	Deaths     int    `json:"deaths"`
	Title      int    `json:"title"`
}

// Config struct holds configuration options for GW2 API
type Config struct {
	APIKey  string `json:"apiKey"`
	BaseURL string `json:"baseURL"`
}

// Currency Is in the Account struct to hold currency values and types
type Currency struct {
	ID    int `json:"id"`
	Value int `json:"value"`
}

// World is a structure that holds the ID and Name of each defined world.
type World struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Population string `json:"population"`
}
