FROM saaskit/alpine:healthcheck
ENV APP_PORT=8000
WORKDIR /
COPY nats/example/pub/.build/app /
EXPOSE $APP_PORT
CMD ["/app"]
