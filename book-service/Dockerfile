FROM alpine:3.7
RUN apk add --update --no-cache openssh ca-certificates zlib
RUN mkdir /app
WORKDIR /app
COPY book-service /app
ENTRYPOINT ["./book-service"]