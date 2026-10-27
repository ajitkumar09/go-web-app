FROM golang:1.22.5 AS build

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN mkdir -p /out && CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/go-web-app .

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=build /out/go-web-app ./go-web-app
COPY --from=build /src/static ./static

EXPOSE 8080

ENTRYPOINT ["/app/go-web-app"]
