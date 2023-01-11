package lax

import (
	"testing"

	"github.com/lsh-0/ppp-go/internal/api"
)

type LaxDummy struct {
	IAPI
}

// a dummy implementation of an 'article-list' endpoint for testing
func (p LaxDummy) ArticleList(opts api.RequestConfig) api.Response {
	return api.Response{}
}

func TestArticleList(t *testing.T) {
	lax := LaxDummy{}
	opts := api.RequestConfig{}
	lax.ArticleList(opts)
}
