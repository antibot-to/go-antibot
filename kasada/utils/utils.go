package utils

import (
	"errors"
	"regexp"
)

func GetScript(body []byte) (string, error) {

	re := `src="(?P<src>.+)"`
	params := RegParams(re, string(body))
	if src, ok := params["src"]; ok {
		return src, nil
	}

	return "", errors.New("script not found")
}

func RegParams(regEx, rawsting string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(rawsting)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if name == "" {
			continue
		}
		if i > 0 && i <= len(match) {
			if match[i] == "" {
				continue
			}
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
