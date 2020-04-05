FROM alpine:latest
RUN addgroup -S appgroup && adduser -S appuser -G appgroup -h /app
USER appuser
WORKDIR /app
COPY web /app/web
RUN chmod +x web
CMD ./web serve --DBSTRING "`echo $DBSTRING`"
