package main

import (
	"fmt"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

// IdeDeploy deploys and mounts a folder
func IdeDeploy(dir string, info string) error {
	fmt.Println("Deploying IDE...")
	if dir != "" {
		err := preflightInHomePath(dir)
		if err != nil {
			return err
		}
	}
	return ideDockerRun(dir, info)
}

// IdeDestroy destroys ide
func IdeDestroy() error {
	fmt.Println("Destroying IDE...")
	fmt.Println(Sys("docker kill iosdk-theia"))
	return nil
}

// ideDockerRun starts the ide
// it also mounts the project folder if the directory is not empty
func ideDockerRun(dir string, info string) (err error) {
	image := IdeImage + ":" + Version
	if err = dockerPull(image); err != nil {
		return err
	}
	mount := ""
	if dir != "" {
		dir, err = filepath.Abs(dir)
		LogIf(err)
		if err == nil {
			dir = fixPathDockerToolbox(dir, info)
			mount = fmt.Sprintf("-v %s:/home/project", dir)
		}
	}

	openwhiskIP := dockerIP("iosdk-openwhisk")
	if openwhiskIP == nil {
		return fmt.Errorf("cannot find openwhisk")
	}

	uid := ""
	if RuntimeOS == "linux" {
		usr, err := user.Current()
		LogIf(err)
		if err == nil {
			uid = fmt.Sprintf("-u %s", usr.Uid)
		}
	}

	command := fmt.Sprintf(`docker run -d -p 3000:3000
	--rm --name iosdk-theia -e HOME=/home/project 
	%s %s --add-host=openwhisk:%s %s`, mount, uid, *openwhiskIP, image)
	Sys(command)
	return nil
}

// OpenWhiskDockerWait wait for openwhisk to be
func OpenWhiskDockerWait() error {
	fmt.Println(Sys("docker exec iosdk-whisk waitready"))
	return nil
}

// on windows if you are using Docker Toolbox you need to change the path format
// it expects an absolute path starting with `<drive>:` in windows
func fixPathDockerToolbox(dir string, info string) string {
	if RuntimeOS != "windows" {
		return dir
	}
	var search = regexp.MustCompile(`Operating System: Boot2Docker`)
	if search.FindString(info) == "" {
		return dir
	}
	dir = strings.ReplaceAll(dir, "\\", "/")
	if dir[1] == ':' {
		dir = "//" + string(dir[0]) + dir[2:]
	}
	return dir
}
