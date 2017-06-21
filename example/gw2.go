package main

import "github.com/dkettman/gw2"

// Config stores the configuration for the application Right now, there are two required objects required:
// - apiKey: <String> - API key generated in the Account page of guildwars2.com
// - baseURL: <String> - baseURL to use for API access. Currently only 'https://api.guildwars2.com/'
var Config gw2.Config = gw2.LoadConfig("config.json")

func main() {

	account := gw2.Account{}
	account.GetAccountInfo(&Config)

	gw2.PrettyPrint(account)
}
