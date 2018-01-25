# exmaple api with golang

## Install go with package installer

[download](https://golang.org/dl/) |
[insatll](https://golang.org/doc/install?download)

## Set PATH in .bashrc or .zshrc

```sh
export GOROOT=/usr/local/go or export GOROOT=~/bin/go # manual install
export GOPATH=$HOME/go
export LGOBIN=$GOPATH/bin
export GOPROJECT=$GOPATH/src
export PATH=$PATH:$LGOBIN:$GOPATH:$GOROOT/bin
export GO15VENDOREXPERIMENT=1
```

## Install dependencies management

docs from [https://github.com/Masterminds/glide](https://github.com/Masterminds/glide)

```bash
brew install glide
# or
curl https://glide.sh/get | sh
# or
sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update
sudo apt-get install glide
```

## install dependencies of project

```sh
cd $GOPROJECT/myproject
glide install
```

## Run

```go
go run *.go
```

gomon from [https://www.npmjs.com/package/go-mon](https://github.com/johannesboyne/gomon)

```npm
 npm install -g go-mon
```

usage

```go
gomon *.go
```
