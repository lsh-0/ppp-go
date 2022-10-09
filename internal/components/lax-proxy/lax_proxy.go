package lax

import (
	"github.com/lsh-0/ppp-go/internal/api"
	"github.com/lsh-0/ppp-go/internal/types"
)

func ArticleList2(opts api.Request) api.Response[types.ArticleList] {
	return api.Response[types.ArticleList]{}
}

func ArticleList(opts api.Request) types.ArticleList {
	return types.ArticleList{} // todo: more of an ArticleSnippetList
}

func Article(id int, opts api.Request) types.Article {
	return types.Article{}
}

func RelatedArticleList(id int, opts api.Request) types.ArticleList {
	return types.ArticleList{} // todo: more of an ArticleSnippetList
}

func ArticleVersionList(id int, opts api.Request) types.ArticleList {
	return types.ArticleList{} // todo: more of an ArticleSnippetList
}

func ArticleVersion(id int, version int, opts api.Request) types.Article {
	return types.Article{}
}
