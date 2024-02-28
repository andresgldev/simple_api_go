FROM golang:1.21-alpine AS build

WORKDIR /app


RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache

# COPY ./go.sum .
# COPY ./go.mod .
# COPY .env .
COPY ./. .
# RUN go mod download

RUN --mount=type=cache,target=/gomod-cache \
    go mod download

RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache \
    go build -o /myapp main.go

FROM alpine:3.19

COPY --from=build /myapp ./.
COPY .env /.env

EXPOSE 8080

CMD [ "/myapp" ]