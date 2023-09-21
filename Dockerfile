FROM golang:1.19.13-alpine3.18 AS Builder
WORKDIR /go/src
COPY . .
RUN go build -o auth-service

FROM alpine
WORKDIR /service
COPY --from=Builder /go/src/auth-service .
RUN touch .env
ENTRYPOINT ["./auth-service"]