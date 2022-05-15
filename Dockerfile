FROM golang:1.16-alpine as build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./graph ./graph
COPY ./internal ./internal

COPY *.go ./

RUN go build -o ./appBuilt

CMD [ "./appBuilt" ]