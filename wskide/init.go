package wskide

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/src-d/go-git.v4"
)

var (
	initLangFlag = initCmd.Flag("language", "SDK language").Default("javascript").String()
	initDirFlag  = initCmd.Flag("directory", "SDK dir").Default("project").String()
)

// Init io-sdk
func Init(dir, lang string) error {

	repo := fmt.Sprintf("https://github.com/pagopa/io-sdk-%s", lang)
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      repo,
		Progress: ioutil.Discard,
	})

	return err

}
