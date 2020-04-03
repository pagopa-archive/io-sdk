package wskide

import (
	"errors"
	"fmt"
	"regexp"
)

func ExampleLogIf() {
	fmt.Println(LogIf(nil))
	err := fmt.Errorf("generated error")
	if LogIf(err) {
		fmt.Println(LastError)
	}
	// Output:
	// false
	// generated error
}

func shouldPanic(flag bool) (err error) {
	defer Recover(&err)
	if flag {
		err := errors.New("Panic mode")
		FatalIf(err)
	}
	fmt.Println("no panic")
	return nil
}

func ExampleFatalIf() {
	fmt.Println(shouldPanic(false))
	fmt.Println(shouldPanic(true))
	// Output:
	// no panic
	// <nil>
	// Panic mode
}

func ExampleSys() {
	*DryRunFlag = false
	Sys("/bin/echo 1 2 3")
	Sys("/bin/echo 3", "4", "5")
	Sys("@sh -c", "echo foo >/tmp/foo")
	fmt.Print(Sys("cat /tmp/foo"))
	Sys("@sh -c", "echo bar >/tmp/bar")
	fmt.Print(Sys("@cat /tmp/bar"))
	*DryRunFlag = true
	// Output:
	// 1 2 3
	// 3 4 5
	// foo
	// foo
	// bar
}

func ExampleSysErr() {
	*DryRunFlag = false
	SysErr("/bin/echo 1 2 3")
	SysErr("/bin/echo 3", "4", "5")
	SysErr("@sh -c", "echo foo >/tmp/foo")
	fmt.Print(Sys("cat /tmp/foo"))
	SysErr("@sh -c", "echo bar >/tmp/bar")
	fmt.Print(Sys("@cat /tmp/bar"))
	_, err := SysErr("false")
	fmt.Println("ERR", err)
	_, err = SysErr("donotexist")
	fmt.Println("ERR", err)
	*DryRunFlag = true
	// Output:
	// 1 2 3
	// 3 4 5
	// foo
	// foo
	// bar
	// ERR exit status 1
	// ERR exec: "donotexist": executable file not found in $PATH
}

func ExampleSysCd() {
	*DryRunFlag = false
	Sys("@mkdir -p /tmp/example-syscd/hello-i-am-a-dir")
	fmt.Println(SysCd("/tmp/example-syscd", "@ls"))
	*DryRunFlag = true
	// Output:
	// hello-i-am-a-dir
}

func ExampleSysSh() {
	*DryRunFlag = false
	show(SysSh("@rm -Rf /tmp/simple ; mkdir /tmp/simple ; ls -d /tmp/simple"))
	SysSh("cd /tmp/simple ; touch hello ; ls")
	*DryRunFlag = true
	// Output:
	// /tmp/simple
	//
	// hello
}

func ExampleDryRun() {
	DryRunPush("first", "second")
	fmt.Println(Sys("dummy"))
	fmt.Println(Sys("dummy", "alpha", "beta"))
	fmt.Printf("'%s'\n", Sys("dummy"))
	// Output:
	// dummy
	// first
	// dummy alpha beta
	// second
	// dummy
	// ''
}

func ExampleDryRunErr() {
	DryRunPush("first", "second", "!third")
	out, err := SysErr("dummy")
	fmt.Println(1, out, err)
	out, err = SysErr("dummy", "alpha", "beta")
	fmt.Println(2, out, err)
	out, err = SysErr("dummy")
	fmt.Println(3, "out", out, "err", err)
	out, err = SysErr("dummy")
	fmt.Println(4, "out", out, "err", err)
	// Output:
	// dummy
	// 1 first <nil>
	// dummy alpha beta
	// 2 second <nil>
	// dummy
	// 3 out  err third
	// dummy
	// 4 out  err <nil>
}

func ExampleShowError() {
	ShowError(nil)
	ShowError(errors.New("error"))
	// Output:
	// *** ERROR: error ***
}

func ExampleRun() {
	//*DryRunFlag = false
	DryRunPush("", "", "wrong")
	fmt.Println(Run("docker pull busybox"))
	fmt.Println(Run("@docker pull busybox"))
	fmt.Println(Run("@docker pull busybox1"))
	// Output:
	// docker pull busybox
	// <nil>
	// <nil>
	// wrong
}

func ExampleRandomString() {
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	rnd := RandomString(32)
	fmt.Printf("%s\n", re.ReplaceAll([]byte(rnd), []byte("*")))
	// Output:
	// ********************************
}

func ExampleInput() {
	DryRunPush("hello", "world")
	fmt.Println(Input("how are you", "fine"))
	fmt.Println(Input("", ""))
	// Output:
	// hello
	// world
}

func ExampleSelect() {
	DryRunPush("hello", "world")
	fmt.Println(Select("how are you", "fine"))
	fmt.Println(Select("", ""))
	// Output:
	// hello
	// world
}
