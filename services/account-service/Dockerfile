FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN go build -o /out/account-service ./cmd/account-service

FROM gcr.io/distroless/base-debian12
COPY --from=build /out/account-service /account-service
EXPOSE 8081
ENTRYPOINT ["/account-service"]
