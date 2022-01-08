# GOROOT is for compiler/tools that comes from go installation.
# linux
$ export PATH=$PATH:$GOROOT/bin
# window
$ export PATH=$PATH:$GOROOT\\bin

# GOPATH is for your own go projects / 3rd party libraries (downloaded with "go get").
# linux
$ export GOBIN=$GOPATH/bin
# window
$ export GOBIN=$GOPATH\\bin

# check Environment Variable for go
$ go env