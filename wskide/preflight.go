package wskide

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/mitchellh/go-homedir"
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
	err := preflightEnsureDockerVersion()
	if err != nil {
		return err
	}
	err = preflightInHomePath(dir)
	if err != nil {
		return err
	}
	return nil
}
