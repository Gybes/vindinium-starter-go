Assuming you have the [suggested](http://golang.org/doc/code.html) directory layout simply:

    go get github.com/Gybes/vindinium-starter-go

Then copy the sample main.go file into your src directory:

    cp $GOPATH/src/github.com/Gybes/vindinium-starter-go/sample_bot/main.go $GOPATH/src/<your_bot_repo>

Put your secret key in a file called `secretKey` in the same dir as you put main.go:

    cd $GOPATH/src/<your_bot_repo>
    cat pbpaste > secretKey        # this will be different depending on your environment... hopefully you get the drift.

Run your bot!

    go run main.go arena 10
    go run main.go training 10
    go run main.go training 10 http://localhost:9000
