# IO-SDK development process

![Test Build](https://github.com/pagopa/io-sdk/workflows/Test%20Build/badge.svg)
![Create Release And Upload Assets](https://github.com/pagopa/io-sdk/workflows/Create%20Release%20And%20Upload%20Assets/badge.svg)

# Using IO-SDK as a developer

If you are one of those heroic people wanting to help and you can deal with bugs, crashes and the occasional complete destruction of your computer, you can follow with those instructions.

After you read this document, do not forget to [check also this other document](DEBUG.md) for debugging tips.

## Prerequisites

Supported development platforms are:
- Linux Ubuntu 18.x ([read installation guidelines](docs/Prerequisites/Linux/Ubuntu.md))
- Mac OS Catalina ([read installation guidelines](docs/Prerequisites/Mac/OSX.md))
- Windows 10 Professional build 2003 with WSL2 and Ubuntu 18.04 ([read installation guidelines](docs/Prerequisites/Windows/10.md))

Also your workstation needs at least 8gb of memory.

It may work on other configurations but it is not tested.

## First setup of the development environment

If the prerequisites are satisfied, you can make the first setup and test the development environment as follows:

```
git clone https://github.com/pagopa/io-sdk
cd io-sdk
bash setup.sh
source source-me-first
make
````

In detail, each command performs these activities:
- download the sources from github
- enter in the repository folder
- install dependencies with `./setup.sh`
- activate and verify the environment with `source source-me-first`
- build everything with `make`

If all the test passes, congratulations, your development environment is ready.

**If this procedure does not work in one of the supported installations it is a bug, please [report it](https://github.com/pagopa/io-sdk/issues).**

## Daily development tasks (after first setup)

In general you do not need to repeat the entire setup procedure every day.

If everything was setup, just type:  `source source-me-first` to setup the environment.

You can update images with:

```
make build
```

You can start development services. Do:

```
./iosdk/iosdk init --wskprops
./iosdk/iosdk start --skip-pull-images --skip-ide
```

Note the `--wskprops` that will setup locally the configuration file to access and deploy actions in OpenWhisk.
Note also the `--skip-pull-images` that will use the local copy without downloading images from the net.
Note the command `--skip-ide` exclude the theia ide for edit code direcly on the application.

If you are lucky, your browser will open and you will get the IO-SDK user interface.

Try also `wsk action list` and you should see the list of actions

## Admin User Interface

To develop the admin user interface, do:

```
cd admin
make devel
```

Now you have a development version of the UI in `http://localhost:5000`.

**NOTE:** on windows with WSL2 you will need to setup a tunnel to reach that port. [See the section on installing on Windows](docs/Prerequisites/Windows/10.md).

Code is writtend javascript, based on [Svelte](https://svelte.dev/). Sources are under `src`

You can edit code and it will automatically reload.

Do `make build` from the top level to embed the code you modified in the image.

## Backend Actions

The server side part of the backend are [OpenWhisk](http://openwhisk.apache.org), action written in various languages.

Source code is under `admin/actions/src`

You can edit sources.

You can deploy actions after changes with `make deploy`

## The CLI

The CLI, that is used to launch and initialize the environment is under `iosdk` and it is written in the Go programming language.

To develop the cli, do `cd iosdk` and use `make`.

The cli embeds the version number of the images to refer to them correctly.

In development, the version is just the branch name.

When you push to master, the CI builds images and pushes to the Docker Hub, so the CLI can retrieve the version stored in Docker Hub.

There is a test suite for the CLI and you generally should ensure tests are passed with `make test`.

## Snapshot and Releases

Images are pushed to docker hub only by the CI, and they uses a secret token, so to get your changes in the public images on docker hub you need to open a Pul Requests.

If you have the token (reserved to release manager) you can do, after a `docker login` also:

- `make push` to push images tagged with current branch
- `make snapshot` to create a "snapshot" release (in format `%Y.%m%d.%H%M-snapshot`)
- `git tag <release> ; git push --tags` to create a release (a manual process)

## Base images

The server, the ide and the installer all uses some base images, that are pretty stable and change rarely.

If you really need to change them, you need to get a password to write in some public repository (`pagopa` on DockerHub currently)

`cd images ; make`

They are tagged with the date you built them. If you change the tag you will also have to change the tag in `admin/Dockerfile` and `ide/Dockerfile`

The process is manual because it is infrequent and needs awareness of what you are doing.
