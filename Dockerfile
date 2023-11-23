FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates && apt-get install -y tzdata && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY main /var/clash_request/clrequest
WORKDIR /var/clash_request/
VOLUME [ "/var/clash_request/config.json" ]
VOLUME [ "/var/clash_request/data" ]
ENTRYPOINT [ "./clrequest" ]
CMD [ "start" ]