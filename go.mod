module github.com/pagopa/io-sdk

go 1.13

require (
	github.com/markbates/pkger v0.14.0
	github.com/pagopa/io-sdk/wskide v0.0.0-20200211113841-2dacb9565e81
	github.com/sergi/go-diff v1.1.0 // indirect
	golang.org/x/crypto v0.0.0-20200210222208-86ce3cb69678 // indirect
)

replace github.com/pagopa/io-sdk/wskide => ./wskide
