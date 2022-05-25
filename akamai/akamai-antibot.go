package akamai

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SensorPayload struct {
	Abck      string `json:"abck"`
	Url       string `json:"url"`
	Bm_sz     string `json:"bm_sz"`
	UserAgent string `json:"userAgent"`
}
type V1SensorPayload struct {
	Abck      string  `json:"abck"`
	Url       string  `json:"url"`
	Version   float32 `json:"version"`
	UserAgent string  `json:"userAgent"`
}
type SensorResponse struct {
	Sensor    string `json:"sensor"`
	UserAgent string `json:"userAgent"`
}

func (c *Antibot) GenerateSensor(abck, url, bm_sz, userAgent string) (string, error) {
	payload := &SensorPayload{
		Abck:      abck,
		Url:       url,
		Bm_sz:     bm_sz,
		UserAgent: userAgent,
	}
	payloadBuffer := new(bytes.Buffer)
	json.NewEncoder(payloadBuffer).Encode(payload)

	request, err := http.NewRequest(http.MethodPost, "https://api.antibot.to/akamai/generate", payloadBuffer)
	if err != nil {
		return "", err
	}

	request.Header = http.Header{
		"content-type": {"application/json"},
		"apikey":       {c.apiKey},
	}

	response, err := c.client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var sensorResp SensorResponse
	json.Unmarshal(respBody, &sensorResp)

	return sensorResp.Sensor, nil

}

func (c *Antibot) GenerateV1Sensor(abck, url, userAgent string, version float32) (string, error) {
	payload := &V1SensorPayload{
		Abck:      abck,
		Url:       url,
		Version:   version,
		UserAgent: userAgent,
	}
	payloadBuffer := new(bytes.Buffer)
	json.NewEncoder(payloadBuffer).Encode(payload)

	request, err := http.NewRequest(http.MethodPost, "https://api.antibot.to/akamai/generate", payloadBuffer)
	if err != nil {
		return "", err
	}

	request.Header = http.Header{
		"content-type": {"application/json"},
		"apikey":       {c.apiKey},
	}

	response, err := c.client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var sensorResp SensorResponse
	json.Unmarshal(respBody, &sensorResp)

	return sensorResp.Sensor, nil

}
