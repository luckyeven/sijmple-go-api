# Step 1: build binary
FROM docker.io/library/alpine:3.17.2 AS build
RUN apk update && apk upgrade && apk add --no-cache go gcc g++
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=1 GOOS=linux go build

# Step 2: deployment image
FROM docker.io/library/alpine:3.17.2
WORKDIR /app
COPY --from=build /app/simple-go-api /app/simple-go-api
EXPOSE 8080
USER 1001
ENTRYPOINT ["/app/simple-go-api"]
