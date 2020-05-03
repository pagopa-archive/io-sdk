package main

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
	fi, err := os.Stat(dir)
	if err == nil {
		switch mode := fi.Mode(); {
		case mode.IsDir():
			// do directory stuff
			return dir, nil
		case mode.IsRegular():
			panic("existing path that is not a directory")
		}
	}
	// error, file does not exist

	if repo == "" {
		fmt.Println("Select one of the available templates for importers, or provide your own.")
		fmt.Println("The javascript template is for Excel import.")
		fmt.Println("The java template is for SQL import.")
		fmt.Println("The python template is for REST import.")
		fmt.Println("The github template requires a github repo (user/path).")

		tmpls := map[string]string{
			"js":   "javascript",
			"java": "java",
			"py":   "python",
			"gh":   "github",
		}

		repo, err = SelectTemplate(tmpls)

		if err != nil {
			fmt.Println(err)
		}
	}

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
