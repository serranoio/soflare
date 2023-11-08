
# builder
FROM golang:latest AS builder
RUN apt-get update && apt-get install -y gcc sqlite3 libsqlite3-dev
WORKDIR /app
COPY . .
RUN go build -o main .

EXPOSE 8080

# start PocketBase
CMD [ "./main", "serve", "--http=0.0.0.0:8080" ]