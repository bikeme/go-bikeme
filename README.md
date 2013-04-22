# Usage

Create a Go workspace, for example:

    mkdir -p ~/Documents/mygo

Set the GOPATH environment variable to this directory:

    export GOPATH=~/Documents/mygo

Create a `src/` directory:

    mkdir ~/Documents/mygo/src

Clone the repo inside the `src/` directory:

    git clone git@github.com:damselem/go-bikeme.git

Compile and run the application:

    cd ~/Documents/mygo/src/go-bicing
    go run main.go
