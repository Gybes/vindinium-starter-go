Assuming you have the [suggested](http://golang.org/doc/code.html) directory layout simply:

    go get github.com/Gybes/vindinium-starter-go

Then copy the sample main.go file into your src directory:

    cp $GOPATH/src/github.com/Gybes/vindinium-starter-go/sample_main/main.go $GOPATH/src/

Examples:

    go run main.go mySecretKey arena 10
    go run main.go mySecretKey training 10
    go run main.go mySecretKey training 10 http://localhost:9000
