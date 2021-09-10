package gomusixmatch

import (
	"net/url"

	musixmatchParams "github.com/milindmadhukar/go-musixmatch/params"
)

func processParams(params ...musixmatchParams.Param) (musixmatchParams.Params, error) {
	p := musixmatchParams.Params{
		UrlParams: url.Values{},
	}
	for _, param := range params {
		param(&p)
		if p.Err != nil {
			return musixmatchParams.Params{
				UrlParams: url.Values{},
			}, p.Err
		}
	}

	return p, nil
}
