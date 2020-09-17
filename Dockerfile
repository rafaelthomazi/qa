FROM golang:latest AS build

WORKDIR /go/src/github.com/rafaelthomazi/qa

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN cd server && \
    CGO_ENABLED=0 go build -a -ldflags="-s -w" -installsuffix cgo -o /go/bin/qa-server .

# Runtime image
FROM alpine:latest AS runtime

RUN apk --no-cache add ca-certificates

COPY --from=build /go/bin/qa-server /usr/local/bin/qa-server

CMD [ "/usr/local/bin/qa-server" ]