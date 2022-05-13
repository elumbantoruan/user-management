FROM golang:1.18-alpine as build-env

RUN apk add git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/user-service ./cmd/app

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=build-env /app/user-service /
EXPOSE 8088
CMD ["/user-service"]