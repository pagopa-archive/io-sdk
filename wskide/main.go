package wskide

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/browser"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

// VerboseFlag is flag for verbose

var (
	// TestModeFlag enable behaviours useful for testing

	// global flags
	debugFlag = kingpin.Flag("debug", "Enable Debug log").Hidden().Default("false").Bool()
	traceFlag = kingpin.Flag("trace", "Enable Trace log").Hidden().Default("false").Bool()
	testFlag  = kingpin.Flag("test", "Enable Test Mode").Hidden().Default("false").Bool()

	// hidden global flags
	skipDockerVersion = kingpin.Flag("skip-docker-version", "Skip check of docker version").Hidden().Default("false").Bool()

	// hidden debug commands
	debugCmd        = kingpin.Command("debug", "debug").Hidden()
	ideDeployCmd    = debugCmd.Command("ide-deploy", "Create IDE deployment").Hidden()
	ideDestroyCmd   = debugCmd.Command("ide-destroy", "Destroy IDE deployment").Hidden()
	whiskDeployCmd  = debugCmd.Command("whisk-deploy", "Create Whisk deployment").Hidden()
	whiskDestroyCmd = debugCmd.Command("whisk-destroy", "Destroy Whisk deployment").Hidden()
	redisDeployCmd  = debugCmd.Command("redis-deploy", "Create Redis deployment").Hidden()
	redisDestroyCmd = debugCmd.Command("redis-destroy", "Destroy Redis deployment").Hidden()
	inputCmd        = debugCmd.Command("input", "Input test").Hidden()
	inputArgCmd     = inputCmd.Arg("input arg", "input arg").Default("").String()
	inputSelectFlag = inputCmd.Flag("select", "select").Bool()
	// start, stop, init and status
	startCmd    = kingpin.Command("start", "Start Development Enviroment")
	startDirArg = startCmd.Arg("dir", "Project dir").Required().String()
	// init
	initCmd          = kingpin.Command("init", "Initialise SDK Repository")
	initDirArg       = initCmd.Arg("directory", "work directory").Required().String()
	initRepoArg      = initCmd.Arg("repo", "Repository").Default("").String()
	initWhiskKeyFlag = initCmd.Flag("whisk-apikey", "Whisk API Key").Default("").String()
	initIOKeyFlag    = initCmd.Flag("io-apikey", "IO API Key").Default("").String()

	// stop
	stopCmd = kingpin.Command("stop", "Stop Development Environment")
	// status
	statusCmd = kingpin.Command("status", "Check Containers Status")
)

func parseDebug(cmd string) bool {
	switch cmd {
	case ideDeployCmd.FullCommand():
		IdeDeploy("")
	case ideDestroyCmd.FullCommand():
		IdeDestroy()
	case whiskDeployCmd.FullCommand():
		WhiskDeploy()
	case whiskDestroyCmd.FullCommand():
		WhiskDestroy()
	case redisDeployCmd.FullCommand():
		RedisDeploy()
	case redisDestroyCmd.FullCommand():
		RedisDestroy()
	case inputCmd.FullCommand():
		if !*inputSelectFlag {
			fmt.Printf("result: '%s'\n", Input("Input Test", *inputArgCmd))
		} else {
			fmt.Printf("select: '%s'\n", Select("Select Test", *inputArgCmd))
		}
	default:
		return false
	}
	return true
}

func parse(cmd string) {
	// debugging (hidden) commands
	if parseDebug(cmd) {
		return
	}
	// user visible commands
	switch cmd {
	// Start
	case startCmd.FullCommand():
		err := Start(*startDirArg)
		ShowError(err)
		if err == nil {
			time.Sleep(2 * time.Second)
			browser.OpenURL(BrowserURL)
		}
	// Stop
	case stopCmd.FullCommand():
		Stop()
	// Init
	case initCmd.FullCommand():
		dir, err := filepath.Abs(*initDirArg)
		if err == nil {
			log.Debug("init dir ", dir)
			err := Init(dir, *initRepoArg, os.Stderr)
			if err == nil {
				log.Debug("configuring")
				err = Configure(dir)
			}
		}
		ShowError(err)
	// Status
	case statusCmd.FullCommand():
		dockerStatus("openwhisk")
		dockerStatus("redis")
		dockerStatus("ide-js")
	default:
		kingpin.Usage()
	}
}

// Main entrypoint for wskide
func Main() {
	cmd := kingpin.Parse()
	fmt.Println(cmd)
	if cmd != initCmd.FullCommand() {
		if err := ConfigLoad(); err != nil {
			log.Info(err)
			fmt.Println("You need to run 'iosdk init ', first.")
			os.Exit(1)
		}
	}
	if *debugFlag {
		log.SetLevel(log.DebugLevel)
	}
	if *traceFlag {
		log.SetLevel(log.TraceLevel)
	}
	parse(cmd)
}
