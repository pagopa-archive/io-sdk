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
	fmt.Println("Destroying Whisk...")
	Sys("docker exec iosdk-openwhisk stop")
	fmt.Println()
	return nil
}

// return empty string if ok, otherwise the error
func whiskDockerRun() string {
	image := WhiskImage + ":" + Version
	if err := dockerPull(image); err != nil {
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
-e CONTAINER_EXTRA_ENV=__OW_REDIS=%s
-e CONFIG_FORCE_whisk_users_guest=%s
-v %s %s`,
		*redisIP, Config.WhiskAPIKey, dockerSock, image)
	_, err := SysErr(cmd)
	if err != nil {
		return "cannot start server: " + err.Error()
	}
	err = Run("docker exec iosdk-openwhisk waitready")
	if err != nil {
		return "server readyness error: " + err.Error()
	}
	return ""
}
