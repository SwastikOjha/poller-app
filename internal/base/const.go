package base

import "bytes"

const (
	ApiURI     string = "https://v6.exchangerate-api.com/v6/"
	ConfigPath string = "./cmd/config/config.json"
)

var (
	Buf bytes.Buffer
)
