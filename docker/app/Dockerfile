# Build environment
FROM golang:1.15-alpine as build-env
WORKDIR /go-cloudnative-practice

RUN apk update && apk add --no-cache gcc musl-dev git

# Cache modules by placing them in Docker, downlading them, then copying all other files to Docker.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate

# Deployment environment
FROM alpine
RUN apk update && apk add --no-cache bash mysql-client

COPY --from=build-env /go-cloudnative-practice/bin/app /go-cloudnative-practice/
COPY --from=build-env /go-cloudnative-practice/bin/migrate /go-cloudnative-practice/
COPY --from=build-env /go-cloudnative-practice/migrations /go-cloudnative-practice/migrations

COPY --from=build-env /go-cloudnative-practice/docker/app/bin /usr/local/bin/go-cloudnative-practice/
RUN chmod +x /usr/local/bin/go-cloudnative-practice/*
# <ADDED> Manual execusion permissions to /migrate, since originally unable to run it.
RUN chmod +x /go-cloudnative-practice/migrate

EXPOSE 8080
CMD ["/go-cloudnative-practice/app"]