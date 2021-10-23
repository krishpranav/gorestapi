## Building gorestapi

### make sure you have the latest version of golang

## Building:
```
ruby setup.rb
```

### testing the api:

## Main page
- main page you will get response ```Hello World```
```
$ curl http://localhost:8080/
```


## Get response about the albums 
- get response of the albums that are stored locally
```
$ curl http://localhost:8080/albums
```

## About restapi response
- you will get to know about the restapi version, author by curling this request
```
$ curl http://localhost:8080/about
```

### running tests:
```
$ go test
```

### Building Frontend:
```bash
$ cd views/
$ yarn build
$ cd ../
$ go run main.go
```
- go to [localhost](http://localhost:8080) and in the console you can find the api request 