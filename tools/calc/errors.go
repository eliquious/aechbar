package main

type StatusCode int

// Success codes
const (
	OK StatusCode = iota + 2000
)

// Error codes
const (
	InternalServerError StatusCode = iota + 5000
)

var statusCodes = map[StatusCode]string{

	// Success
	OK: "OK",

	// General errors
	InternalServerError: "InternalServerError",
}
