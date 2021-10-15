## Building gorestapi

### make sure you have the latest version of golang

## Building:
```
ruby setup.rb
```

### running tests:
```
$ go test
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