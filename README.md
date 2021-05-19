# licenser

A simple command line tool to add a license to a project and assign copyright to Github contributors. 

## Installation

``` shell
go get github.com/nandanrao/licenser
```

## Usage

The command: 

``` shell
licenser > LICENSE
```

Results in an MIT license with the following, where the Github URL is created from the git remote information (defaults to "origin" remote): 

``` text
Copyright (c) 2021 licenser contributors
(https://github.com/nandanrao/licenser/graph/contributors)
```

You can add additional copyright holders with positional arguments:

``` shell
licenser "Nandan Rao"> LICENSE
```

Which results in an MIT license with the following: 

``` text

Copyright (c) 2021 Nandan Rao and licenser contributors
(https://github.com/nandanrao/licenser/graph/contributors)
```

There is also a "year" flag:

``` shell
licenser -y 2019-2020 > LICENSE
```

You can pick the git remote name:


``` shell
licenser -r not-the-origin > LICENSE
```

