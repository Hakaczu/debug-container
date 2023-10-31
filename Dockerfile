FROM golang:1.21.3-alpine as build
WORKDIR /app
COPY go.mod server.go ./
RUN go build

FROM busybox:latest
LABEL maintainer="sebastian@devgru.com.pl"
COPY --from=build /app/server /server/
ENTRYPOINT  ["/server/server"]