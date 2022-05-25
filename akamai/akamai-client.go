package akamai

import "net/http"

type Antibot struct {
	apiKey string
	client *http.Client
}

func NewAkamaiAntibot(apiKey string) *Antibot {
	return &Antibot{apiKey: apiKey, client: &http.Client{}}
}
