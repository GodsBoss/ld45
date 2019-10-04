FROM golang:1.12.10
RUN mkdir -p /go/src/github.com/GodsBoss
RUN mkdir /go/.cache
RUN chown -R 1000:1000 /go
USER 1000:1000
ENV GOCACHE=/go/.cache
WORKDIR /go/src/github.com/GodsBoss
RUN go get -u -v github.com/gopherjs/gopherjs
