package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/dixonwille/wmenu/v5"
	log "github.com/sirupsen/logrus"
	"github.com/tcnksm/go-input"
	"gopkg.in/alecthomas/kingpin.v2"
)

// LastError seen
var LastError error

// FatalIf panics if error is not nil
func FatalIf(err error) {
	LastError = err
	if err != nil {
		log.Error(err)
		panic(err)
	}
}

// LogIf logs a warning if the error is not nil
// returns true if the err is not nil
func LogIf(err error) bool {
	LastError = err
	if err != nil {
		log.Warn(err)
		return true
	}
	return false
}

// Debug messages
func Debug(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Recover recovers from a panic returning an error
func Recover(perr *error) {
	if r := recover(); r != nil {
		switch v := r.(type) {
		case error:
			*perr = v
		case string:
			*perr = fmt.Errorf("%s", v)
		default:
			*perr = fmt.Errorf("%v", v)
		}
	}
}

// DryRunFlag allows for dry execution
var DryRunFlag = kingpin.Flag("dry-run", "show commands instead of executing them").
	Short('n').Default("false").Bool()

// buffer for dry run rusults
var dryRunBuf = []string{}

// DryRunPush saves dummy results for dry run execution
func DryRunPush(buf ...string) {
	dryRunBuf = buf
}

// DryRunPop returns a value from the buffer of dry run results
// returns an empty string if the  buffer is empty
func DryRunPop(buf ...string) string {
	res := ""
	if len(dryRunBuf) > 0 {
		res = dryRunBuf[0]
		dryRunBuf = dryRunBuf[1:]
	}
	return res
}

// Sys executes a command in a convenient way:
// it splits the paramenter in arguments if separated by spaces,
// then accepts multiple arguments;
// logs errors in stderr and prints output in stdout;
// also returns output as a string, or empty if errors.
// If the command starts with "@" do not print the output.
// It also honor the DryRunFlag, in this case it always print the command
func Sys(cli string, args ...string) string {
	re := regexp.MustCompile(`[\r\t\n\f ]+`)
	a := strings.Split(re.ReplaceAllString(cli, " "), " ")
	params := args
	if len(a) > 1 {
		params = append(a[1:], args...)
	}
	exe := strings.TrimPrefix(a[0], "@")
	silent := strings.HasPrefix(a[0], "@")

	if *DryRunFlag {
		if len(params) > 0 {
			fmt.Printf("%s %s\n", exe, strings.Join(params, " "))
		} else {
			fmt.Println(exe)
		}
		return DryRunPop()
	}

	log.Tracef("< %s %v\n", exe, params)
	cmd := exec.Command(exe, params...)
	out, err := cmd.CombinedOutput()
	res := string(out)
	log.Tracef("> %s", res)
	if !LogIf(err) && !silent {
		fmt.Printf(res)
	}
	return res
}

// SysErr executes a command in a convenient way:
// it splits the paramenter in arguments if separated by spaces,
// then accepts multiple arguments;
// logs errors in stderr and prints output in stdout;
// also returns output as a string, or an error if there is an error
// If the command starts with "@" do not print the output.
// It also honor the DryRunFlag, in this case it always print the command
func SysErr(cli string, args ...string) (string, error) {
	re := regexp.MustCompile(`[\r\t\n\f ]+`)
	a := strings.Split(re.ReplaceAllString(cli, " "), " ")
	params := args
	if len(a) > 1 {
		params = append(a[1:], args...)
	}
	exe := strings.TrimPrefix(a[0], "@")
	silent := strings.HasPrefix(a[0], "@")
	if *DryRunFlag {
		if len(params) > 0 {
			fmt.Printf("%s %s\n", exe, strings.Join(params, " "))
		} else {
			fmt.Println(exe)
		}
		res := DryRunPop()
		if strings.HasPrefix(res, "!") {
			return "", errors.New(res[1:])
		}
		return res, nil
	}

	log.Tracef("< %s %v\n", exe, params)
	cmd := exec.Command(exe, params...)
	out, err := cmd.CombinedOutput()
	res := string(out)
	if err != nil {
		log.Tracef("> ERROR: %s", err.Error())
		return "", err
	}
	log.Tracef("> %s", res)
	if !silent {
		fmt.Printf(res)
	}
	return res, nil
}

// Run executes a command, without capturing input, in a convenient way:
// it splits the paramenter in arguments if separated by spaces,
// then accepts multiple arguments;
//returns the error
// It also honor the DryRunFlag, in this case it always prints the command
func Run(cli string, args ...string) error {
	re := regexp.MustCompile(`[\r\t\n\f ]+`)
	a := strings.Split(re.ReplaceAllString(cli, " "), " ")
	params := args
	if len(a) > 1 {
		params = append(a[1:], args...)
	}
	exe := strings.TrimPrefix(a[0], "@")
	silent := strings.HasPrefix(a[0], "@")

	if !silent {
		if len(params) > 0 {
			fmt.Printf("%s %s\n", exe, strings.Join(params, " "))
		} else {
			fmt.Println(exe)
		}
	}

	if *DryRunFlag {
		errs := DryRunPop()
		if errs == "" {
			return nil
		}
		return errors.New(errs)
	}

	log.Tracef("< %s %v\n", exe, params)
	cmd := exec.Command(exe, params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// SysCd works as Sys,
// but cd into the directory first,
// and restore the original directory after.
func SysCd(cd string, cli string, args ...string) string {
	orig, err := os.Getwd()
	FatalIf(err)
	FatalIf(os.Chdir(cd))
	res := Sys(cli, args...)
	FatalIf(os.Chdir(orig))
	return res
}

// SysSh execute a command as a shell script;
// if it starts with "@" it does not print the output;
// returns the output as a string.
func SysSh(cmd string) string {
	if strings.HasPrefix(cmd, "@") {
		return Sys("@sh -c", strings.TrimPrefix(cmd, "@"))
	}
	return Sys("sh -c", cmd)
}

// ShowError prints an error if not nil
func ShowError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func mkMap(key string, value interface{}) map[string]interface{} {
	return map[string]interface{}{
		key: value,
	}
}

func addMap(data map[string]interface{}, key string, value interface{}) map[string]interface{} {
	data[key] = value
	return data
}

func mkErr(err interface{}) map[string]interface{} {
	switch v := err.(type) {
	case error:
		return mkMap("error", v.Error())
	case string:
		return mkMap("error", v)
	default:
		return mkMap("error", fmt.Sprintf("%v", err))
	}
}

// RandomString generate random string of given lengtth
func RandomString(length int) string {

	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
		"1234567890"

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Input asks for an input,
// uses predef as the default
// if you use the empty string it will ask for a value
// if the user press ^c it returns the empty string
func Input(query string, predef string) string {
	if *DryRunFlag {
		return DryRunPop()
	}
	inputUI := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}
	input, err := inputUI.Ask(query, &input.Options{
		Default:  predef,
		Required: true,
		Loop:     true,
	})
	if err != nil {
		return ""
	}
	return input
}

// Select asks for an input,
// uses predef as the options, comma separated
// if the user press ^c it returns the empty string
func Select(query string, options string) string {
	if *DryRunFlag {
		return DryRunPop()
	}
	inputUI := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}
	sel := strings.Split(options, ",")
	input, err := inputUI.Select(query, sel, &input.Options{
		Default:  sel[0],
		Required: true,
		Loop:     true,
	})
	if err != nil {
		return ""
	}
	return input
}

// SelectTemplate asks for template selection and build the corresponding github URI
func SelectTemplate(tmpls map[string]string) (string, error) {
	const baseURL = "https://github.com/"
	const defaultProj = "pagopa/io-sdk"
	var resURI = ""

	// Define header
	menu := wmenu.NewMenu("Please select the template:")

	// Define the default FQDN {githubURL}/{pagopaNAME}/{templateNAME}
	menu.Action(func(opts []wmenu.Opt) error {
		resURI = fmt.Sprintf("%s/%v-%v", baseURL, defaultProj, opts[0].Text)
		return nil
	})

	// Enumerate the input options. Only GH has a custom action, asking for <user>/<project>
	menu.Option(tmpls["js"], tmpls["js"], true, nil)
	menu.Option(tmpls["java"], tmpls["java"], false, nil)
	menu.Option(tmpls["py"], tmpls["py"], false, nil)
	menu.Option(tmpls["gh"], tmpls["gh"], false, func(o wmenu.Opt) error {
		fmt.Println("Enter the github user/path:")
		input := bufio.NewReader(os.Stdin)
		customProj, err := input.ReadString('\n')

		// detect OS newline, eventually flushing it
		if runtime.GOOS == "windows" {
			customProj = strings.Replace(customProj, "\r\n", "", -1)
		} else {
			customProj = strings.Replace(customProj, "\n", "", -1)
		}

		resURI = fmt.Sprintf("%v/%v", baseURL, customProj)
		return err
	})

	// Launch the menu
	err := menu.Run()

	if err != nil {
		return "", err
	}

	return resURI, nil
}
