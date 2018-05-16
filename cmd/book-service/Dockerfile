FROM alpine:3.7

RUN mkdir -p /app
WORKDIR /app

ADD book.json /app/book.json
ADD book-cli /app/book-cli

CMD ["./book-cli"]
