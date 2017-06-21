package gw2

import "encoding/json"

// PrettyPrint will pretty print any variable in JSON format
func PrettyPrint(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	println(string(b))
}
