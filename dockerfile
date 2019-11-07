FROM golang:1-alpine

WORKDIR /go/src/app
COPY . .
RUN apk update && apk upgrade && apk add git
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o ./theapp -v .
RUN git rev-parse HEAD > commit_hash

FROM alpine:latest
WORKDIR /etc
RUN apk update && apk upgrade && apk add ca-certificates
COPY --from=0 /go/src/app/theapp /bin
COPY --from=0 /go/src/app/metadata /etc
COPY --from=0 /go/src/app/commit_hash /etc
EXPOSE 8001
CMD export HASH=$(cat /etc/commit_hash); theapp