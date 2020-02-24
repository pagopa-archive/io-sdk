package wskide

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/protocol/packp/sideband"
)

// Init io-sdk
func Init(dir, lang string, log sideband.Progress) error {
	err := preflightInHomePath(dir)
	if err != nil {
		return err
	}
	fmt.Printf("Preparing work directory %s for %s\n", dir, lang)
	repo := fmt.Sprintf("https://github.com/pagopa/io-sdk-%s", lang)
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL:      repo,
		Progress: log,
	})
	if err != nil {
		return err
	}
	fmt.Println("Done.")
	return err
}
