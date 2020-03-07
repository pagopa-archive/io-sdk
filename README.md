# io-sdk

![Test Build](https://github.com/pagopa/io-sdk/workflows/Test%20Build/badge.svg)
![Create Release And Upload Assets](https://github.com/pagopa/io-sdk/workflows/Create%20Release%20And%20Upload%20Assets/badge.svg)

# Using IO-SDK as an user

Whenm ready, you can go to the release page and download the binary for your system. 
If you do it now, as everything is in flux, you may have something completely broken. So don't do it, yet.

# Using IO-SDK as a developer

If you are one of those heroic people wanting to help and you can deal with bugs, crashes and the occasional complete destruction of you computer, you can follow with those instructions. 

## Prerequisites

Install 
- Go v1.13
- Nodejs v10.x  
- a recent version of GNU MAKE
- [OpenWhisk Cli v1.0.0](https://github.com/apache/openwhisk-cli/releases)

and be sure it is available on the path. If you don't know how to do this, well, I doubt you can help...

````
git clone https://github.com/pagopa/io-sdk
cd io-sdk
make
````

You can now start development. Do:

```
bin/iosdk init driver
bin/iosdk start driver
```

If you are lucky, your browser will open and you will get the IO-SDK user interface.

## Backend User Interface

To develop the backend user interface, do:

```
cd backend
make devel
```

You can find a development version of the UI in `http://localhost:5000`.

Code is writtend javascript, based on [Svelte](https://svelte.dev/). Sources are under `src``

## Backend Actions

The server side part of the backend are [OpenWhisk](http://openwhisk.apache.org), action written in various languages. 

Source code is under `actions/src`

You can now deploy actions after changes with `make deploy`



