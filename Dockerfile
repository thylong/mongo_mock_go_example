FROM golang:1.7

ADD . /go/src/github.com/thylong/mongo_mock_go_example
WORKDIR /go/src/github.com/thylong/mongo_mock_go_example

# Fetching Dependencies.
RUN go get github.com/tools/godep
RUN godep restore

RUN go install github.com/thylong/mongo_mock_go_example

EXPOSE 8080

CMD ["/go/bin/mongo_mock_go_example"]
