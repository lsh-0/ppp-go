package main

import (

	// "github.com/lsh-0/ppp-go/internal/components/lax-proxy/interface"
	"github.com/lsh-0/ppp-go/internal/http"
	"github.com/lsh-0/ppp-go/internal/types"
	"github.com/lsh-0/ppp-go/internal/utils"
	// "github.com/emvi/null"
)

func main() {
	url := "https://api.elifesciences.org/articles"
	output_fname := http.URLFilename(url)

	if !utils.FileExists(output_fname) {
		output_fname = http.Download(url)
	}

	article_list := utils.ReadJSON[types.ArticleList](output_fname)
	utils.Pprint(utils.ToJSON(article_list))
}
