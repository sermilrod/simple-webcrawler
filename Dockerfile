FROM golang:1.11.5-alpine3.9 AS build
ENV APP_PATH /go/src/github.com/sermilrod/simple-webcrawler
RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH
ADD . $APP_PATH
RUN apk add --update git gcc make libc-dev
RUN rm -rf vendor simple-webcrawler
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build
RUN go test -v ./... --benchmem

FROM alpine:3.9
RUN apk add --no-cache ca-certificates
COPY --from=build /go/src/github.com/sermilrod/simple-webcrawler/simple-webcrawler .
ENTRYPOINT ["/simple-webcrawler"]
