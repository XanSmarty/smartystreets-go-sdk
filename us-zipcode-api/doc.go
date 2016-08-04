package us_zipcode

import "net/http"

type requestSender interface {
	Send(*http.Request) ([]byte, error)
}
