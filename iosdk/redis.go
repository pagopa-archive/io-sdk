package main

import (
	"fmt"
)

// RedisDeploy deploys openwhisk standalone
func RedisDeploy() error {
	fmt.Println("Deploying Redis...")
	fmt.Println(redisDockerRun())
	return nil
}

// RedisDestroy destroys openwhisk standalone
func RedisDestroy() error {
	fmt.Println("Destroying Redis...")
	fmt.Println(Sys("docker stop iosdk-redis"))
	return nil
}

// return empty string if ok, otherwise the error
func redisDockerRun() string {
	if err := dockerPull(RedisImage); err != nil {
		return err.Error()
	}
	cmd := fmt.Sprintf(`docker run -d -p 6379:6379
--rm --name iosdk-redis --hostname redis %s`, RedisImage)
	_, err := SysErr(cmd)
	if err != nil {
		return "cannot start redis: " + err.Error()
	}
	return ""
}
