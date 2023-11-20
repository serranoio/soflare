# builder
FROM golang:latest AS builder
RUN apt-get update && apt-get install -y gcc sqlite3 libsqlite3-dev
WORKDIR /app
COPY . .
RUN go build -o main .

# Fly.ctl
# curl -L https://fly.io/install.sh | sh
# download git
# sudo apt install git


EXPOSE 8080

# start PocketBase
CMD [ "./main" ]

