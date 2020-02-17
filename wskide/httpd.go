package wskide

import (
	"net/http"
	"strconv"
)

var (
	httpPortArg = httpCmd.Flag("port", "httpd port").Default("8080").Uint16()
	httpDirArg  = httpCmd.Flag("directory", "directory to serve").Default("src").ExistingDir()
)

// Httpd starts httpd server
func Httpd(port uint16, dir string) error {
	// http.Handle("/", http.FileServer(http.Dir("./src")))
	http.Handle("/", http.FileServer(http.Dir(dir)))
	// http.HandleFunc("/ping", ping)
	listen := "127.0.0.1:" + strconv.FormatInt(int64(port), 10)
	err := http.ListenAndServe(listen, nil)
	return err
}

// // Ping start ping
// func ping(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("pong"))
// }
