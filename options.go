package chttp

import "time"

type Option func(o *options)

type options struct {
	timeout time.Duration
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.timeout = timeout
	}
}
