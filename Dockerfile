FROM alpine:latest
RUN addgroup -S appgroup && adduser -S appuser -G appgroup -h /app
COPY web /app/web
RUN chown -R appuser:appgroup /app && chmod +x /app/web
USER appuser
WORKDIR /app
CMD ./web serve --DBSTRING "`echo $DBSTRING`"
