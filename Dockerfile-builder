FROM debian:bullseye

RUN apt-get update &&\
  apt-get upgrade -y

COPY docker/entrypoint.sh /usr/local/bin/entrypoint.sh

RUN chmod u+x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
