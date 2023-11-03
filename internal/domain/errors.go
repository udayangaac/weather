package domain

import "errors"

var (

	// ErrNotModified is an error variable indicating that something is not modified.
	ErrNotModified = errors.New("not modified")

	// ErrLocationNotFound is an error indicating that a location was not found.
	ErrLocationNotFound = errors.New("location not found")

	// ErrServiceBusy is an error indicating that internal services are not responding.
	ErrServiceBusy = errors.New("service busy")
)
