FROM ubuntu:14.04
MAINTAINER Mickey Yawn <mickey.yawn@turner.com>

### update Ubuntu
RUN apt-get update

### set timezone
ENV TZ=America/New_York
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

### add binary to image
ADD cloud-go-ref /var/www/cloud-go-ref/

#### run binary
CMD ["/var/www/cloud-go-ref/cloud-go-ref"]

EXPOSE 80
