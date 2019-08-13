# anonuuid

:wrench: anonymize UUIDs

![anonuuid Logo](https://raw.githubusercontent.com/moul/anonuuid/master/assets/anonuuid.png)

[![CircleCI](https://circleci.com/gh/moul/golang-repo-template.svg?style=shield)](https://circleci.com/gh/moul/golang-repo-template)
[![GoDoc](https://godoc.org/moul.io/golang-repo-template?status.svg)](https://godoc.org/moul.io/golang-repo-template)
[![License](https://img.shields.io/github/license/moul/golang-repo-template.svg)](https://github.com/moul/golang-repo-template/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/moul/golang-repo-template.svg)](https://github.com/moul/golang-repo-template/releases)
[![Go Report Card](https://goreportcard.com/badge/moul.io/golang-repo-template)](https://goreportcard.com/report/moul.io/golang-repo-template)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/golang-repo-template/badge)](https://www.codefactor.io/repository/github/moul/golang-repo-template)
[![codecov](https://codecov.io/gh/moul/golang-repo-template/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/golang-repo-template)
[![Docker Metrics](https://images.microbadger.com/badges/image/moul/golang-repo-template.svg)](https://microbadger.com/images/moul/golang-repo-template)
[![Sourcegraph](https://sourcegraph.com/github.com/moul/anonuuid/-/badge.svg)](https://sourcegraph.com/github.com/moul/anonuuid?badge)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)


**anonuuid** anonymize an input string by replacing all UUIDs by an anonymized
new one.

The fake UUIDs are cached, so if `anonuuid` encounter the same real UUIDs multiple
times, the translation will be the same.

## Usage

```console
$ anonuuid --help
NAME:
   anonuuid - Anonymize UUIDs outputs

USAGE:
   anonuuid [global options] command [command options] [arguments...]

VERSION:
   1.0.0-dev

AUTHOR(S):
   Manfred Touron <https://github.com/moul>

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --hexspeak		Generate hexspeak style fake UUIDs
   --random, -r		Generate random fake UUIDs
   --keep-beginning	Keep first part of the UUID unchanged
   --keep-end		Keep last part of the UUID unchanged
   --prefix, -p 	Prefix generated UUIDs
   --suffix 		Suffix generated UUIDs
   --help, -h		show help
   --version, -v	print the version
   ```

## Example

Replace all UUIDs and cache the correspondance.

```command
$ anonuuid git:(master) ✗ cat <<EOF | anonuuid
VOLUMES_0_SERVER_ID=15573749-c89d-41dd-a655-16e79bed52e0
VOLUMES_0_SERVER_NAME=hello
VOLUMES_0_ID=c245c3cb-3336-4567-ada1-70cb1fe4eefe
VOLUMES_0_SIZE=50000000000
ORGANIZATION=fe1e54e8-d69d-4f7c-a9f1-42069e03da31
TEST=15573749-c89d-41dd-a655-16e79bed52e0
EOF
VOLUMES_0_SERVER_ID=00000000-0000-0000-0000-000000000000
VOLUMES_0_SERVER_NAME=hello
VOLUMES_0_ID=11111111-1111-1111-1111-111111111111
VOLUMES_0_SIZE=50000000000
ORGANIZATION=22222222-2222-2222-2222-222222222222
TEST=00000000-0000-0000-0000-000000000000
```

---

Inline

```command
$ echo 'VOLUMES_0_SERVER_ID=15573749-c89d-41dd-a655-16e79bed52e0 VOLUMES_0_SERVER_NAME=bitrig1 VOLUMES_0_ID=c245c3cb-3336-4567-ada1-70cb1fe4eefe VOLUMES_0_SIZE=50000000000 ORGANIZATION=fe1e54e8-d69d-4f7c-a9f1-42069e03da31 TEST=15573749-c89d-41dd-a655-16e79bed52e0' | ./anonuuid
VOLUMES_0_SERVER_ID=00000000-0000-0000-0000-000000000000 VOLUMES_0_SERVER_NAME=bitrig1 VOLUMES_0_ID=11111111-1111-1111-1111-111111111111 VOLUMES_0_SIZE=50000000000 ORGANIZATION=22222222-2222-2222-2222-222222222222 TEST=00000000-0000-0000-0000-000000000000
```

---

```command
$ curl -s https://api.pathwar.net/achievements\?max_results\=2 | anonuuid | jq .
{
  "_items": [
    {
      "_updated": "Thu, 30 Apr 2015 13:00:58 GMT",
      "description": "You",
      "_links": {
        "self": {
          "href": "achievements/00000000-0000-0000-0000-000000000000",
          "title": "achievement"
        }
      },
      "_created": "Thu, 30 Apr 2015 13:00:58 GMT",
      "_id": "00000000-0000-0000-0000-000000000000",
      "_etag": "b1e9f850accfcb952c58384db41d89728890a69f",
      "name": "finish-20-levels"
    },
    {
      "_updated": "Thu, 30 Apr 2015 13:01:07 GMT",
      "description": "You",
      "_links": {
        "self": {
          "href": "achievements/11111111-1111-1111-1111-111111111111",
          "title": "achievement"
        }
      },
      "_created": "Thu, 30 Apr 2015 13:01:07 GMT",
      "_id": "11111111-1111-1111-1111-111111111111",
      "_etag": "c346f5e1c4f7658f2dfc4124efa87aba909a9821",
      "name": "buy-30-levels"
    }
  ],
  "_links": {
    "self": {
      "href": "achievements?max_results=2",
      "title": "achievements"
    },
    "last": {
      "href": "achievements?max_results=2&page=23",
      "title": "last page"
    },
    "parent": {
      "href": "/",
      "title": "home"
    },
    "next": {
      "href": "achievements?max_results=2&page=2",
      "title": "next page"
    }
  },
  "_meta": {
    "max_results": 2,
    "total": 46,
    "page": 1
  }
}
```

## Install

Using go

- `go get github.com/moul/anonuuid/...`

## Changelog

### [v1.1.0](https://github.com/moul/anonuuid/releases/tag/v1.1.0) (2018-04-02)

* Switch from `Godep` to `Glide`
* Add mutex to protect the cache field ([@QuentinPerez](https://github.com/QuentinPerez))
* Switch from `Party` to `Godep`
* Support of `--suffix=xxx`, `--keep-beginning` and `--keep-end` options ([#4](https://github.com/moul/anonuuid/issues/4))
* Using **party** to stabilize vendor package versions ([#8](https://github.com/moul/anonuuid/issues/8))
* Add homebrew package ([#6](https://github.com/moul/anonuuid/issues/6))

[full commits list](https://github.com/moul/anonuuid/compare/v1.0.0...master)

### [v1.0.0](https://github.com/moul/anonuuid/releases/tag/v1.0.0) (2015-10-07)

**Initial release**

#### Features

* Support of `--hexspeak` option
* Support of `--random` option
* Support of `--prefix` option
* Anonymize input stream
* Anonymize files

## License

© 2015-2019 [Manfred Touron](https://manfred.life) -
[Apache-2.0 License](https://github.com/moul/golang-repo-template/blob/master/LICENSE)
