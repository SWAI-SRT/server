FROM alpine:latest

WORKDIR /app

COPY swai .
RUN chmod +x swai

CMD ["./swai"]