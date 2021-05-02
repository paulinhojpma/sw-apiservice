FROM golang:latest
ENV GO111MODULE=on

ADD . /sw-sys/api-service
WORKDIR /sw-sys/api-service/
#RUN ls
COPY go.mod .
COPY go.sum .
#RUN go mod download
#ENV GO111MODULE=off
#WORKDIR /go/src/github.com/sw-sys/escola-service/cmd/escola-service
RUN go build

CMD ["./api-service"]
