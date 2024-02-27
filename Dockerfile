FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
#COPY ./go.mod .
COPY ./. ./

RUN go build -o /myapp main.go

FROM alpine:3.19

COPY --from=build /myapp ./

EXPOSE 8080

CMD [ "/myapp" ]