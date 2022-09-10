FROM golang:1.19.1-bullseye as build
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/bin/is-mercury-in-retrograde


FROM gcr.io/distroless/static
COPY --from=build /usr/bin/is-mercury-in-retrograde .
ENTRYPOINT ["/is-mercury-in-retrograde"]
