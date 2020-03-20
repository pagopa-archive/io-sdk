package wskide

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
)

//MinDockerVersion required
const MinDockerVersion = "18.06.3-ce"

// config file json
type jsonConfig struct {
	Apikey string `json:"apikey"`
}

var (
	home string = os.Getenv("HOME")
	name string = ".iosdk"
)

func config() {
	fmt.Print("Enter apikey: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	res := &jsonConfig{
		Apikey: text}
	json, err := json.MarshalIndent(res, "", " ")

	// add \n at the end of the json
	var b uint8 = 10
	json = append(json, b)

	if err != nil {
		return
	}
	configFile := fmt.Sprintf("%s/%s", home, name)
	err = ioutil.WriteFile(configFile, json, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Wrote apikey on", configFile, "ok")
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
