package gomusixmatch

import (
	"net/url"

	musixmatchParams "github.com/milindmadhukar/go-musixmatch/params"
)

func processParams(musixMatchUrl string, params ...musixmatchParams.Param) (string, error) {
	p := musixmatchParams.Params{
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
