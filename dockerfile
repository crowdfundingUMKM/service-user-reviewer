#syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /service-Ureviewer

EXPOSE 8080

CMD [ "/service-Ureviewer" ]