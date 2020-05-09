FROM golang:1.11.5

ENV APP_NAME go_demo
ENV PORT 8080

COPY main.go /go/src/${APP_NAME}/main.go
WORKDIR /go/src/${APP_NAME}

RUN go get ./
RUN go build -o ${APP_NAME} main.go

CMD ./${APP_NAME}

EXPOSE ${PORT}