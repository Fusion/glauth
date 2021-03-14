package handler

import (
	"context"

	"github.com/GeertJohan/yubigo"
	"github.com/glauth/glauth/pkg/config"
	"github.com/go-logr/logr"
)

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	Logger   logr.Logger
	Config   *config.Config
	Context  *context.Context
	YubiAuth *yubigo.YubiAuth
}

// newOptions initializes the available default options.
func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// newOptions initializes the available default options.
func NewOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Logger provides a function to set the logger option.
func Logger(val logr.Logger) Option {
	return func(o *Options) {
		o.Logger = val
	}
}

// Config provides a function to set the config option.
func Config(val *config.Config) Option {
	return func(o *Options) {
		o.Config = val
	}
}

// Context provides a function to set the context option.
func Context(val *context.Context) Option {
	return func(o *Options) {
		o.Context = val
	}
}

// YubiAuth provides a function to set the yubiauth option.
func YubiAuth(val *yubigo.YubiAuth) Option {
	return func(o *Options) {
		o.YubiAuth = val
	}
}
