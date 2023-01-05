package lax

import (
	"strconv"

	"github.com/lsh-0/ppp-go/internal/api"
	"github.com/lsh-0/ppp-go/internal/types"
)

type API interface {
	ArticleList(opts api.RequestConfig) api.Response[types.ArticleSnippetList]
	Article(id int, opts api.RequestConfig) (api.Response[types.Article], error)
	RelatedArticleList(id int, opts api.RequestConfig) api.Response[types.ArticleSnippetList]
	ArticleVersionList(id int, opts api.RequestConfig) api.Response[types.ArticleSnippetList]
	ArticleVersion(id int, version int, opts api.RequestConfig) api.Response[types.Article]
}

type Lax struct {
	API
}

type LaxProxy struct {
	API
}

// proxy results from a remote Lax 'article-list' endpoint
func (p LaxProxy) ArticleList(opts api.RequestConfig) api.Response[types.ArticleSnippetList] {
	return api.Request[types.ArticleSnippetList]("/articles", opts)
}

func (p LaxProxy) Article(id int, opts api.RequestConfig) (api.Response[types.Article], error) {
	idstr := strconv.Itoa(id)
	return api.Request[types.Article]("/articles/"+idstr, opts), nil
}

func (p LaxProxy) RelatedArticleList(id int, opts api.RequestConfig) api.Response[types.ArticleSnippetList] {
	idstr := strconv.Itoa(id)
	return api.Request[types.ArticleSnippetList]("/articles/"+idstr+"/related", opts)
}

func (p LaxProxy) ArticleVersionList(id int, opts api.RequestConfig) api.Response[types.ArticleSnippetList] {
	idstr := strconv.Itoa(id)
	return api.Request[types.ArticleSnippetList]("/articles/"+idstr+"/versions", opts)
}

func (p LaxProxy) ArticleVersion(id int, version int, opts api.RequestConfig) api.Response[types.Article] {
	idstr := strconv.Itoa(id)
	versionstr := strconv.Itoa(version)
	return api.Request[types.Article]("/articles/"+idstr+"/versions/"+versionstr, opts)
}

// a local implementation of Lax's 'article-list' endpoint
func (p Lax) ArticleList(opts api.RequestConfig) api.Response[types.ArticleSnippetList] {
	/* TODO, obvs. */
	return *new(api.Response[types.ArticleSnippetList])
}
