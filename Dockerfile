FROM golang:1.17

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN go build -o main ./recv.go

CMD /app/main