FROM golang:1.16.7 as Builder

ADD src /app

WORKDIR /app

RUN go build -o app .

FROM alpine:3.14.1

COPY --from=Builder /app/app app

CMD ["./app"]
