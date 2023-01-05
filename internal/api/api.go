package api

// internal communication between components uses a common function interface of:
//  `(endpoint, options)` or `(endpoint, param1, param2, paramN, options)`
// returning a *http-like* response.
// the http-like response uses a status code, a type instance of the body content if
// content type was successfully negotiated, etc.

import (
	"io"
	"net/http"

	foo "github.com/lsh-0/ppp-go/internal/http"
	"github.com/lsh-0/ppp-go/internal/utils"
)

type ContentType struct {
	// aka the 'mime' type
	ContentType string
	Version     int
	// a content type is deprecated if it's not using the most recent content type *version*
	Deprecated bool
}

// does Go have tuples? do I really need a struct here?
type KeyVal struct {
	Key string
	Val any
}

type RequestConfig struct {
	// when true, response is not decompressed nor is the json marshalled.
	// you can serve the bytes up directly to the response writer.
	Proxy bool

	// acceptable content types.
	ContentTypeList []ContentType
	// api key, if any, for making authenticated requests.
	ApiKey string
	// a list of key+vals
	KeywordArgs []KeyVal

	// trait, see 'paged'
	Page    int
	PerPage int
	Order   string

	// trait, see 'subjected'
	SubjectList []string

	// trait, see 'container'
	ContainingList []string
}

type Response[T any] struct {
	// https://pkg.go.dev/net/http#Response
	HttpResponse http.Response

	// instantiated type.
	// todo: how to make this generic?
	Content T //interface{}

	// content type information derived from the response
	ContentType ContentType

	// the request was successfully authenticated
	Authenticated bool

	Deprecated bool
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

// makes a request to the api gateway for the given `endpoint` path,
// returning an api.Response containing an instance of the given type `T`.
func Request[T any](endpoint string, opts RequestConfig) Response[T] {
	url := "https://api.elifesciences.org" + endpoint

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	bytes, _ := io.ReadAll(resp.Body)

	t_inst := utils.FromJSON[T](bytes)

	return Response[T]{
		HttpResponse: *resp,
		Content:      t_inst,
		ContentType: ContentType{
			ContentType: "foo",
			Version:     1,
			Deprecated:  false,
		},
		Authenticated: false,
	}
}
