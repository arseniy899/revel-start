# GOlang
* [Get it] (https://golang.org/dl)
* [To get started with syntax] (https://www.tutorialspoint.com/go/index.htm)
## Writing down sample HTTP server
Simply put that file `http.go` anywhere:
```
package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":81", nil)
}
```
Start it up: `go run http-servers.go`
# More complex 
[Compairing the most top web frameworks](https://nordicapis.com/7-frameworks-to-build-a-rest-api-in-go/)
## Revel

### To get sources and run
```
go get github.com/arseniy899/revel-start
revel run revel-start
```

### Go to http://localhost:81/ and you'll see:

    "It works"

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites

## Help

* The [Golang refrence](https://www.tutorialspoint.com/go/index.htm).
* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).
* The [Bootstrap sample app](https://github.com/richtr/baseapp).

