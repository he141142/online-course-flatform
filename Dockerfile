FROM golang:1.19.6-alpine AS build-env

RUN apk --no-cache add git


RUN apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /godocker -ldflags="-s -w" cmd/main.go


FROM alpine:3.18

COPY --from=build-env /godocker /godocker
COPY --from=build-env /app/dev.env dev.env
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV PORT=8080

EXPOSE $PORT
ENTRYPOINT ["/godocker", "dev"]