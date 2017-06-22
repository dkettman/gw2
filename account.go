package gw2

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

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

// Currency Is in the Account struct to hold currency values and types
type Currency struct {
	ID    int `json:"id"`
	Value int `json:"value"`
}

// GetAccountInfo accesses the API and gathers basic account-based information:
// user name, world, character list, wallet
func (c *Client) GetAccountInfo() error {
	val := Caller(&c.Config, "v2/account")
	dec := json.NewDecoder(strings.NewReader(string(val)))
	for {
		if err := dec.Decode(&c.Account); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}

	c.resolveWorld()
	c.getCharacters()
	c.getWallet()
	return nil
}

func (c *Client) resolveWorld() error {

	found := false
	val := Caller(&c.Config, "v2/worlds?ids="+fmt.Sprintf("%d", c.Account.WorldID))
	dec := json.NewDecoder(strings.NewReader(string(val)))
	for dec.More() {
		t, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if t == "name" {
			found = true
			continue
		}
		if found {
			c.Account.World = fmt.Sprintf("%v", t)
			break
		}
	}
	return nil
}

func (c *Client) getCharacters() error {
	c.Account.Characters = make(map[string]Character)
	val := Caller(&c.Config, "v2/characters")
	chars := make([]string, 0)
	json.Unmarshal(val, &chars)

	for _, ch := range chars {
		var char Character
		char.Core = CharacterCore{Name: ch}
		c.Account.Characters[ch] = char
	}
	return nil
}

func (c *Client) getWallet() error {
	val := Caller(&c.Config, "v2/account/wallet")
	dec := json.NewDecoder(strings.NewReader(string(val)))
	for {
		if err := dec.Decode(&c.Account.Wallet); err == io.EOF {
			break
		} else if err != nil {
			return err
		}

	}
	return nil
}
