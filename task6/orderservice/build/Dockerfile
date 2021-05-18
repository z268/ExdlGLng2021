FROM golang:1.16

WORKDIR /go/src/github.com/z268/ExdlGLng2021/task6/orderservice

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
RUN go get -u github.com/swaggo/swag/cmd/swag

COPY ./. .
RUN swag init -g cmd/main.go

RUN go build cmd/main.go

COPY ./build/*.sh .
RUN chmod +x ./*.sh

CMD ["./start_http.sh"]