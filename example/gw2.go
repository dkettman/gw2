package main

import (
	"fmt"
	"os"

	"github.com/dkettman/gw2"
)

func main() {

	cfg, cfgErr := gw2.LoadConfig("config.json")

	if cfgErr != nil {
		fmt.Printf("%v", cfgErr)
		os.Exit(1)
	}

	c := gw2.NewClient(cfg)

	err := c.GetAccountInfo()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	gw2.PrettyPrint(c)
}
