FROM busybox:latest

LABEL maintainer="sebastian@devgru.com.pl"

COPY ./debug.sh debug.sh

RUN chmod +x debug.sh

ENTRYPOINT  ["/debug.sh"]