# Specifies a parent image
FROM golang:1.23
LABEL authors="Hessam"

# Creates an app directory to hold your app’s source code
WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

# Copies everything from your root directory into /app
COPY . .

# production -------------
#RUN go build -v -o /usr/local/bin/app

#CMD ["/usr/local/bin/app"]
# -----------------------

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o tmp/main main.go" --command=./tmp/main
