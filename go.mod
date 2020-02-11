module github.com/pagopa/io-sdk

go 1.12

replace github.com/pagopa/io-sdk/wskide => ./wskide

require (
	github.com/pagopa/io-sdk/wskide v0.0.0-00010101000000-000000000000
	github.com/sciabarracom/openwhisk-ide/wskide v0.0.0-20200211103852-c06dd04df0ed
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)
