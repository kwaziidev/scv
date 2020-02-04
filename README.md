# scv
Generate scaffold project layout for Go.

[https://github.com/golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## Installation
Download scv by using:

```
$ go get -u github.com/gloomyzerg/scv
```

## Create a new project

```
$ scv [projectName]
```

## Examples

```
$ scv example
? Choose a mode:  [Use arrows to move, type to filter]
  min
  std 
> full

? Choose a mode: full
generate success!

$ tree
.
└── example
    ├── api
    ├── assets
    ├── build
    │   ├── ci
    │   └── package
    ├── cmd
    ├── configs
    ├── deployments
    ├── docs
    ├── examples
    ├── githooks
    ├── init
    ├── internal
    │   ├── app
    │   └── pkg
    ├── pkg
    ├── scripts
    ├── test
    ├── third_party
    ├── tools
    ├── vendor
    ├── web
    │   ├── app
    │   ├── static
    │   └── template
    └── website

27 directories, 0 files
```