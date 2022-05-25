package kasada

import "net/http"

type Antibot struct {
	apiKey string
	client *http.Client
}

func NewKasadaAntibot(apiKey string) *Antibot {
	return &Antibot{apiKey: apiKey, client: &http.Client{}}
}
