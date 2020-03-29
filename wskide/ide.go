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

	command := fmt.Sprintf(`docker run -d -p 3000:3000 --rm --name ide-js 
	--add-host=openwhisk:%s %s %s %s`, *openwhiskIP, mount, *extraArgIde, IdeJsImage)
	//OpenWhiskDockerWait()
	Sys(command)
	return nil
}

// OpenWhiskDockerWait wait for openwhisk to be
func OpenWhiskDockerWait() error {
	fmt.Println(Sys("docker exec openwhisk waitready"))
	return nil
}
