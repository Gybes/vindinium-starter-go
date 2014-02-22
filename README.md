## Package Setup

Assuming you have the [suggested](http://golang.org/doc/code.html) directory layout simply:

    go get github.com/Gybes/vindinium-starter-go

Then copy the sample main.go file into your src directory:

    cp $GOPATH/src/github.com/Gybes/vindinium-starter-go/sample_main/main.go $GOPATH/src/<your_bot_repo>

And create a `secretKey` file with your bot's secret key in it:

    cd $GOPATH/src/<your_bot_repo>
    cat pbpaste > secretKey        # This will vary depending on your environment... hopefully you get the drift.

Examples:

    go run main.go arena 10
    go run main.go training 10
    go run main.go training 10 http://localhost:9000


## Writing your own bot

Interfaces are implicitly satisfied in Go so any type that implements the same method signature as that in the Bot interface can be used with this package.

[Here's](http://pastebin.com/GGcEVZek) an example of a bot implementation tacked onto the sample main file.
