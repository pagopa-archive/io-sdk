package wskide

import (
	"fmt"
	"os"
)

//MinDockerVersion required
const MinDockerVersion = "18.06.3-ce"

// BrowserURL to access
const BrowserURL = "http://localhost:3280/"

// IdeJsImage is the image for the ide
const IdeJsImage = "actionloop/ide-js:config"

// OpenwhiskStandaloneImage is the image for the standalone openwhisk
const OpenwhiskStandaloneImage = "actionloop/iosdk:config"

// APIHOST to send messages
const APIHOST = "https://api.cd.italia.it/api/v1"

// RedisImage is the image for redis
const RedisImage = "library/redis:5"

// IOAPIHOST to send messages
const IOAPIHOST = "https://api.cd.italia.it/api/v1"

const name string = ".iosdk"

var home string = os.Getenv("HOME")

var configFile = fmt.Sprintf("%s/%s", home, name)
