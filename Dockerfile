FROM golang:alpine AS build_base

WORKDIR /tmp/go_app

COPY . .

RUN go mod download
RUN go build -o ./sample_app .

FROM alpine:latest

COPY --from=build_base /tmp/go_app/sample_app /app/sample_app

EXPOSE 9080

CMD ["/app/sample_app"]