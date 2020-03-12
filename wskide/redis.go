package wskide

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
	fmt.Println(Sys("docker stop redis"))
	return nil
}

// return empty string if ok, otherwise the error
func redisDockerRun() string {
	err := Run("docker pull " + RedisImage)
	if err != nil {
		return "cannot pull " + RedisImage
	}
	cmd := fmt.Sprintf(`docker run -d 
--rm --name redis --hostname redis %s`, RedisImage)
	_, err = SysErr(cmd)
	if err != nil {
		return "cannot start redis: " + err.Error()
	}
	return ""
}
