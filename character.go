package gw2

import (
	"encoding/json"
	"io"
	"log"
	"net/url"
	"strings"
)

// GetDetails will gather details of a given character in an Account struct.
func (ch *Character) GetDetails(c *Client) {
	v := Caller(&c.Config, "v2/characters/"+url.PathEscape(ch.Core.Name)+"/core")
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
