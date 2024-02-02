package futugrpc

import (
	"crypto/tls"
	"time"
)

type Options struct {
	keepAlive             time.Duration
	timeout               time.Duration
	maxConnectionAge      time.Duration
	maxConnectionAgeGrace time.Duration
	maxRetry              int
	addr                  string
	clientId              string
	tlsConfig             *tls.Config
}

type Option func(*Options)

func WithAddr(addr string) Option {
	return func(o *Options) {
		o.addr = addr
	}
}

func WithTlsConfig(tlsC tls.Config) Option {
	return func(o *Options) {
		o.tlsConfig = &tlsC
	}
}

func WithKeepAlive(keepAlive time.Duration) Option {
	return func(o *Options) {
		o.keepAlive = keepAlive
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.timeout = timeout
	}
}

func WithMaxConnectionAge(maxConnectionAge time.Duration) Option {
	return func(o *Options) {
		o.maxConnectionAge = maxConnectionAge
	}
}

func WithMaxConnectionAgeGrace(maxConnectionAgeGrace time.Duration) Option {
	return func(o *Options) {
		o.maxConnectionAgeGrace = maxConnectionAgeGrace
	}
}

func WithMaxRetry(maxRetry int) Option {
	return func(o *Options) {
		o.maxRetry = maxRetry
	}
}

func WithClientId(clientId string) Option {
	return func(o *Options) {
		o.clientId = clientId
	}
}
