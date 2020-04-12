package main

import "github.com/pagopa/io-sdk/wskide"

// Version is the current version - it will be set when built
var Version = "master"

//main
func main() {
	wskide.Main(Version)
}
