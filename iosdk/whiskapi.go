package main

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
			Config.WhiskAPIHostLocal,
			operation)
	}
	return fmt.Sprintf("%s/api/v1/namespaces/%s/%s",
		Config.WhiskAPIHostLocal,
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
	data []byte) (*http.Request, error) {
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

func whiskPackageUpdate(pkg string, data []byte) map[string]interface{} {
	invoke := fmt.Sprintf("packages/%s?overwrite=true",
		pkg)
	req, err := whiskPut(invoke, data)
	if err != nil {
		return mkErr(err)
	}
	return whiskCall(req)
}

type whiskKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type whiskWrap struct {
	Name       string          `json:"name"`
	Namespace  string          `json:"namespace"`
	Parameters []whiskKeyValue `json:"parameters"`
}

func whiskConfigKeyValues(m map[string]string) []whiskKeyValue {
	res := make([]whiskKeyValue, 0)
	for k, v := range m {
		kw := whiskKeyValue{
			Key:   k,
			Value: v,
		}
		res = append(res, kw)
	}
	return res
}

// WhiskUpdatePackageParameters update a package with a given map of parameters
func WhiskUpdatePackageParameters(pkg string, m map[string]string) map[string]interface{} {
	wrap := whiskWrap{
		Name:       pkg,
		Namespace:  Config.WhiskNamespace,
		Parameters: whiskConfigKeyValues(m),
	}

	buf, err := json.Marshal(wrap)
	if err != nil {
		return mkErr(err)
	}
	return whiskPackageUpdate(pkg, buf)
}
