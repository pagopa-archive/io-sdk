package wskide

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// ConfigState for spew
var sp spew.ConfigState = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
	SpewKeys:                true,
	DisableMethods:          true,
}

func show(args ...interface{}) {
	fmt.Println(args...)
}

func showf(format string, args ...interface{}) {
	fmt.Printf(format+"\\n", args...)
}

func dump(args ...interface{}) {
	sp.Dump(args...)
}

func fdump(filenames ...string) {
	for _, filename := range filenames {
		buf, err := ioutil.ReadFile(filename)
		if err == nil {
			show(string(buf))
		} else {
			show(err)
		}
	}
}

func grep(search string, data ...interface{}) {
	re := regexp.MustCompile(search)
	lines := strings.Split(sp.Sdump(data...), "\n")
	for _, line := range lines {
		//show(line)
		if re.Match([]byte(line)) {
			show(strings.TrimSpace(line))
		}
	}
}

type recovering func()

func capture(fn recovering) {
	defer func() {
		if r := recover(); r != nil {
			show("capture:", r)
		}
	}()
	fn()
}

func debug(arg ...interface{}) {
	log.Debug(arg...)
}

func trace(arg ...interface{}) {
	log.Trace(arg...)
}

func traceOn() {
	log.SetLevel(log.TraceLevel)
}

func traceOff() {
	log.SetLevel(log.DebugLevel)
}

func TestMain(m *testing.M) {
	flag.Parse()
	*DryRunFlag = true
	log.SetOutput(os.Stderr)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	})
	log.SetLevel(log.DebugLevel)
	//log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}
