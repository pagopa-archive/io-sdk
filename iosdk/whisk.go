package main

import (
	"fmt"
	"runtime"
)

// WhiskDeploy deploys openwhisk standalone
func WhiskDeploy() error {
	fmt.Println("Deploying Whisk...")
	whiskDockerRun()
	return nil
}

// WhiskDestroy destroys openwhisk standalone
func WhiskDestroy() error {
	Sys("docker exec iosdk-openwhisk stop")
	fmt.Println("Destroying Whisk: iosdk-openwhisk")
	return nil
}

// pulling all the runtime images before starting
// return the whisk image to start
func pullImages() (string, error) {
	image := WhiskImage + ":" + Version
	if err := dockerPull(image); err != nil {
		return "", err
	}
	if err := dockerPull(PythonImage); err != nil {
		return "", err
	}
	if err := dockerPull(NodeJSImage); err != nil {
		return "", err
	}
	return image, nil 
}

// return empty string if ok, otherwise the error
func whiskDockerRun() string {
	image, err := pullImages() 
        if err != nil {
		return err.Error()
	}
	redisIP := dockerIP("iosdk-redis")
	if redisIP == nil {
		return "cannot locate redis"
	}
	dockerSock := "/var/run/docker.sock:/var/run/docker.sock"
	if runtime.GOOS == "windows" {
		dockerSock = "/" + dockerSock
	}
	cmd := fmt.Sprintf(`docker run -d -p 3280:3280
--rm --name iosdk-openwhisk --hostname openwhisk
-e CONTAINER_EXTRA_ENV=__OW_REDIS_IP=%s -e CONTAINER_EXTRA_ENV1=__OW_REDIS_PASSWORD=%s
-e CONFIG_FORCE_whisk_users_guest=%s
-v %s %s`,
		*redisIP, RedisPassword, Config.WhiskAPIKey, dockerSock, image)
	_, err = SysErr(cmd)
	if err != nil {
		return "cannot start server: " + err.Error()
	}
	err = Run("docker exec iosdk-openwhisk waitready")
	if err != nil {
		return "server readyness error: " + err.Error()
	}
	return ""
}
