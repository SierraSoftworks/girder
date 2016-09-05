package errors

import "runtime/debug"

type StacktraceProvider interface {
	Get() interface{}
}

var Stacktrace StacktraceProvider = &DefaultStacktraceProvider{}

type DefaultStacktraceProvider struct {
}

func (p *DefaultStacktraceProvider) Get() interface{} {
	return string(debug.Stack())
}
