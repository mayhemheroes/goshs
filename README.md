![Version](https://img.shields.io/badge/Version-v0.3.1-green)
[![GitHub](https://img.shields.io/github/license/patrickhener/goshs)](https://github.com/patrickhener/goshs/blob/master/LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/patrickhener/goshs)
[![GitHub issues](https://img.shields.io/github/issues-raw/patrickhener/goshs)](https://github.com/patrickhener/goshs/issues)
![goreleaser](https://github.com/patrickhener/goshs/workflows/goreleaser/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/patrickhener/goshs)](https://goreportcard.com/report/github.com/patrickhener/goshs)

<img src="https://github.com/patrickhener/image-cdn/blob/main/goshs-logo-github.png" alt="goshs-logo" width="85">

goshs is a replacement for Python's `SimpleHTTPServer`. It allows uploading and downloading via HTTP/S with either self-signed certificate or user provided certificate and you can use HTTP basic auth.

<kbd><img src="https://github.com/patrickhener/image-cdn/blob/main/goshs-screenshot.png" alt="goshs-screenshot"></kbd>

# Features
* Download or view files
  * Bulk download as .zip file
* Upload files (Drag & Drop)
* Basic Authentication
* Transport Layer Security (HTTPS)
  * self-signed
  * provide own certificate
* Non persistent clipboard
  * Download clipboard entries as .json file
* WebDAV support
* Read-Only and Upload-Only mode
* Silent mode (no webserver output)
* Retrieve json on cli

# Installation

## Release
You can download the executable from the [release section](https://github.com/patrickhener/goshs/releases)

## Go

```bash
go get -u github.com/patrickhener/goshs
go install github.com/patrickhener/goshs@latest
```

## Build yourself

Building requirements are [ugilfy-js](https://www.npmjs.com/package/uglify-js) and [wt](https://github.com/wellington/wellington). After installing those 2 packages you can easily just:

```bash
git clone https://github.com/patrickhener/goshs.git
cd goshs
make build
```

# Usage

```bash
> goshs -h

goshs v0.3.1
Usage: goshs [options]

Web server options:
  -i,  --ip           The ip/if-name to listen on             (default: 0.0.0.0)
  -p,  --port         The port to listen on                   (default: 8000)
  -d,  --dir          The web root directory                  (default: current working path)
  -w,  --webdav       Also serve using webdav protocol        (default: false)
  -wp, --webdav-port  The port to listen on for webdav        (default: 8001)
  -ro, --read-only    Read only mode, no upload possible      (default: false)
  -uo, --upload-only  Upload only mode, no download possible  (default: false)
  -si, --silent       Running without dir listing             (default: false)

TLS options:
  -s,  --ssl          Use TLS
  -ss, --self-signed  Use a self-signed certificate
  -sk, --server-key   Path to server key
  -sc, --server-cert  Path to server certificate

Authentication options:
  -b, --basic-auth    Use basic authentication (user:pass - user can be empty)

Misc options:
  -V  --verbose       Activate verbose log output             (default: false)
  -v                  Print the current goshs version

Usage examples:
  Start with default values:    	./goshs
  Start with wevdav support:    	./goshs -w
  Start with different port:    	./goshs -p 8080
  Start with self-signed cert:  	./goshs -s -ss
  Start with custom cert:       	./goshs -s -sk <path to key> -sc <path to cert>
  Start with basic auth:        	./goshs -b secret-user:$up3r$3cur3
  Start with basic auth empty user:	./goshs -b :$up3r$3cur3
```

# Examples

**Serve from your current directory**

`goshs`

**Serve from your current directory with webdav enabled on custom port**

`goshs -w -wp 8081`

**Serve from another directory**

`goshs -d /path/to/directory`

**Serve from port 1337**

`goshs -p 1337`

**Password protect the service**

`goshs -b secret-user:VeryS3cureP4$$w0rd`

*Please note:* goshs uses HTTP basic authentication. It is recommended to use SSL option with basic authentication to prevent from credentials beeing transfered in cleartext over the line.

**Use TLS connection**

*Self-Signed*

`goshs -s -ss`

*Provide own certificate*

`goshs -s -sk server.key -sc server.crt`

**Run in silent mode**  
This mode will omit the dir listing on the web interface. Also you will not have access to the clipboard or upload form. Still you could upload a file using the corresponding post request (see examples).

`goshs -si`

**Retrieve the directory listing in json format**  
You can now retrieve the directory listing in *json* format. This is meant to be used with curl for example in environments where you do not have a browser on hand.
```bash
curl -s localhost:8000/?json | jq
[
  {
    "name": ".git/",
    "is_dir": true,
    "is_symlink": false,
    "symlink_target": "",
    "extension": "",
    "size_bytes": 4096,
    "last_modified": "2023-02-28T15:38:11.982+01:00"
  },
  {
    "name": ".github/",
    "is_dir": true,
    "is_symlink": false,
    "symlink_target": "",
    "extension": "",
    "size_bytes": 4096,
    "last_modified": "2023-02-28T10:27:35.524+01:00"
  },
  {
    "name": ".gitignore",
    "is_dir": false,
    "is_symlink": false,
    "symlink_target": "",
    "extension": ".gitignore",
    "size_bytes": 48,
    "last_modified": "2023-02-20T07:58:46.436+01:00"
  },
  ... snip ...
```

Or with path:

```bash
curl -s localhost:8000/utils?json | jq
[
  {
    "name": "utils.go",
    "is_dir": false,
    "is_symlink": false,
    "symlink_target": "",
    "extension": ".go",
    "size_bytes": 2218,
    "last_modified": "2023-02-28T15:28:54.783+01:00"
  },
  {
    "name": "utils_test.go",
    "is_dir": false,
    "is_symlink": false,
    "symlink_target": "",
    "extension": ".go",
    "size_bytes": 2012,
    "last_modified": "2023-02-28T15:28:12.748+01:00"
  }
]
```


# Credits

A special thank you goes to *sc0tfree* for inspiring this project with his project [updog](https://github.com/sc0tfree/updog) written in Python.
