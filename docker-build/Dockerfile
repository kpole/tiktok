FROM golang:1.18-alpine3.17



RUN mkdir "/mw"
ENV PATH="/mw:${PATH}"
COPY /docker-build/ffmpeg-amd64-linux.tar.gz /mw/ffmpeg.tar.gz
RUN tar -zxvf /mw/ffmpeg.tar.gz -C /mw/
RUN rm /mw/ffmpeg.tar.gz

COPY biz/ /go/src/biz/
COPY pkg/ /go/src/pkg/
COPY go.mod /go/src/
COPY go.sum /go/src/
COPY router.go /go/src/
COPY router_gen.go /go/src/
COPY main.go /go/src/

WORKDIR "/go/src/"
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOOS=linux
RUN go env -w GOARCH=amd64
RUN go mod tidy
RUN go build -o offer_tiktok

RUN chmod +x /go/src/offer_tiktok

EXPOSE 18005
ENTRYPOINT ["./offer_tiktok"]

