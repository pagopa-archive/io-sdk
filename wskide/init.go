package wskide

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

var (
	initLangFlag = initCmd.Flag("language", "SDK language").Default("javascript").String()
	initDirFlag  = initCmd.Flag("directory", "SDK dir").Default("project").String()
)

// Init openwhisk-ide
func Init() error {

	repo := fmt.Sprintf("https://github.com/pagopa/io-sdk-%s", *initLangFlag)
	_, err := git.PlainClone(*initDirFlag, false, &git.CloneOptions{
		URL:      repo,
		Progress: os.Stdout,
	})

	return err

}
