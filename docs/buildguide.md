## Building gorestapi

### make sure you have the latest version of golang

### macOS:
```
$ git clone https://github.com/krishpranav/gorestapi.git
$ cd gorestapi
$ go mod tidy 
$ go run main.go
```

### linux:
```
$ git clone https://github.com/krishpranav/gorestapi.git
$ cd gorestapi
$ go mod tidy 
$ go run main.go
```

### running tests:
```
$ go test
```

### testing the api:

- curl http://localhost:8080/
- main page you will get response ```Hello World```


- curl http://localhost:8080/albums
- get response of the albums that are stored