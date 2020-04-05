FROM ubuntu:latest
USER root
WORKDIR /root
COPY web /root/web
RUN chmod +x web
CMD ./web serve --DBSTRING "`echo $DBSTRING`"
