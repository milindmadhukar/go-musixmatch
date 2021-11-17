package gomusixmatch

import (
	"net/url"

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
