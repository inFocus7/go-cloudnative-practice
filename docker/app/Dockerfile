# Build environment
FROM golang:1.15-alpine as build-env
WORKDIR /go-cloudnative-practice

RUN apk update && apk add --no-cache gcc musl-dev git

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

# Deployment environment
FROM alpine
RUN apk update && apk add --no-cache bash

COPY --from=build-env /go-cloudnative-practice/bin/app /go-cloudnative-practice/

EXPOSE 8080
CMD ["/go-cloudnative-practice/app"]