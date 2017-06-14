# rdm - ReadMe scaffolding

[![CircleCI](https://img.shields.io/circleci/project/github/sbstjn/rdm.svg)](https://circleci.com/gh/sbstjn/rdm)
[![GitHub release](https://img.shields.io/github/release/sbstjn/rdm.svg)](https://github.com/sbstjn/rdm/releases)
[![license](https://img.shields.io/github/license/sbstjn/rdm.svg)](https://github.com/sbstjn/rdm/blob/master/LICENSE.md)

Command line tool to create default `LICENSE.md`, `CODE_OF_CONDUCT.md` and `README.md` files for your typical GitHub open source project.

## Features

- Create `LICENSE.md` ([MIT](https://opensource.org/licenses/MIT) or [unlicense](http://unlicense.org/))
- Create `CODE_OF_CONDUCT.md` (By [Contributor Covenant](http://contributor-covenant.org/version/1/4/)) 

```bash
$ > rdm

Scaffold common GitHub project files.

Usage:
  rdm [command]

Available Commands:
  conduct     Generate CODE_OF_CONDUCT.md file
  help        Help about any command
  license     Generate LICENSE.md file
  version     Show command version

Flags:
      --config string   config file (default is $HOME/.rdm.yaml)
  -f, --force           Overwrite files
  -o, --out string      output path (default is current directory)
```

## Install

```bash
$ > brew tap sbstjn/bin
$ > brew install rdm
$ > rdm version
0.0.3
```

## Usage

### Project License

Choose either to [unlicense your code](http://unlicense.org/), or use the default [MIT](https://opensource.org/licenses/MIT) license for your project.

```bash
$ > rdm license # -o destination/path, default is current directory
```

### Code of Conduct

Generate a `CODE_OF_CONDUCT.md` file based on [Contributor Covenant](http://contributor-covenant.org/version/1/4/).

```bash
$ > rdm conduct # -o destination/path, default is current directory
```

## Contribute

Run `make bindata` after you update a template. 🙆

### Todo

- [ ] Generate README.md
  - [ ] Support for different sections (shields.io etc.)
- [x] Generate LICENSE.md (MIT, Unlicense)
- [x] Generate CODE_OF_CONDUCT.md
- [ ] Presets for different languages (Node, Go)
- [ ] Use git config for default values
