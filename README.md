# interpreter <!-- omit from toc -->

- [Overview](#overview)
- [Dev Setup](#dev-setup)

## Overview

This repo is my implementation of an interpreter created while reading Thorsten Ball's "Writing an Interpreter in Go".

## Dev Setup

This repo uses [magefiles](https://magefile.org/), so you will need the mage exe that is at least v1.13 in your PATH. If you make a temporary folder, you can install mage via:

```bash
git clone https://github.com/magefile/mage
cd mage
go run bootstrap.go
```

The bootstrap is necessary to end up with an exe that includes version information.
