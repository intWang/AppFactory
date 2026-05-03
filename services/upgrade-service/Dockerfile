FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN go build -o /out/upgrade-service ./cmd/upgrade-service

FROM gcr.io/distroless/base-debian12
COPY --from=build /out/upgrade-service /upgrade-service
EXPOSE 8082
ENTRYPOINT ["/upgrade-service"]
