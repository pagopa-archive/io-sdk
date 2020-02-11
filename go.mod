module github.com/sciabarracom/openwhisk-ide

go 1.12

replace github.com/sciabarracom/openwhisk-ide/wskide => ./wskide

require (
	github.com/sciabarracom/openwhisk-ide/wskide v0.0.0-00010101000000-000000000000
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)
