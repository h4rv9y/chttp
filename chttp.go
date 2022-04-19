package chttp

import (
	"net/http"
	"strings"
)

type client struct {
	httpClient http.Client
}

func NewClient(opts ...Option) *client {
	var (
		c client
		o options
	)
	for _, opt := range opts {
		opt(&o)
	}
	c.httpClient.Timeout = o.timeout

	return &c
}

var httpScheme = []byte("http")

func hasProtocolScheme(url string) bool {
	if len(url) < 8 {
		return false
	}
	prefix := url[:8]
	for i := 0; i < 4; i++ {
		if prefix[i] != httpScheme[i] {
			return false
		}
	}
	if url[4] == ':' && url[5] == '/' && url[6] == '/' {
		return true
	}
	if url[4] == 's' && url[5] == ':' && url[6] == '/' && url[7] == '/' {
		return true
	}
	return false
}

func (c *client) Get(url string) (r *request) {
	if !hasProtocolScheme(url) {
		url = strings.Join([]string{"http://", url}, "")
	}
	return &request{
		c:      c,
		url:    url,
		method: "GET",
		meta:   make(map[string][]string),
		params: make(map[string][]string),
	}
}

func (c *client) Post(url string) (r *request) {
	if !hasProtocolScheme(url) {
		url = strings.Join([]string{"http://", url}, "")
	}
	return &request{
		c:      c,
		url:    url,
		method: "POST",
		meta:   make(map[string][]string),
		params: make(map[string][]string),
	}
}

func (c *client) Put(url string) (r *request) {
	if !hasProtocolScheme(url) {
		url = strings.Join([]string{"http://", url}, "")
	}
	return &request{
		c:      c,
		url:    url,
		method: "PUT",
	}
}

func (c *client) Delete(url string) (r *request) {
	if !hasProtocolScheme(url) {
		url = strings.Join([]string{"http://", url}, "")
	}
	return &request{
		c:      c,
		url:    url,
		method: "DELETE",
	}
}

var DefaultClient = NewClient()

func Get(url string) (r *request) {
	if !hasProtocolScheme(url) {
		url = strings.Join([]string{"http://", url}, "")
	}
	return &request{
		c:      DefaultClient,
		url:    url,
		method: "GET",
	}
}

func Post(url string) (r *request) {
	if !hasProtocolScheme(url) {
		url = strings.Join([]string{"http://", url}, "")
	}
	return &request{
		c:      DefaultClient,
		url:    url,
		method: "POST",
	}
}

func Put(url string) (r *request) {
	if !hasProtocolScheme(url) {
		url = strings.Join([]string{"http://", url}, "")
	}
	return &request{
		c:      DefaultClient,
		url:    url,
		method: "PUT",
	}
}

func Delete(url string) (r *request) {
	if !hasProtocolScheme(url) {
		url = strings.Join([]string{"http://", url}, "")
	}
	return &request{
		c:      DefaultClient,
		url:    url,
		method: "DELETE",
	}
}
