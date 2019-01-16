FROM golang

WORKDIR /go/src

COPY . .

RUN go get -d -v ./...
RUN go install -v ./cmd/...

CMD [ "tsuru-reviewapp-server" ]
