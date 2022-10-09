package api

// internal communication between components uses a common function interface of:
//  `(endpoint, options)` or `(endpoint, param1, param2, paramN, options)`
// returning a *http-like* response.
// the http-like response uses a status code, a type instance of the body content if
// content type was successfully negotiated, etc.

import (
	"net/http"
	"net/url"
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

type Request struct {
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
	http.Response

	// instantiated type.
	// todo: how to make this generic?
	Content T //interface{}

	// content type information derived from the response
	ContentType ContentType

	// the request was successfully authenticated
	Authenticated bool
}

func request[T any](endpoint string, opts Request) Response[T] {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "api.elifesciences.org"
	// q := u.Query()
	// for keyval in opts.KeywordArgs {
	// 	q.set(keyval.Key, keyval.Val)
	// }

	return Response[T]{}
}
