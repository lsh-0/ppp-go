package http

// HTTP utils.
// nothing to do with APIs, types, components, projects etc.

import (
	"encoding/base32"
	"io"
	"net/http"
	"os"

	"github.com/lsh-0/ppp-go/internal/utils"
)

// copies one set of headers into another.
// todo: if not re-usable, move back into api.go
// todo: opportunity here to add/remove headers
func CopyHeader(src, dest http.Header) {
	for header, value_list := range src {
		for _, value := range value_list {
			dest.Add(header, value)
		}
	}
}

func URLFilename(url string) string {
	return base32.StdEncoding.EncodeToString([]byte(url))
}

// downloads `url` to a byte array
// why would anyone want status or headers or such? pft
func Download(url string) []byte {
	resp, err := http.Get(url)
	utils.Check(err)
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	utils.Check(err)

	return bytes
}

// downloads `url` to a base32 (fs safe) filename
func DownloadToFile(url string) string {
	output_fname := URLFilename(url)
	err := os.WriteFile(output_fname, Download(url), 0644)
	utils.Check(err)
	return output_fname
}

// downloads `url` returning an instance of the given type `T`
// for example: DownloadJSON[ArticleList]("https://api.elifesciences.org/articles")
func DownloadJSON[T any](url string) T {
	return utils.FromJSON[T](Download(url))
}
