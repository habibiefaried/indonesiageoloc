FROM alpine:latest
RUN addgroup -S appgroup && adduser -S appuser -G appgroup -h /app
COPY indonesiageoloc /app/indonesiageoloc
RUN chown -R appuser:appgroup /app && chmod +x /app/indonesiageoloc
USER appuser
# CMD /app/indonesiageoloc serve --DBSTRING "`echo $DBSTRING`"
