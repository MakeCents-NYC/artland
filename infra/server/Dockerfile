FROM golang:latest

WORKDIR /go/src/github.com/MakeCents-NYC/artland
COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."
