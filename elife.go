package main

import (
	"encoding/base32"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	// "github.com/emvi/null"
)

func pprint(val interface{}) {

	switch valtype := val.(type) {
	// list of somethings
	case []interface{}:
		pprint("list of somethings")
		for i, v := range valtype {
			fmt.Print("index:", i, ": ")
			pprint(v)
		}

	// map of somethings
	case map[string]interface{}:
		pprint("maps of somethings")
		for k, v := range valtype {
			fmt.Print("key:", k, ": ")
			pprint(v)
		}

	case string:
		fmt.Println(val)

	default:
		fmt.Println("unknown:", val)
	}

}

// mapping jsonschema to go structs:
// 'allOf' == condense sets of properties, for example, 'Article'
// 'oneOf' == variable types, when type is array. for example, 'ArticlePoa' or 'ArticleVor'
// 'anyOf' == condense properties but nilable, for example, 'Image.thumbnail' and 'Image.social'.

type Copyright struct {
	Statement string `json:"statement"`
	License   string `json:"license"`
	Holder    string `json:"holder,omitempty"`
}

type Paragraph struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Image struct {
	Id   string `json:"id"`
	Type string `json:"image"`
	//Image string // not really, it's actually a 'misc/image'
}

type Subject struct {
	Id              string      `json:"id"`
	Name            string      `json:"name"`
	ImpactStatement string      `json:"impactStatement,omitempty"`
	AimsAndScope    []Paragraph `json:"aimsAndScope,omitempty"`
	Image           *struct {
		Banner    Image `json:"banner"`
		Thumbnail Image `json:"thumbnail"`
	} `json:"image,omitempty"`
}

type Article struct {
	// https://github.com/elifesciences/api-raml/blob/develop/src/snippets/article.v1.yaml
	Id                string    `json:"id"`
	Version           int       `json:"version"`
	Type              string    `json:"type"`
	Doi               string    `json:"doi"`
	AuthorLine        string    `json:"authorLine"`
	Title             string    `json:"title"`
	TitlePrefix       string    `json:"titlePrefix,omitempty"`
	Stage             string    `json:"stage"`       // 'preview' or 'published'
	Published         string    `json:"published"`   // date-time
	VersionDate       string    `json:"versionDate"` // date-time
	StatusDate        string    `json:"statusDate"`  // date-time
	Volume            int       `json:"volume"`
	ElocationId       string    `json:"elocationId"`
	Pdf               string    `json:"pdf"` // url
	Subjects          []Subject `json:"subjects"`
	ResearchOrganisms []string  `json:"researchOrganisms,omitempty"` // set
	Image             *struct {
		Thumbnail Image `json:"thumbnail,omitempty"`
		Social    Image `json:"social,omitempty"`
	} `json:"image,omitempty"`

	// present but shouldn't be?
	Copyright Copyright `json:"copyright"`
}

type ArticlePoa struct {
	Article
	Status string `json:"status"`
}

type ArticleVor struct {
	// https://github.com/elifesciences/api-raml/blob/develop/src/snippets/article.v1.yaml
	Article

	// https://github.com/elifesciences/api-raml/blob/2.8.0/src/snippets/article-vor.v1.yaml
	Status          string `json:"status"` // "vor"
	ImpactStatement string `json:"impactStatement,omitempty"`
	FiguresPDF      string `json:"figuresPdf,omitempty"` // url
	// not yet implemented?
	// ReviewedDate    string `json:"reviewedDate"` // date-time
	// CurationLabels  []string `json:"curationLabels"` // set actually, min of one
}

type ArticleList struct {
	Total int `json:"total"`
	// https://github.com/elifesciences/api-raml/blob/develop/dist/model/article-list.v1.json#L14
	Items []ArticleVor `json:"items"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func url_filename(url string) string {
	return base32.StdEncoding.EncodeToString([]byte(url))
}

// downloads `url` to a base32 (fs safe) filename
func download(url string) string {
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	check(err)

	output_fname := url_filename(url)
	err = os.WriteFile(output_fname, bytes, 0644)
	check(err)

	return output_fname
}

func read_json[T any](filename string) T {
	bytes, err := os.ReadFile(filename)
	check(err)

	i := new(T)
	json.Unmarshal(bytes, &i)

	return *i
}

func to_json(data interface{}) string {
	b, err := json.Marshal(data)
	check(err)
	return string(b[:])
}

func file_exists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}

func main() {
	url := "https://api.elifesciences.org/articles"
	output_fname := url_filename(url)

	if !file_exists(output_fname) {
		output_fname = download(url)
	}

	article_list := read_json[ArticleList](output_fname)
	pprint(to_json(article_list))
}
