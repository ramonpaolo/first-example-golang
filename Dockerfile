FROM golang:1.21.6-alpine

WORKDIR /usr/src/app

COPY ./go.mod ./

RUN go mod download && go mod verify

COPY ./ ./

RUN go build

CMD [ "./simple-server" ]