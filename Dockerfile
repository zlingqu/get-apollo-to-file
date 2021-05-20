FROM alpine:3.12.1
WORKDIR /app
ADD build .
CMD ["/app/get-apollo-to-file","-h"]