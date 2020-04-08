package wskide

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/protocol/packp/sideband"
)

// Init io-sdk
func Init(dir, repo string, log sideband.Progress) (string, error) {

	appDir := "importer"
	if ConfigLoad() == nil {
		appDir = Config.AppDir
	}

	var err error
	if dir == "" {
		dir = Input("Work Directory (can already exists)", appDir)
	}
	if dir == "" {
		return "", errors.New("Directory not specified")
	}
	dir, err = filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	err = preflightInHomePath(dir)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		fmt.Println("Using existing work directory.")
		return dir, nil
	}

	if repo == "" {
		fmt.Println("Select one of the available templates for importers, or provide your own.")
		fmt.Println("The javascript template is for Excel import.")
		fmt.Println("The java template is for SQL import.")
		fmt.Println("The python template is for REST import.")
		fmt.Println("The github template requires a github repo (user/path).")
		opt := Select("Which template:", "javascript,java,python,github")
		if opt == "" {
			return "", fmt.Errorf("aborted template selection")
		}
		if opt == "github" {
			repo = Input("GitHub user/path", "")
		} else {
			repo = fmt.Sprintf("pagopa/io-sdk-%s", opt)
		}
	}
	repo = fmt.Sprintf("https://github.com/%s", repo)

	fmt.Printf("Preparing work directory %s for %s\n", dir, repo)
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL:      repo,
		Progress: log,
	})
	if err != nil {
		return "", err
	}
	fmt.Println("Done.")
	return dir, err
}
