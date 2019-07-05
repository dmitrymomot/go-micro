FROM saaskit/alpine:healthcheck
ENV APP_PORT=8000
WORKDIR /
COPY test/.build/app /
EXPOSE $APP_PORT
CMD ["/app"]
