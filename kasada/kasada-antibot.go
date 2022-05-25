package kasada

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"io"
	"net/http"
	Url "net/url"
	"strings"
)

type CdPayload struct {
	Time int `json:"time"`
}

func (c *Antibot) GenerateCD() (string, error) {
	request, err := http.NewRequest(http.MethodGet, "https://api.antibot.to/kasada/x-kpsdk-cd", nil)
	if err != nil {
		return "", err
	}
	request.Header = http.Header{
		"apikey":       {c.apiKey},
		"content-type": {"application/json"},
	}
	response, err := c.client.Do(request)
	if err != nil {
		return "", err
	}
	responseByte, err := io.ReadAll(response.Body)

	return string(responseByte), nil
}

func (c *Antibot) GenerateCustomCD(duration int) (string, error) {
	payload := &CdPayload{
		Time: duration,
	}
	payloadBuffer := new(bytes.Buffer)
	json.NewEncoder(payloadBuffer).Encode(payload)

	request, err := http.NewRequest(http.MethodPost, "https://api.antibot.to/kasada/x-kpsdk-cd", payloadBuffer)
	if err != nil {
		return "", err
	}
	request.Header = http.Header{
		"apikey":       {c.apiKey},
		"content-type": {"application/json"},
	}
	response, err := c.client.Do(request)
	if err != nil {
		return "", err
	}
	responseByte, err := io.ReadAll(response.Body)
	return string(responseByte), nil

}
func (c *Antibot) GenerateFullCD() (string, error) {

	request, err := http.NewRequest(http.MethodPost, "https://api.antibot.to/kasada/x-kpsdk-cd", strings.NewReader("{}"))
	if err != nil {
		return "", err
	}
	request.Header = http.Header{
		"apikey":       {c.apiKey},
		"content-type": {"application/json"},
	}
	response, err := c.client.Do(request)
	if err != nil {
		return "", err
	}
	responseByte, err := io.ReadAll(response.Body)
	return string(responseByte), nil

}

func (c *Antibot) SolveTl(resp *http.Response, url, ua, platform string) (*bytes.Reader, error) {
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var buf bytes.Buffer
	z := zlib.NewWriter(&buf)

	if _, err = z.Write(respBytes); err != nil {
		return nil, err
	}
	if err = z.Close(); err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, "https://api.antibot.to/kasada/x-kpsdk-ct", &buf)
	if err != nil {
		return nil, err
	}

	request.Header = http.Header{
		"content-type": {"application/octet-stream"},
		"apikey":       {c.apiKey},
	}

	params := Url.Values{}
	params.Add("url", url)
	params.Add("ua", ua)
	params.Add("platform", platform)
	request.URL.RawQuery = params.Encode()

	SolvedResp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer SolvedResp.Body.Close()

	bodyBytes, err := io.ReadAll(SolvedResp.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(bodyBytes), nil

}
