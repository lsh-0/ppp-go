package recommendations

import (
	"errors"

	"github.com/lsh-0/ppp-go/internal/api"
	"github.com/lsh-0/ppp-go/internal/components/lax"
	"github.com/lsh-0/ppp-go/internal/components/search"
	"github.com/tidwall/gjson"
)

// searches for article content matching the subject of the article with the given `article_id`.
// returns a list of `application/vnd.elife.search+json` json strings
// each string is probably an article POR or VOR snippet:
// https://github.com/elifesciences/api-raml/blob/b0fcfa2b76b61b71ab830bcd41929f1758375851/src/model/article-list.v1.yaml#L14-L15
func find_articles_by_subject(article_id int64, lax_api lax.IAPI, search_api search.IAPI) ([]string, error) {

	empty_response := []string{}

	request_config := api.RequestConfig{}
	lax_api_response, error := lax_api.Article(article_id, request_config)
	if error != nil {
		return empty_response, error
	}

	subject_id := gjson.Get(lax_api_response.JSONContent, "subjects.0.id").String()
	if subject_id == "" {
		return empty_response, nil
	}
	subject_list := []string{subject_id}

	search_term := ""
	sort_by := "date"
	search_content_types := []string{
		"editorial", "feature", "insight", "research-advance", "research-article",
		"research-communication", "registered-report", "replication-study",
		"review-article", "scientific-correspondence", "short-report", "tools-resources"}
	request_config = api.RequestConfig{}
	search_api_response, error := search_api.Search(search_term, sort_by, search_content_types, subject_list, request_config)
	if error != nil {
		return empty_response, error
	}

	var result_list []string

	// remove self from results
	search_result_list := gjson.Get(search_api_response.JSONContent, "items").Array()
	if len(search_result_list) == 0 {
		return empty_response, errors.New("bad search response, 'items' not found")
	}
	for _, search_result := range search_result_list {
		search_result_json := search_result.String()
		search_result_id := gjson.Get(search_result_json, "id").Int()
		if search_result_id == article_id {
			continue
		}
		result_list = append(result_list, search_result_json)
	}

	return result_list, nil
}
