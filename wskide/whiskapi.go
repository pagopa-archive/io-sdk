package wskide

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func whiskURL(operation string) string {
	if operation[:1] == "/" {
		return fmt.Sprintf("%s/api/v1/namespaces%s",
			Config.WhiskAPIHost,
			operation)
	}
	return fmt.Sprintf("%s/api/v1/namespaces/%s/%s",
		Config.WhiskAPIHost,
		Config.WhiskNamespace,
		operation)
}

func whiskAuth() (string, string) {
	up := strings.Split(Config.WhiskAPIKey, ":")
	return up[0], up[1]
}

func whiskPost(action string,
	args map[string]interface{}) (*http.Request, error) {
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", whiskURL(action),
		bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(whiskAuth())
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func whiskPut(action string,
	args p) (*http.Request, error) {
	data, err := json.Marshal(args)

	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", whiskURL(action),
		bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(whiskAuth())
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func whiskCall(req *http.Request) map[string]interface{} {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return mkErr(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return mkErr(err)
	}
	// encode answer
	var objmap map[string]interface{}
	err = json.Unmarshal(body, &objmap)
	if err != nil {
		return mkErr(err)
	}
	return objmap
}

func whiskInvoke(action string, args map[string]interface{},
	blocking bool, result bool) map[string]interface{} {
	invoke := fmt.Sprintf("actions/%s?blocking=%t&result=%t",
		action, blocking, result)
	req, err := whiskPost(invoke, args)
	if err != nil {
		return mkErr(err)
	}
	return whiskCall(req)
}

func whiskPackageUpdate(action string, args p) map[string]interface{} {
	invoke := fmt.Sprintf("packages/%s?overwrite=true",
		action)
	// action, blocking, result)
	req, err := whiskPut(invoke, args)
	if err != nil {
		return mkErr(err)
	}
	return whiskCall(req)
}

type kv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type p struct {
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	Parameters []kv   `json:"parameters"`
}

func prova() {
	jsonConfig := p{
		Name:      "iosdk",
		Namespace: "guest",
		Parameters: []kv{
			kv{Key: "a", Value: "1"},
			kv{Key: "apikey", Value: Config.WhiskAPIKey},
		},
	}
	fmt.Println(whiskPackageUpdate("iosdk", jsonConfig))
}
