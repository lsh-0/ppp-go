package api

// internal communication between components.

// whether the component is proxying the request to a remote API,
// or has a local implementation,
// it will receive a list of zero or more parameters followed by an `api.RequestConfig` and
// return an `api.Response`.

// an `api.Response` extends the builtin `http.Response` with api specific parsing.

import (
	"errors"
	"io"
	"mime"
	"net/http"
	"strconv"

	foo "github.com/lsh-0/ppp-go/internal/http"

	"github.com/lsh-0/ppp-go/internal/log"
)

type RequestConfig struct {

	// acceptable content types
	// an array of mime:mime-version
	ContentTypeList []map[string]int

	// api key, if any, for making authenticated requests.
	ApiKey string

	// a list of key+vals
	KeywordArgs map[string]any

	// trait, see 'paged'
	Page    int
	PerPage int
	Order   string

	// trait, see 'subjected'
	SubjectList []string

	// trait, see 'container'
	ContainingList []string
}

type Response struct {
	// https://pkg.go.dev/net/http#Response
	HttpResponse http.Response

	// the API may respond with an error
	Error bool

	// response body as a string, not bytes, regardless of content encoding.
	// we don't expect to receive any binary content from the API
	Content string

	// response body as a JSON string, but only if JSON-type mime response
	JSONContent string

	// aka the 'mime' type, "application/vnd.elife.article-list+json"
	ContentType string

	// mime type 'version' parameter, the 1 in "version=1"
	ContentTypeVersion int

	// if the content type version has been deprecated in favour of a newer version.
	// only happens if a specific, deprecated, content type version has been requested.
	ContentVersionDeprecated bool

	// the request was successfully authenticated if an api key was given
	Authenticated bool
}

// proxies a request from a http server to https://api.elifesciences.org
// writing the response directly to the given `respWriter`
func ProxyHttpRequest(respWriter http.ResponseWriter, extReq *http.Request) {

	extReq.URL.Scheme = "https"
	extReq.URL.Host = "api.elifesciences.org"

	intReq := http.Request{
		Method: extReq.Method,
		URL:    extReq.URL,
		Header: http.Header{},
		// POST+PUT
		// Body: extReq.Body,
		ContentLength: extReq.ContentLength,
		Close:         extReq.Close,
	}
	foo.CopyHeader(extReq.Header, intReq.Header)

	resp, err := http.DefaultTransport.RoundTrip(&intReq)
	if err != nil {
		http.Error(respWriter, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	foo.CopyHeader(resp.Header, respWriter.Header())
	respWriter.WriteHeader(resp.StatusCode)
	io.Copy(respWriter, resp.Body)
}

/* extracts the content type ('mime') and it's version parameter from the given `content_type_mime`.
 */
func ParseContentType(content_type_mime string) (string, int, error) {
	if content_type_mime == "" {
		return "", 0, errors.New("empty mime")
	} else {
		content_type, parameter_map, error := mime.ParseMediaType(content_type_mime)
		if error != nil {
			return "", 0, error
		}
		parameter_version := parameter_map["version"]
		// it's ok for responses to exclude a version parameter.
		// for example, 'application/problem+json' has no 'version' parameter.
		if parameter_version == "" {
			return content_type, 0, nil
		}
		content_type_version, error := strconv.Atoi(parameter_version)
		if error != nil {
			return "", 0, error
		}
		return content_type, content_type_version, nil
	}
}

// makes a request to the api gateway for the given `endpoint` path,
// returning an `api.Response` that extends the built-in `http.Response`.
func Request(endpoint string, opts RequestConfig) Response {
	url := "https://api.elifesciences.org" + endpoint

	// ensure url is valid?
	// ensure `endpoint` starts with '/' ?

	resp, error := http.Get(url)
	if error != nil {
		log.Error("failed to fetch URL '", url, "' with error: ", error)
	} else {
		defer resp.Body.Close()
	}

	content_bytes, _ := io.ReadAll(resp.Body)

	response_content_type := resp.Header.Get("Content-Type")
	content_type, content_type_version, error := ParseContentType(response_content_type)
	if error != nil {
		log.Warn("failed to correctly parse content type", response_content_type)
	}

	in_error := content_type == "application/problem+json"

	return Response{
		HttpResponse:             *resp,
		Error:                    in_error,
		Content:                  string(content_bytes),
		// JSONContent:           ...
		ContentType:              content_type,
		ContentTypeVersion:       content_type_version,
		ContentVersionDeprecated: false,
		Authenticated:            false,
	}
}
