package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/alecthomas/units"
	"github.com/coreos/go-semver/semver"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
)

func preflightEnsureDockerVersion() error {
	if *skipDockerVersion {
		return nil
	}
	version, err := dockerVersion()
	if err != nil {
		return err
	}
	vA := semver.New(MinDockerVersion)
	vB := semver.New(strings.TrimSpace(version))
	if vB.Compare(*vA) == -1 {
		return fmt.Errorf("Installed docker version %s is no longer supported", vB)
	}
	return nil
}

func preflightInHomePath(dir string) error {
	// do not check if the directory is empty
	if dir == "" {
		return nil
	}
	homePath, err := homedir.Dir()
	if err != nil {
		return err
	}
	dir, err = filepath.Abs(dir)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(dir, homePath) {
		return fmt.Errorf("work directory %s should be below your home directory %s;\nthis is required to be accessible by Docker", dir, homePath)
	}
	return nil
}

// Preflight perform preflight checks
func Preflight(dir string) error {
	err := preflightDockerMemory()
	if err != nil {
		return err
	}
	err = preflightEnsureDockerVersion()
	if err != nil {
		return err
	}
	err = preflightInHomePath(dir)
	if err != nil {
		return err
	}
	return nil
}

// preflightDockerMemory checks docker memory
func preflightDockerMemory() error {
	out, err := SysErr("@docker info")
	if err != nil {
		return fmt.Errorf("Docker is not running")
	}
	var search = regexp.MustCompile(`Total Memory: (.*)`)
	result := search.FindString(string(out))
	if result == "" {
		return fmt.Errorf("Docker is not running")
	}
	mem := strings.Split(result, ":")
	memory := strings.TrimSpace(mem[1])
	n, err := units.ParseStrictBytes(memory)
	log.Debug("mem:", n)
	//fmt.Println(n)
	if n <= int64(MinDockerMem) {
		return fmt.Errorf("IOSDK needs 4GB memory allocatable on docker")
	}

	return nil

}
