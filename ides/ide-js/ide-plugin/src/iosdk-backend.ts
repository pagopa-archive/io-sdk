/**
 * Generated using theia-plugin-generator
 */

import * as theia from '@theia/plugin';

export function start(context: theia.PluginContext) {
    const informationMessageTestCommand = {
        id: 'hello-world-example-generated',
        label: "build"
    };
    const options: theia.TerminalOptions = {
        name: "Bash terminal",
        shellPath: "/bin/bash",
        shellArgs: ["-l"],
        cwd: "/home/theia/plugins/ide-plugin",
        env: { "TERM": "screen", "TEST": "prova" },
    };
    context.subscriptions.push(theia.commands.registerCommand(informationMessageTestCommand, (...args: any[]) => {
        theia.window.showInformationMessage('Hello World!');
        const terminal = theia.window.createTerminal(options);
        terminal.show();
        terminal.sendText("yarn build", true);
    }));

}

export function stop() {

}
