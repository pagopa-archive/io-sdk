
# IO-SDK

The app IO is the App to access Italian Government services.

The IO-SDK is a software development kit to develop integrations with the app IO.

This SDK allows to easily develop importers to send messages to italian citizens using IO.

This document is about how to start to use the sdk.

If you want to contribute, check the [development](DEVEL.md) document.

## Using IO-SDK 

To use IO-SDK you need to install Docker Desktop in Windows or Mac, or just Docker in Linux.

You also need an API Key for IO.

Download and install one of the releases `iosdk`.

If you have `go`, you can build the bleeding edge (master) with `go get github/pagopa/io-sdk/iosdk`.

Master is supposed to be stable as development should happen in branches.

Initialize the environment with `iosdk init`.

It will ask for:

- the work directory where you are going put your code. It *must* be below your home directory
- a template to use, either one of the available templates to import Excel, SQL or REST data, or any third-parties templates that will (eventually) be available
- the IO Api Key

Once configured you can start the sdk with `iosdk start`. 
It will then open the user interface at `http://localhost:3280`.

Other commands are `iosdk status` to check the status or `iosdk stop` to stop the sdk.

You can check how the sdk works [in this video](https://youtu.be/qnnCg_cIwWc) (it is in Italian). 
