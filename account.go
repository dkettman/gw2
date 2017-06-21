package gw2

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"strings"
)

func (a *Account) GetAccountInfo(c *Config) {
	val := Caller(c, "v2/account")
	dec := json.NewDecoder(strings.NewReader(string(val)))
	for {
		if err := dec.Decode(a); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	a.resolveWorld(c)
	a.getCharacters(c)
}

func (a *Account) resolveWorld(c *Config) {

	found := false
	val := Caller(c, "v2/worlds?ids="+fmt.Sprintf("%d", a.WorldID))
	dec := json.NewDecoder(strings.NewReader(string(val)))
	for dec.More() {
		t, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if t == "name" {
			found = true
			continue
		}
		if found {
			a.World = fmt.Sprintf("%v", t)
			break
		}
	}
}

func (a *Account) getCharacters(c *Config) {
	a.Characters = make(map[string]Character)
	val := Caller(c, "v2/characters")
	chars := make([]string, 0)
	json.Unmarshal(val, &chars)

	for _, ch := range chars {
		var char Character
		char.Core = CharacterCore{Name: ch}
		a.Characters[ch] = char

	}
}

func (ch *Character) GetDetails(c *Config) {
	v := Caller(c, "v2/characters/"+url.PathEscape(ch.Core.Name)+"/core")
	dec := json.NewDecoder(strings.NewReader(string(v)))
	for {
		var charcore CharacterCore
		if err := dec.Decode(&charcore); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		ch.Core = charcore
	}
}
