FROM golang:1.21.5-alpine3.19 as builder
WORKDIR /api
COPY . .
RUN go mod tidy
RUN go build -o main cmd/main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /api/main .
EXPOSE 1323
CMD [ "./main" ]