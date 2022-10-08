package http

import (
	"encoding/base32"
	"io"
	"net/http"
	"os"

	"github.com/lsh-0/ppp-go/internal/utils"
)

func URLFilename(url string) string {
	return base32.StdEncoding.EncodeToString([]byte(url))
}

// downloads `url` to a base32 (fs safe) filename
func Download(url string) string {
	resp, err := http.Get(url)
	utils.Check(err)
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	utils.Check(err)

	output_fname := URLFilename(url)
	err = os.WriteFile(output_fname, bytes, 0644)
	utils.Check(err)

	return output_fname
}
