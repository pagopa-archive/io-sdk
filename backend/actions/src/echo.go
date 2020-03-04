package main

// Main function for the action
func Main(obj map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	res["body"] = obj
	return res
}
