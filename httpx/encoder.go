package httpx

import (
	"encoding/json"
	"net/http"
)

// Encoder describes an object capable of encoding a response.
type Encoder interface {
	// Encode writes a serialization of v to w, optionally using additional
	// information from the http.Request to do so.
	Encode(w http.ResponseWriter, r *http.Request, v interface{}) error

	// ContentType gets a string that will become the Content-Type header
	// when responding through w to the specified http.Request.
	// Most of the time the argument will be ignored, but occasionally
	// details in the request, or even in the headers in the ResponseWriter may
	// change the content type.
	ContentType(w http.ResponseWriter, r *http.Request) string
}

// JSONEncoder is an Encoder for JSON.
var JSONEncoder Encoder = (*jsonEncoder)(nil)

type jsonEncoder struct{}

func (*jsonEncoder) Encode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func (*jsonEncoder) ContentType(w http.ResponseWriter, r *http.Request) string {
	return "application/json; charset=utf-8"
}
