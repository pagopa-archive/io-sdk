# Tips for debugging tests

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



