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

    cd ~/Documents/mygo/src/go-bikeme
    go run main.go

## Contributing

Please use the issues page to report any bug or suggest new features. If you feel brave:

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Run the tests and make sure they all pass (`go test  ./...`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request
