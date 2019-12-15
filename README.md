# daddy

[![Build Status](https://travis-ci.com/artberri/daddy.svg?branch=master)](https://travis-ci.com/artberri/daddy)
[![Go Report Card](https://goreportcard.com/badge/artberri/daddy)](https://goreportcard.com/report/artberri/daddy)
[![Coverage Status](https://coveralls.io/repos/github/artberri/daddy/badge.svg?branch=master)](https://coveralls.io/github/artberri/daddy?branch=master)

`daddy` is a command line interface to manage DNS records in GoDaddy.

The idea for creating this tool is to be able to create ACME challenges in automatic tools.

## Downloads

[In this link](https://github.com/artberri/daddy/releases) you will find the packages for every supported platform. Please download the proper package for your operating system and architecture. You can also download older versions.

## Installation

To install `daddy`, find the [appropriate package](https://github.com/artberri/daddy/releases)
for your system and download it.

After downloading `daddy`, unzip the package. `daddy` runs as a single binary. Any other files in the package can be safely removed and `daddy` will still function.

The final step is to make sure that the `daddy` binary is available on the PATH. See [this page](https://stackoverflow.com/questions/14637979/how-to-permanently-set-path-on-linux) for instructions on setting the PATH on Linux and Mac. [This page](https://stackoverflow.com/questions/1618280/where-can-i-set-path-to-make-exe-on-windows) contains instructions for setting the PATH on Windows.

## Usage

Run the help command and follow the instructions.

```bash
daddy --help
```

You need to setup your API key and secret to obtain/set data. For example:

```bash
daddy list --key=1234567689 --secret=1234566
```

If you do not want to pass the parameters on each command you can create
the file `$HOME/.daddy.yaml` and put your configuration there. For example:

```taml
---
key: 1234567689
secret: 1234567689
```

### Available commands

| Command | Description            |
|---------|------------------------|
| help    | Help about any command |
| list    | List owned domains     |
| show    | Show DNS records       |
| add     | Add DNS record         |
| update  | Update DNS records     |
| remove  | Remove DNS records     |

## License

daddy. CLI to manage DNS records in GoDaddy.

Copyright © 2019 Alberto Varela <alberto@berriart.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
