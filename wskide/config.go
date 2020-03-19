package wskide

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

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
}
