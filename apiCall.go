package gw2

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Caller will return a list of all registered world names
func Caller(config *Config, endpoint string) []byte {

	log.Printf("Endpoint: %s\n", endpoint)
	c := &http.Client{}
	req, err := http.NewRequest("GET", config.BaseURL+endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+config.APIKey)
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return result
}
