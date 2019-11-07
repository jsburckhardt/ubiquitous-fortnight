FROM golang:1-alpine

WORKDIR /go/src/app
COPY . .
RUN apk update && apk upgrade && apk add git
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o ./app -v .

FROM alpine:latest
RUN apk update && apk upgrade && apk add ca-certificates
COPY --from=0 /go/src/app/app /bin
EXPOSE 8001
CMD ["app"]