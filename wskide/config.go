package wskide

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

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

	// no error.. file should not be present
	jsonFile, _ := os.Open(configFile)
	buf, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(buf, &ioSDKConfig)

	// struct to interface
	v := reflect.ValueOf(ioSDKConfig)
	typeOfS := v.Type()

	// response map
	response := make(map[string]string)

	// parse interface, ask/read user input and assign value to response[]
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Enter %s: (%s) ", typeOfS.Field(i).Name, v.Field(i).Interface())
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		response[typeOfS.Field(i).Name] = scanner.Text()
		if response[typeOfS.Field(i).Name] == "" {
			response[typeOfS.Field(i).Name] = v.Field(i).Interface().(string)
		}
	}

	// the json
	res := &IoSDKConfig{
		WhiskAPIHost:   response["WhiskAPIHost"],
		WhiskAPIKey:    response["WhiskAPIKey"],
		WhiskNamespace: response["WhiskNamespace"],
		IoAPIKey:       response["IoAPIKey"]}

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
