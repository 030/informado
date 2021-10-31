# Informado

[![GoDoc Widget]][GoDoc]
[![Build Status](https://travis-ci.org/030/informado.svg?branch=master)](https://travis-ci.org/030/informado)
[![Go Report Card](https://goreportcard.com/badge/github.com/030/informado)](https://goreportcard.com/report/github.com/030/informado)
[![StackOverflow SE Questions](https://img.shields.io/stackexchange/stackoverflow/t/informado.svg?logo=stackoverflow)](https://stackoverflow.com/tags/informado)
[![DevOps SE Questions](https://img.shields.io/stackexchange/devops/t/informado.svg?logo=stackexchange)](https://devops.stackexchange.com/tags/informado)
[![ServerFault SE Questions](https://img.shields.io/stackexchange/serverfault/t/informado.svg?logo=serverfault)](https://serverfault.com/tags/informado)
![Issues](https://img.shields.io/github/issues-raw/030/informado.svg)
![Pull requests](https://img.shields.io/github/issues-pr-raw/030/informado.svg)
![Total downloads](https://img.shields.io/github/downloads/030/informado/total.svg)
![License](https://img.shields.io/github/license/030/informado.svg)
![Repository Size](https://img.shields.io/github/repo-size/030/informado.svg)
![Contributors](https://img.shields.io/github/contributors/030/informado.svg)
![Commit activity](https://img.shields.io/github/commit-activity/m/030/informado.svg)
![Last commit](https://img.shields.io/github/last-commit/030/informado.svg)
![Release date](https://img.shields.io/github/release-date/030/informado.svg)
![Latest Production Release Version](https://img.shields.io/github/release/030/informado.svg)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=030_informado&metric=bugs)](https://sonarcloud.io/dashboard?id=030_informado)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=030_informado&metric=code_smells)](https://sonarcloud.io/dashboard?id=030_informado)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=030_informado&metric=coverage)](https://sonarcloud.io/dashboard?id=030_informado)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=030_informado&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=030_informado)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=030_informado&metric=ncloc)](https://sonarcloud.io/dashboard?id=030_informado)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_informado&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=030_informado)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=030_informado&metric=alert_status)](https://sonarcloud.io/dashboard?id=030_informado)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=030_informado&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=030_informado)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=030_informado&metric=security_rating)](https://sonarcloud.io/dashboard?id=030_informado)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=030_informado&metric=sqale_index)](https://sonarcloud.io/dashboard?id=030_informado)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=030_informado&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=030_informado)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/2810/badge)](https://bestpractices.coreinfrastructure.org/projects/2810)
[![codecov](https://codecov.io/gh/030/informado/branch/master/graph/badge.svg)](https://codecov.io/gh/030/informado)
[![BCH compliance](https://bettercodehub.com/edge/badge/030/informado?branch=master)](https://bettercodehub.com/results/030/informado)
[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-web.svg)](https://golangci.com/r/github.com/030/informado)
[![informado](https://snapcraft.io//informado/badge.svg)](https://snapcraft.io/informado)
[![codebeat badge](https://codebeat.co/badges/60706232-493c-4527-b0c9-9e38f682b68c)](https://codebeat.co/projects/github-com-030-informado-master)

<a href="https://informado.releasesoftwaremoreoften.com">\
<img src="https://github.com/030/informado/raw/master/assets/logo/logo.png" height="100"></a>

Use this Go library or the tool to read various RSS feeds. Note that Atom and
Reddit feeds can be parsed as well.

## Installation

### Ubuntu

```bash
sudo snap install informado
```

## Usage

Create an `informado` directory:

```bash
mkdir ~/.informado
```

and subsequently an `rss-feed-urls.csv` file:

```bash
type,url
atom,https://github.com/golang/go/releases.atom
```

Once the file has been created, run:

```bash
./informado
```

Once informado has been completed, a `~/.informado/last-run-time.txt` has been
created that contains the Epoch time when the tool was run. The next time
informado is run it will lookup the time and only show newer messages. If one
would like to view all messages, then the time has to be changed in the
`.informado` file.

```bash
docker run \
  -v /home/${USER}/.informado:/home/informado/.informado \
  -it informado:3.0.0
```

### Slack

Create a `~/.informado/creds.yml` file:

```bash
---
slackChannel: x
slackToken: y
```

## Stargazers over time

[![Stargazers over time](https://starchart.cc/030/informado.svg)](https://starchart.cc/030/informado)

[GoDoc]: https://godoc.org/github.com/030/informado
[GoDoc Widget]: https://godoc.org/github.com/030/informado?status.svg
