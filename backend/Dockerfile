FROM golang:latest

ADD . /dockerdev
WORKDIR /dockerdev

RUN go build ./cmd/server -o /server

#Final stage
FROM debian:buster

EXPOSE 8081

WORKDIR /
COPY --from=build-env /server /

CMD ["/server"]