package wskide

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
)

//MinDockerVersion required
const MinDockerVersion = "18.06.3-ce"

// BrowserURL to access
const BrowserURL = "http://localhost:3280/"

// IdeJsImage is the image for the ide
const IdeJsImage = "actionloop/ide-js:ide-plugin"

// OpenwhiskStandaloneImage is the image for the standalone openwhisk
const OpenwhiskStandaloneImage = "actionloop/iosdk:latest"

// RedisImage is the image for redis
const RedisImage = "library/redis:5"

// IOAPIHOST to send messages
const IOAPIHOST = "https://api.cd.italia.it/api/v1"

// IoSDKConfig is the global configuration type
type IoSDKConfig struct {
	// WhiskApiHost is the openwhisk api host
	WhiskAPIHost string `json:"whisk-apihost"`
	// WhiskAPIKey is the openwhisk api key
	WhiskAPIKey string `json:"whisk-apikey"`
	// WhiskNamespace is the openwhisk namespace
	WhiskNamespace string `json:"whisk-namespace"`
	// IoAPIKey is the io api key
	IoAPIKey string `json:"io-apikey"`
}

// Config is the global configuration
var Config *IoSDKConfig

// LoadConfig load the configuration
func LoadConfig() error {
	configFile, err := homedir.Expand("~/.iosdk")
	if err != nil {
		return err
	}
	if _, err := os.Stat(configFile); err != nil {
		return err
	}
	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	json.Unmarshal(buf, &Config)
	return nil
}
