// types.go
package types

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
