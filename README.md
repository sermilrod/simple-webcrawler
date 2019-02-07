# Simple Webcrawler
A simple webcrawler written in go

# Building the artefact
```
$ docker build -t sermilrod/simple-webcrawler .
$ docker push sermilrod/simple-webcrawler
```

# How to use it
```
$ docker run --rm sermilrod/simple-webcrawler -url http://foo.bar
$ docker run --rm sermilrod/simple-webcrawler -url http://foo.bar -url http://bibidi.babidi.boo
```

# Running the tests
```
$ go test -v ./... --benchmem
```
