FROM scratch
MAINTAINER Lei Cao "lei.cao.life@gmail.com"
EXPOSE 8085
WORKDIR /app
# copy binary into image
COPY beefer /app/
# copy the views into image
COPY views /app/views/
# copy the static files into image
COPY static /app/static/
ENTRYPOINT ["./beefer"]
