#First build binary
FROM golang:1.14 AS build
WORKDIR /go/src/learn-docker
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

#Copy only binary to final image
FROM alpine
WORKDIR /app

#Copy binary
COPY --from=build /go/src/learn-docker/app .
EXPOSE 3000

CMD ["./app"]