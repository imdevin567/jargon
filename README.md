![Logo][jargon-logo]

[![Build Status](https://travis-ci.org/imdevin567/jargon.svg?branch=master)](https://travis-ci.org/imdevin567/jargon)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/imdevin567/jargon?status.svg)](http://godoc.org/github.com/imdevin567/jargon)
[![Go Report Card](https://goreportcard.com/badge/github.com/imdevin567/jargon)](https://goreportcard.com/report/github.com/imdevin567/jargon)
![GitHub release](https://img.shields.io/github/release/imdevin567/jargon.svg)

> A proxy tool to translate network traffic to different protocols.

TODO: Description of Jargon

## Get Started

1. Install Jargon

```bash
$ go get github.com/imdevin567/jargon
```

2. Create a `Jargonfile` with route configuration

```yaml
---
routes:
  - name: HTTP to TCP
    input:
      adapter: http
      host: localhost
      port: 9600
      path: /post
      contentType: application/json
    output:
      adapter: tcp
      host: localhost
      port: 8861
      delimiter: \n
```

3. Start Jargon

```bash
$ jargon run
```

You can verify setup by creating a TCP server with `nc`:

```bash
$ nc -l 8861
```

Now send some data to the HTTP endpoint created by Jargon:

```bash
$ curl -X POST -H "Content-Type: application/json" -d '{"active": true}' http://localhost:9600/post
```

You should receive the data on the TCP server:

```bash
$ nc -l 8861
{"active": true}
```

## Goals

* Protocol support
    * ~~HTTP~~
    * ~~TCP~~
    * ~~UDP~~
    * ~~Websocket~~
    * HTTPS
    * Secure websocket (TLS)
    * RPC
    * AMQP
* Docker image
* Multiplex input/output

## Contributing

Community contributions are always greatly appreciated. To start developing Jargon, check out the [guidelines for contributing](https://github.com/imdevin567/jargon/blob/master/.github/CONTRIBUTING.md).

## License

Jargon is released under MIT license. See [LICENSE](https://github.com/imdevin567/jargon/blob/master/LICENSE).

[jargon-logo]: /img/logo.png
