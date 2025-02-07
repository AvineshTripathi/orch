# Build stage for API server
FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build  -a -installsuffix cgo -o build/orch .

FROM alpine:3.21.0
WORKDIR /app
COPY --from=build /app/build/orch /app/orch
EXPOSE 8089
CMD ["./orch"]