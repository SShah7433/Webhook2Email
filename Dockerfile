FROM alpine:latest
WORKDIR /app
COPY dist/webhook2email /app
CMD ["./webhook2email"]