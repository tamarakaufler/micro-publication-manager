FROM alpine:3.7

RUN mkdir -p /app
WORKDIR /app

ADD author.json /app/author.json
ADD author-cli /app/author-cli

CMD ["./author-cli"]
