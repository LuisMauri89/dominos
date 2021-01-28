package advent

import (
	"fmt"
	"net/http"
)

type AuthError struct {
	Realm string
}

func (AuthError) StatusCode() int {
	return http.StatusUnauthorized
}

func (AuthError) Error() string {
	return http.StatusText(http.StatusUnauthorized)
}

func (e AuthError) Headers() http.Header {
	return http.Header{
		"Content-Type":           []string{"text/plain; charset=utf-8"},
		"X-Content-Type-Options": []string{"nosniff"},
		"WWW-Authenticate":       []string{fmt.Sprintf(`Basic realm=%q`, e.Realm)},
	}
}
