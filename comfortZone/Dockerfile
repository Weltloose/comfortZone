FROM golang

COPY comfortZone/ $GOPATH/src/github.com/Weltloose/comfortZone/

ENV GOPROXY=https://goproxy.cn

WORKDIR $GOPATH/src/github.com/Weltloose/comfortZone

RUN go install

WORKDIR $GOPATH/bin

EXPOSE 80

CMD ["./comfortZone"]
