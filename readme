# bandwidth-go

This is a Go client library for [Bandwidth's Communication APIs](https://dev.bandwidth.com/).

At the time of writing, Bandwidth does not provide an official Go SDK for their APIs. Please check the official [Brandwidth SDKs](https://dev.bandwidth.com/sdks) page for updates and a possible release of an official Go SDK in the future.

For now, this library only contains a package for generating [BXML](https://dev.bandwidth.com/docs/voice/bxml). As a base to work from, I copied the Twiml package from [Twilio's Go SDK](https://github.com/twilio/twilio-go).

## ⚠️ WARNING! ⚠️ 

This library is still in active development and is not ready for production use. Be aware that not all BXML verbs and nouns have been fully implemented yet. Because this project started as a copy of Twilio's Twiml package, some of the Twiml verbs and nouns may still be around. If you would like to contribute, please see feel free to submit a pull request.

## Installation

The recommended way to install `bandwidth-go` is by using [Go modules](https://go.dev/ref/mod#go-get).

If you already have an initialized project, you can run the command below from your terminal in the project directory to install the library:

```shell
go get github.com/andersryanc/bandwidth-go
```

If you are starting from scratch in a new directory, you will first need to create a go.mod file for tracking dependencies such as bandwidth-go. This is similar to using package.json in a Node.js project or requirements.txt in a Python project. [You can read more about mod files in the Go documentation](https://golang.org/doc/modules/managing-dependencies). To create the file, run the following command in your terminal:

```shell
go mod init bandwidth-example
```

Once the module is initialized, you may run the installation command from above, which will update your go.mod file to include bandwidth-go.
