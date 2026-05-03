FROM golang:1.23 AS build
ENV GOFLAGS=-mod=vendor
ENV GOWORK=off
ENV CGO_ENABLED=0
WORKDIR /workspace
COPY . .
WORKDIR /workspace/service-manager
RUN go build -o /out/service-manager ./cmd/service-manager

FROM docker:cli
WORKDIR /app
COPY --from=build /out/service-manager /app/service-manager
COPY --from=build /workspace/service-manager/configs /app/configs
EXPOSE 8080
ENTRYPOINT ["/app/service-manager"]
