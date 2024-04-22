package gomusixmatch

import (
	"net/url"
	"os"
	"strings"

	mxmParams "github.com/milindmadhukar/go-musixmatch/params"
)

func processParams(musixMatchUrl string, params ...mxmParams.Param) (string, error) {
	p := mxmParams.Params{
		UrlParams: url.Values{},
	}
	for _, param := range params {
		param(&p)
		if p.Err != nil {
			return "", p.Err
		}
	}

	urlParams := p.UrlParams.Encode()

	if urlParams != "" {
		musixMatchUrl += "&" + urlParams
	}

	return musixMatchUrl, nil
}

func GetApiKeyFromEnvFile(filename string) string {

	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	content := string(file)
	lines := strings.Split(content, "\n")

	for _, line := range lines {

		if strings.Contains(line, "MUSIXMATCH_API_KEY") {
			apiKey := strings.Split(line, "=")
			return apiKey[1]
		}
	}

	panic("API Key not found in the file")
}
