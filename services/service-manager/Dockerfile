FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN go build -o /out/service-manager ./cmd/service-manager

FROM gcr.io/distroless/base-debian12
COPY --from=build /out/service-manager /service-manager
EXPOSE 8080
ENTRYPOINT ["/service-manager"]
