FROM golang:1.23 AS build
ENV GOFLAGS=-mod=vendor
ENV GOWORK=off
ENV CGO_ENABLED=0
WORKDIR /workspace
COPY . .
WORKDIR /workspace/upgrade-service
RUN go build -o /out/upgrade-service ./cmd/upgrade-service

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /out/upgrade-service /app/upgrade-service
COPY --from=build /workspace/upgrade-service/configs /app/configs
EXPOSE 8082
ENTRYPOINT ["/app/upgrade-service"]
