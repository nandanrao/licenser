# licenser

A simple command line tool to add a license to a project and assign copyright to Github contributors. 

## Installation

``` shell
go get github.com/nandanrao/licenser
```

## Usage

The command: 

``` shell
licenser -p nandanrao/licenser > LICENSE
```

Results in an MIT license with the following: 

``` text
Copyright (c) 2021 licenser contributors
(https://github.com/nandanrao/licenser/graph/contributors)
```

While the command: 

``` shell
licenser -p nandanrao/licenser "Nandan Rao"> LICENSE
```

Results in an MIT license with the following: 

``` text

Copyright (c) 2021 Nandan Rao and licenser contributors
(https://github.com/nandanrao/licenser/graph/contributors)
```

There is also a "year" flag:

``` shell
licenser -y 2019-2020 -p nandanrao/licenser > LICENSE
```
