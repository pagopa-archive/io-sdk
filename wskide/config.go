package wskide

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
)

// IoSDKConfig is the global configuration type
type IoSDKConfig struct {
	// WhiskApiHost is the openwhisk api host
	WhiskAPIHost string `json:"whisk-apihost,omitempty"`
	// WhiskAPIKey is the openwhisk api key
	WhiskAPIKey string `json:"whisk-apikey,omitempty"`
	// WhiskNamespace is the openwhisk namespace
	WhiskNamespace string `json:"whisk-namespace,omitempty"`
	// IoAPIKey is the io api key
	IoAPIKey string `json:"io-apikey,omitempty"`
}

var (
	home string = os.Getenv("HOME")
	name string = ".iosdk"
	// Config is the global configuration
	Config *IoSDKConfig
)

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

func config() {

	var configFile = fmt.Sprintf("%s/%s", home, name)
	var ioSDKConfig IoSDKConfig
	var scanner *bufio.Scanner

	jsonFile, _ := os.Open(configFile)
	buf, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(buf, &ioSDKConfig)

	fmt.Printf("Enter whiskapihost: (%s) ", ioSDKConfig.WhiskAPIHost)
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	whiskapihost := scanner.Text()
	if whiskapihost == "" {
		whiskapihost = ioSDKConfig.WhiskAPIHost
	}

	fmt.Printf("Enter whiskapikey: (%s) ", ioSDKConfig.WhiskAPIKey)
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	whiskapikey := scanner.Text()
	if whiskapikey == "" {
		whiskapikey = ioSDKConfig.WhiskAPIKey
	}

	fmt.Printf("Enter whisknamespace: (%s) ", ioSDKConfig.WhiskNamespace)
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	whisknamespace := scanner.Text()
	if whisknamespace == "" {
		whisknamespace = ioSDKConfig.WhiskNamespace
	}

	fmt.Printf("Enter ioapikey: (%s) ", ioSDKConfig.IoAPIKey)
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ioapikey := scanner.Text()
	if ioapikey == "" {
		ioapikey = ioSDKConfig.IoAPIKey
	}

	res := &IoSDKConfig{
		WhiskAPIHost:   whiskapihost,
		WhiskAPIKey:    whiskapikey,
		WhiskNamespace: whisknamespace,
		IoAPIKey:       ioapikey}
	json, err := json.MarshalIndent(res, "", " ")

	// // add \n at the end of the json
	// var b uint8 = 10
	// json = append(json, b)

	// if err != nil {
	// 	return
	// }

	err = ioutil.WriteFile(configFile, json, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Wrote", configFile, "ok")
}
