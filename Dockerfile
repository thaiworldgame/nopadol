FROM scratch
ENV GOROOT /usr/local
## docker set env APP_MODE = PRODUCTION
ENV APP_MODE production
# docker set env to +7:00 time zone
ENV TZ=Asia/Bangkok

#ADD https://golang.org/lib/time/zoneinfo.zip /usr/local/lib/time/
#ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/ca-certificates.crt
#ADD secret/mch_privkey.pem /secret/mch_privkey.pem
#ADD secret/service_account.json /secret/service_account.json

#COPY app /
#COPY admin/html /admin/html
#COPY admin/assets /admin/assets

## Add config file with Global variable such as : time server
#ADD config.yml /

EXPOSE 9999

ENTRYPOINT ["/app"]