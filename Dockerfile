FROM golang:1.21-alpine AS build

WORKDIR /app

# COPY ./go.sum .
# COPY ./go.mod .
# COPY .env .
COPY ./. .
# RUN go mod download

RUN go build -o /myapp main.go

FROM alpine:3.19

COPY --from=build /myapp ./.
COPY .env /.env

EXPOSE 8080

CMD [ "/myapp" ]