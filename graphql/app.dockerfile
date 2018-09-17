FROM golang:1.10.2-alpine3.7 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/cipepser/graphql-microservice-sample/graphql
COPY vendor ../vendor
COPY account ../account
COPY catalog ../catalog
COPY order ../order
COPY grapphql ./
RUN go build -o /go/bin/grapphql

FROM alpine:3.7
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["app"]