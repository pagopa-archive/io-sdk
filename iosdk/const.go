package main

import "runtime"

// ConfigFile is the config file, change number if you need the project to be re-initialized
const ConfigFile = "~/.iosdk.v3"

// Author is the main author
const Author = "Michele Sciabarra"

// Description is the descript
const Description = "iosdk is an SDK to develop connectors for the IO App, check http://github.com/pagopa/io-sdk for more infos"

//MinDockerVersion required
const MinDockerVersion = "18.06.3-ce"

// BrowserURL to access
const BrowserURL = "http://localhost:3280/"

// DockerHubUser is the Docker Hub User to use
const DockerHubUser = "pagopa"

// IdeImage is the image for the ide
const IdeImage = DockerHubUser + "/iosdk-theia"

// WhiskImage is the openwhisk image
const WhiskImage = DockerHubUser + "/iosdk-openwhisk"

// RedisImage is the image for redis
const RedisImage = "library/redis:5"

// RedisPassword is the local redis password for development
const RedisPassword = "password"

// IoAPIHost is the host to send messages
const IoAPIHost = "https://api.io.italia.it/api/v1"

// DefaultWhiskAPIKey is the default whisk api key
const DefaultWhiskAPIKey = "23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP"

// MinDockerMem is the minimum amount of memory required by docker
const MinDockerMem = (4 * 1000 * 1000 * 1000) - 1

// RuntimeOS to be changed for tests
var RuntimeOS = runtime.GOOS
