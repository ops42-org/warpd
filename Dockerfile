FROM debian:bullseye-slim

RUN apt-get update &&\
  apt-get upgrade -y

COPY docker/entrypoint.sh /usr/local/bin/entrypoint.sh
COPY build/warpd-linux-amd64 /usr/local/bin/warpd

RUN chmod u+x /usr/local/bin/entrypoint.sh /usr/local/bin/warpd

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
