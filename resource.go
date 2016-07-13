package rest

import (
	"net/http"
)

// Resource represents an interface information about a rest resource.
type Resource interface {
	Init()


	MainFunc() func(http.ResponseWriter, *http.Request)

	Get(http.ResponseWriter, *http.Request, map[string]string)

	Put(http.ResponseWriter, *http.Request, map[string]string)

	Post(http.ResponseWriter, *http.Request, map[string]string)

	Patch(http.ResponseWriter, *http.Request, map[string]string)

	Delete(http.ResponseWriter, *http.Request, map[string]string)

	Deinit()
}

