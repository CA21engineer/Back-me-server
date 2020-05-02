FROM golang:1.14-alpine as build

WORKDIR /app

COPY . .

RUN apk add --no-cache git \
	&& go get github.com/markbates/refresh \
	&& go build -o ./cmd/ca-zoooom/main

FROM alpine

WORKDIR /app

COPY --from=build /app/main .

RUN addgroup go \
	&& adduser -D -G go go \
	&& chown -R go:go /app/main

CMD ["./main"]