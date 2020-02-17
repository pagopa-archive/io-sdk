package wskide

import (
	"net/http"
	"strconv"

	"github.com/markbates/pkger"
)

var (
	httpPortArg = httpCmd.Flag("port", "httpd port").Default("8080").Uint16()
)

// Httpd starts httpd server
func Httpd(port uint16) error {
	listen := "127.0.0.1:" + strconv.FormatInt(int64(port), 10)
	dir := http.FileServer(pkger.Dir("github.com/pagopa/io-sdk:/backend/public/"))
	err := http.ListenAndServe(listen, dir)
	return err
}
