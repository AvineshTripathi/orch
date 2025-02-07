# Build stage for API server
FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build  -a -installsuffix cgo -o build/orch .
RUN cd provisioner && CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o ../build/provisioner .

FROM alpine:3.21.0
WORKDIR /app
COPY --from=build /app/taskConfig.yaml /app/tasks.yaml
COPY --from=build /app/provisioner/taskPlugins/ /app/taskPlugins
COPY --from=build /app/build/provisioner /app/provisioner
EXPOSE 50051
CMD ["./provisioner"]
