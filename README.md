# ◝(^⌣^)◜ <!-- omit from toc -->

- [Overview](#overview)
- [Dev Setup](#dev-setup)

## Overview

This repo is my implementation of an interpreter for a made up language I call "Hai".

It is being created as I work my way through [Thorsten Ball's "Writing an Interpreter in Go"](https://interpreterbook.com/).

## Dev Setup

Sync this repo in the usual ways, e.g.:

```bash
git clone git@github.com:danbrakeley/hai.git
```

This repo uses a [magefile](https://magefile.org/), so you will need `mage` (>= v1.13) in your path.

To install the latest mage, just make a temporary folder, and do the following in it:

```bash
git clone https://github.com/magefile/mage
cd mage
go run bootstrap.go
```

To test, just run `mage` with no arguments in the root of your local copy of this repo. It should look like this:

```text
$ mage
Targets:
  build       tests and builds all apps
  buildHai    builds cmd/hai (output goes to "local" folder)
  ci          runs all CI tasks
  gen         runs go generate for all packages
  runWaiig    runs unit tests, builds hai until /local, then runs it
  setup       installs cli apps needed for development (not including 'go' or 'mage')
  test        runs tests for all packages
```
