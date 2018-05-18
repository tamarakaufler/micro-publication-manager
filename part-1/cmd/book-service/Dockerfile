FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD book.json /app/book.json
ADD book-cli /app/book-cli

CMD ["./book-cli"]
