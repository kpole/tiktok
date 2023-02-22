FROM alpine:3.16

RUN mkdir "/app"
WORKDIR "/app"

RUN mkdir "/opt"
ENV PATH="/opt:${PATH}"

COPY ffmpeg-amd64-linux.tar.gz /opt/
RUN tar -zxvf /opt/ffmpeg-amd64-linux.tar.gz
COPY offer_tiktok /app/app
RUN chmod +x /app/app
ENTRYPOINT ["./app"]