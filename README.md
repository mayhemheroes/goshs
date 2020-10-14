![Version](https://img.shields.io/badge/Version-v0.0.4-green) [![GitHub](https://img.shields.io/github/license/patrickhener/goshs)](https://github.com/patrickhener/goshs/blob/master/LICENSE) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/patrickhener/goshs) [![GitHub issues](https://img.shields.io/github/issues-raw/patrickhener/goshs)](https://github.com/patrickhener/goshs/issues)

<img src="https://github.com/patrickhener/image-cdn/blob/main/goshs-logo-github.png" alt="goshs-logo" width="85">

goshs is a replacement for Python's `SimpleHTTPServer`. It allows uploading and downloading via HTTP/S with either self-signed certificate or user provided certificate and you can use HTTP basic auth.

> <kbd><img src="https://github.com/patrickhener/image-cdn/blob/main/goshs-screenshot.png" alt="goshs-screenshot"></kbd>

# Installation

## Release
You can download the executable from the [release section](https://github.com/patrickhener/goshs/releases)

## Go

```bash
go get -u github.com/patrickhener/goshs
go install github.com/patrickhener/goshs
```

## Build yourself

```bash
git clone https://github.com/patrickhener/goshs.git
cd goshs
make build
```

# Usage

```bash
goshs -h
Usage of goshs:
  -P string
    	Use basic auth password (user: gopher)
  -d string
    	Web root directory (default "")
  -p int
    	The port (default 8000)
  -s	Use self-signed TLS
  -sc string
    	Path to own server cert
  -sk string
    	Path to own server key
  -ss
    	Use self-signed certificate
```

# Examples

**Serve from your current directory**

`goshs`

**Serve from another directory**

`goshs -d /path/to/directory`

**Serve from port 1337**

`goshs -p 1337`

**Password protect the service**

`goshs -P VeryS3cureP4$$w0rd`

*Please note:* goshs uses HTTP basic authentication. It is recommended to use SSL option with basic authentication to prevent from credentials beeing transfered in cleartext over the line. User is `gopher`.

**Use TLS connection**

*Self-Signed*

`goshs -s -ss`

*Provide own certificate*

`goshs -s -sk server.key -sc server.crt`

# Credits

A special thank you goes to *sc0tfree* for inspiring this project with his project [updog](https://github.com/sc0tfree/updog) written in Python.

# Tutorial Series

I wrote several blog posts how and why I implemented all of this. You can find it [here](https://hesec.de/tags/goshs/) if you are intereseted about the technical background.