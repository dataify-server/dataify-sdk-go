package dataify

import (
	"context"
	"errors"
	"fmt"
)

// DataifyError is implemented by errors produced by this SDK. Callers can
// classify SDK-originated failures with errors.As:
//
//	var sdkErr dataify.DataifyError
//	if errors.As(err, &sdkErr) {
//		// handle a Dataify SDK error
//	}
type DataifyError interface {
	error
	DataifyError() bool
}

// sdkError is a comparable sentinel error that implements DataifyError.
type sdkError string

func (e sdkError) Error() string      { return string(e) }
func (e sdkError) DataifyError() bool { return true }

func newSDKError(message string) error { return sdkError(message) }

// IsDataifyError reports whether err originated from this SDK. It recognizes
// API errors, validation errors, network errors, and the exported sentinel
// errors.
func IsDataifyError(err error) bool {
	var target DataifyError
	return errors.As(err, &target)
}

// ValidationError is returned before a request is sent when a request
// parameter is missing or invalid. Parameter is the JSON field name when one
// is available.
type ValidationError struct {
	Parameter string
	Message   string
}

func (e *ValidationError) Error() string {
	if e == nil {
		return "dataify: validation error"
	}
	if e.Parameter == "" {
		return "dataify: invalid request: " + e.Message
	}
	if e.Message == "" {
		return "dataify: invalid parameter " + e.Parameter
	}
	return fmt.Sprintf("dataify: invalid parameter %s: %s", e.Parameter, e.Message)
}

func (e *ValidationError) DataifyError() bool { return true }

func newValidationError(parameter, format string, args ...any) error {
	return &ValidationError{Parameter: parameter, Message: fmt.Sprintf(format, args...)}
}

// NetworkErrorKind classifies transport failures.
type NetworkErrorKind string

const (
	NetworkErrorKindNetwork  NetworkErrorKind = "network"
	NetworkErrorKindTimeout  NetworkErrorKind = "timeout"
	NetworkErrorKindCanceled NetworkErrorKind = "canceled"
)

var (
	// ErrNetwork is matched by transport failures that are not timeouts or
	// cancellations.
	ErrNetwork = newSDKError("dataify: network request failed")
	// ErrTimeout is matched when the request exceeded the configured timeout or
	// the caller's context deadline.
	ErrTimeout = newSDKError("dataify: request timed out")
	// ErrCanceled is matched when the caller canceled the request context.
	ErrCanceled = newSDKError("dataify: request canceled")
	// ErrInvalidResponse is matched when a successful HTTP response could not be
	// decoded as the expected JSON.
	ErrInvalidResponse = newSDKError("dataify: invalid API response")
)

// NetworkError wraps a transport-level failure while preserving the original
// error for errors.Unwrap, errors.Is, and errors.As.
type NetworkError struct {
	Kind NetworkErrorKind
	Err  error
}

func (e *NetworkError) Error() string {
	if e == nil {
		return "dataify: network error"
	}
	prefix := "dataify: network request failed"
	switch e.Kind {
	case NetworkErrorKindTimeout:
		prefix = "dataify: request timed out"
	case NetworkErrorKindCanceled:
		prefix = "dataify: request canceled"
	}
	if e.Err == nil {
		return prefix
	}
	return prefix + ": " + e.Err.Error()
}

func (e *NetworkError) Unwrap() error { return e.Err }

// Timeout reports whether the failure was a timeout. It makes NetworkError
// compatible with the common interface used by net.Error.
func (e *NetworkError) Timeout() bool {
	return e != nil && e.Kind == NetworkErrorKindTimeout
}

// Temporary is retained for net.Error compatibility.
func (e *NetworkError) Temporary() bool { return e.Timeout() }

func (e *NetworkError) DataifyError() bool { return true }

func (e *NetworkError) Is(target error) bool {
	if e == nil {
		return false
	}
	switch e.Kind {
	case NetworkErrorKindTimeout:
		return errors.Is(target, ErrTimeout)
	case NetworkErrorKindCanceled:
		return errors.Is(target, ErrCanceled)
	default:
		return errors.Is(target, ErrNetwork)
	}
}

func wrapNetworkError(err error) error {
	if err == nil {
		return nil
	}
	var existing *NetworkError
	if errors.As(err, &existing) {
		return err
	}
	kind := NetworkErrorKindNetwork
	switch {
	case errors.Is(err, context.Canceled):
		kind = NetworkErrorKindCanceled
	case errors.Is(err, context.DeadlineExceeded):
		kind = NetworkErrorKindTimeout
	default:
		if timeout, ok := err.(interface{ Timeout() bool }); ok && timeout.Timeout() {
			kind = NetworkErrorKindTimeout
		}
	}
	return &NetworkError{Kind: kind, Err: err}
}
