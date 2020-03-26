package wskide

import (
	"fmt"
	"path/filepath"
)

// IdeDeploy deploys and mounts a folder
func IdeDeploy(dir string) error {
	fmt.Println("Deploying IDE...")
	if dir != "" {
		err := preflightInHomePath(dir)
		if err != nil {
			return err
		}
	}
	return ideDockerRun(dir)
}

// IdeDestroy destroys ide
func IdeDestroy() error {
	fmt.Println("Destroying IDE...")
	fmt.Println(Sys("docker kill ide-js"))
	return nil
}

// ideDockerRun starts the ide
// it also mounts the project folder if the directory is not empty
func ideDockerRun(dir string) (err error) {
	Config, _ := LoadConfig()

	err = Run("docker pull " + IdeJsImage)
	if err != nil {
		return err
	}

	mount := ""
	if dir != "" {
		dir, err = filepath.Abs(dir)
		LogIf(err)
		if err == nil {
			mount = fmt.Sprintf("-v %s:/home/project", dir)
		}
	}

	openwhiskIP := dockerIP("openwhisk")
	if openwhiskIP == nil {
		return fmt.Errorf("cannot find openwhisk")
	}

	command := fmt.Sprintf(`docker run -d -p 3000:3000 --rm --name ide-js  -e CONFIG_FORCE_whisk_users_guest=%s
	--add-host=openwhisk:%s %s %s`, Config.WhiskAPIKey, *openwhiskIP, mount, IdeJsImage)
	//OpenWhiskDockerWait()
	Sys(command)

	err = Run("docker exec ide-js wsk property set apihost http://openwhisk:3280 --apihost http://openwhisk:3280 auth " + Config.WhiskAPIKey + " --auth " + Config.WhiskAPIKey)
	if err != nil {
		fmt.Println("cannot update properties: " + err.Error())
	}

	return nil
}

// OpenWhiskDockerWait wait for openwhisk to be
func OpenWhiskDockerWait() error {
	fmt.Println(Sys("docker exec openwhisk waitready"))
	return nil
}
