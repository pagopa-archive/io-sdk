# Tips to help debugging

## Hidden flags

The `iosdk` has a number of hidden flags useful for development

Global flags

- `--use-default-api-key`: do not generate an api key use the default one
- `--skip-docker-version`: do not check the docker version
- `--skip-pull-images`: do not pull images
- `--skip-open-browser`: no not open the browser
- `--skip-ide`: do not start the ide

Hidden commands:

- `wskprops`: set the .wskprops to use `wsk` with the server
- `ide-deploy`: deploy the ide
- `ide-destroy`: destroy the ide
- `whisk-deploy`: deploy whisk
- `whisk-destroy`: destroy whisk
- `redis-deploy`: deploy redis
- `redis-destroy`: destroy redis

## Checking ide test results

You can run the test suite of the cli with 

```
cd iosdk
make test
```

If a test fails it may be difficult to understand what is wrong in the output.

The test command saves the test output in `difftest.out`

The `difftest.py` comes to the rescue, you pass the test output and returns a list of failed tests , with an index:

```
$ python3 difftest.py difftest.out
0 --- FAIL: ExampleIdeDockerRun (0.00s)
1 --- FAIL: ExampleStart (1.03s)  
```

Then you can see what went wrong in a single test passing the index:

```
$ python3 difftest.py difftest.out 1
31c31
< docker run -d -p 3000:3000 --rm --name iosdk-theia -e HOME=/home/project --add-host=openwhisk: 172.17.0.2 -v /tmp/iosdk-test/javascript:/home/project pagopa/iosdk-theia:test 
---
> docker run -d -p 3000:3000 --rm --name iosdk-theia --add-host=openwhisk:172.17.0.2 -v /tmp/iosdk-test/javascript:/home/project pagopa/iosdk-theia:test
```

## Debugging action tests

Bats tests are shell scripts.

The easiest way to debug them is to execute them line by line using VSCode.

First, configure a keybinding to execute a line in the terminal.

Open the `keybindings.json` as [described here](https://code.visualstudio.com/docs/getstarted/keybindings#_advanced-customization) and add this snippet:

```
{
  "key": "ctrl+enter",
  "command": "workbench.action.terminal.runSelectedText",
  "when": "editorTextFocus"
}
```

Open a terminal, change to `admnin/actions` then do a `source debug-src`.

Open a test and send each line to the terminal with control-enter.



