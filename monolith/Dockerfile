FROM ubuntu:14.04

EXPOSE 8065
VOLUME /mattermost/data

RUN apt-get update && apt-get install -y \
  wget
RUN wget https://github.com/mattermost/platform/releases/download/v1.1.0/mattermost.tar.gz
RUN tar zxf mattermost.tar.gz
ADD config.json /mattermost/config/

WORKDIR /mattermost/bin
ENTRYPOINT /mattermost/bin/platform

