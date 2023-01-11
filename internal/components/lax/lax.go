package lax

import (
	"strconv"

	"github.com/lsh-0/ppp-go/internal/api"
)

type IAPI interface {
	ArticleList(opts api.RequestConfig) api.Response
	Article(id int64, opts api.RequestConfig) (api.Response, error)
	RelatedArticleList(id int64, opts api.RequestConfig) api.Response
	ArticleVersionList(id int64, opts api.RequestConfig) api.Response
	ArticleVersion(id int64, version int, opts api.RequestConfig) api.Response
}

type Lax struct {
	IAPI
}

type LaxProxy struct {
	IAPI
}

// proxy results from a remote Lax 'article-list' endpoint
func (p LaxProxy) ArticleList(opts api.RequestConfig) api.Response {
	return api.Request("/articles", opts)
}

func (p LaxProxy) Article(id int, opts api.RequestConfig) (api.Response, error) {
	idstr := strconv.Itoa(id)
	return api.Request("/articles/"+idstr, opts), nil
}

func (p LaxProxy) RelatedArticleList(id int, opts api.RequestConfig) api.Response {
	idstr := strconv.Itoa(id)
	return api.Request("/articles/"+idstr+"/related", opts)
}

func (p LaxProxy) ArticleVersionList(id int, opts api.RequestConfig) api.Response {
	idstr := strconv.Itoa(id)
	return api.Request("/articles/"+idstr+"/versions", opts)
}

func (p LaxProxy) ArticleVersion(id int, version int, opts api.RequestConfig) api.Response {
	idstr := strconv.Itoa(id)
	versionstr := strconv.Itoa(version)
	return api.Request("/articles/"+idstr+"/versions/"+versionstr, opts)
}

// a local implementation of Lax's 'article-list' endpoint
func (p Lax) ArticleList(opts api.RequestConfig) api.Response {
	/* TODO, obvs. */
	return *new(api.Response)
}
