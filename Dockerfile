FROM golang:1.12.0-alpine3.9
RUN mkdir /app apk update && apk add git
ADD . /app
WORKDIR /app
## Add this go mod download command to pull in any dependencies
RUN go mod download
## Our project will now successfully build with the necessary go libraries included.
RUN go build ./documentService/internal/interface/main .
## Our start command which kicks off
## our newly created binary executable
EXPOSE 8081
CMD ["/main"]