package main

import "encoding/base64"

// Main function for the action
func Main(args map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	if body, ok := args["__ow_body"].(string); ok {
		dec, err := base64.StdEncoding.DecodeString(body)
		if err == nil {
			args["__ow_body"] = string(dec)
		}
	}
	res["body"] = args
	return res
}
