FROM golang:1.23 AS build
ENV GOFLAGS=-mod=vendor
ENV GOWORK=off
ENV CGO_ENABLED=0
WORKDIR /workspace
COPY . .
WORKDIR /workspace/account-service
RUN go build -o /out/account-service ./cmd/account-service

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /out/account-service /app/account-service
COPY --from=build /workspace/account-service/configs /app/configs
EXPOSE 8081
ENTRYPOINT ["/app/account-service"]
