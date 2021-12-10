package foohandler

import (
	"fmt"
	"html"
	"net/http"
)

func (f FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "FOO, %q", html.EscapeString(r.URL.Path))
}
