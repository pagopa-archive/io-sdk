package wskide

import (
	"fmt"
	"os"
)

// Author is the main author
const Author = "Michele Sciabarra"

// Description is the descript
const Description = "io-sdk is an SDK to develop IO App connectorsm, check http://github.com/pagopa/io-sdk "

//MinDockerVersion required
const MinDockerVersion = "18.06.3-ce"

// BrowserURL to access
const BrowserURL = "http://localhost:3280/"

// IdeJsImage is the image for the ide
const IdeJsImage = "actionloop/ide-js:latest"

// OpenwhiskStandaloneImage is the image for the standalone openwhisk
const OpenwhiskStandaloneImage = "actionloop/iosdk:latest"

// APIHOST to send messages
const APIHOST = "https://api.cd.italia.it/api/v1"

// RedisImage is the image for redis
const RedisImage = "library/redis:5"

// IOAPIHOST to send messages
const IOAPIHOST = "https://api.cd.italia.it/api/v1"

const name string = ".iosdk"

var home string = os.Getenv("HOME")

var configFile = fmt.Sprintf("%s/%s", home, name)
