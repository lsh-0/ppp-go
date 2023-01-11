package search

import (
	"github.com/lsh-0/ppp-go/internal/api"
)

type IAPI interface {
	Search(search_term string, sort_by string, search_content_types []string, subject_list []string, opts api.RequestConfig) (api.Response, error)
}

type Search struct {
	IAPI
}

type SearchProxy struct {
	IAPI
}
