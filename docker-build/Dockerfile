FROM alpine:3.16

RUN mkdir "/app"
WORKDIR "/app"

RUN mkdir "/mw"
ENV PATH="/mw:${PATH}"

COPY ffmpeg-amd64-linux.tar.gz /mw/ffmpeg.tar.gz
RUN tar -zxvf /mw/ffmpeg.tar.gz -C /mw/
RUN rm /mw/ffmpeg.tar.gz
COPY offer_tiktok /app/app
RUN chmod +x /app/app
ENTRYPOINT ["./app"]
