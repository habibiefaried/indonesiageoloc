FROM ubuntu:latest
USER root
WORKDIR /root
COPY indonesiageoloc /root/indonesiageoloc
RUN chmod +x indonesiageoloc
CMD ./indonesiageoloc serve --DBSTRING "`echo $DBSTRING`"
