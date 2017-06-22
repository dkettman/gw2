package gw2

// Client struct that holds both Config and Account for a given connection
type Client struct {
	Account Account
	Config  Config
}

// NewClient creates a new client which holds the config and all
// related account information
func NewClient(c Config) Client {
	return Client{Account: Account{}, Config: c}
}
