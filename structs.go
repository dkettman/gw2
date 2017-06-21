package gw2

type Account struct {
	Name       string `json:"name"`
	WorldID    int    `json:"world"`
	World      string
	Characters map[string]Character
}

// Config struct holds configuration options for GW2 API
type Config struct {
	APIKey  string `json:"apiKey"`
	BaseURL string `json:"baseURL"`
}

type Character struct {
	Core CharacterCore
}

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

// World is a structure that holds the ID and Name of each defined world.
type World struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Population string `json:"population"`
}
