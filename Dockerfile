FROM golang:1.12.1-stretch
RUN mkdir /go/src/expvar-test
WORKDIR /go/src/expvar-test
ADD . .
#RUN go get -d -v ./...
#RUN go install -v ./...
EXPOSE 8001
EXPOSE 8002
EXPOSE 8003
ENTRYPOINT /go/src/expvar-test/start.sh
