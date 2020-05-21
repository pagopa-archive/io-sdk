# IO-SDK development process

![Test Build](https://github.com/pagopa/io-sdk/workflows/Test%20Build/badge.svg)
![Create Release And Upload Assets](https://github.com/pagopa/io-sdk/workflows/Create%20Release%20And%20Upload%20Assets/badge.svg)

# Using IO-SDK as a developer

If you are one of those heroic people wanting to help and you can deal with bugs, crashes and the occasional complete destruction of you computer, you can follow with those instructions. 

## Prerequisites

Install 
- Go v1.13
- Nodejs v10.x 
- a recent version of GNU MAKE
- [OpenWhisk Cli v1.0.0](https://github.com/apache/openwhisk-cli/releases)
- [Docker](https://www.docker.com/get-started)

and be sure it is available on the path. If you don't know how to install all of this, well, I doubt you can develop IO-SDK

## Start Development Environment

Execute the following commands

````
git clone https://github.com/pagopa/io-sdk
cd io-sdk
````

Now you can build images locally:

```
make build
```

You can now start development. Do:

```
./iosdk/iosdk init --wskprops
./iosdk/iosdk start --skip-pull-images
```

Note the `--wskprops` that will setup locally the configuration file to access and deploy actions in OpenWhisk.
Note also the `--skip-pull-images` that will use the local copy without downloading images from the net

If you are lucky, your browser will open and you will get the IO-SDK user interface.

Try also `wsk action list` and you should see the list of actions

## Admin User Interface

To develop the admin user interface, do:

```
cd backend
make devel
```

Now you can  can find a development version of the UI in `http://localhost:5000`.

Code is writtend javascript, based on [Svelte](https://svelte.dev/). Sources are under `src`

You can edit codde and it will automatically reload.

## Backend Actions

The server side part of the backend are [OpenWhisk](http://openwhisk.apache.org), action written in various languages. 

Source code is under `actions/src`

You can now deploy actions after changes with `make deploy`

## The CLI

The CLI, that is used to launch and initialize the environment is under `iosdk` and it is written in the Go programming language.

To develop the cli, do `cd iosdk` and use `make`. In the cli there is the version number of the images. In development it is just the branch name. In production the build actions will build images and match them with the version store in the cli to retrieve the right version of the image.

There is a test suite and you generally should ensure tests are passed with `make test`.

## Releases

You can build "branch" images and binaries just typing "make" at the top level. When you build a branch image it is built and pushed to docker hub tagged with the branch name. You can then use those images with `go get github.com/pagopa/io-sdk/iosdk#branch`.

Finally you build "snapshot" images and installers. Those images will be tagged with a  snapshot tag from date and time (in format `%Y.%m%d.%H%M-snapshot`) . This will also build installers for each platform.

You can build "release" images and installers with `make VER=X.Y.Z-T release`. 

## Base images

The server, the ide and the installer all uses some base images, that are pretty stable and change rarely.

If you really need to change them, you need to get a password to write in some public repository (`pagopa` on DockerHub currently)

`cd images ; make`

They are tagged with the date you built them. If you change the tag you will also have to change the tag in `admin/Dockerfile` and `ide/Dockerfile`

The process is manual because it is infrequent and needs awareness of what you are doing.


