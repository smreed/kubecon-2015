FROM ubuntu:14.04

# Copied from MAINTAINER Peter Parente <parente@cs.unc.edu>
MAINTAINER Steve Reed <steven.m.reed@gmail.com>

RUN echo "deb http://archive.ubuntu.com/ubuntu $(lsb_release -sc) main universe" > /etc/apt/sources.list
RUN echo "deb http://archive.ubuntu.com/ubuntu $(lsb_release -sc)-updates main universe" >> /etc/apt/sources.list
RUN apt-get update

RUN apt-get -y -q install \
  nodejs \
  nodejs-legacy \
  npm \
  wget

RUN wget https://github.com/hakimel/reveal.js/archive/3.0.0.tar.gz
RUN tar xzf 3.0.0.tar.gz
RUN mv /reveal.js-3.0.0 /revealjs

RUN mkdir -p /revealjs/md

WORKDIR /revealjs

RUN npm install -g grunt-cli
RUN npm install
RUN sed -i Gruntfile.js -e 's/port: port,/port: port, hostname: "",/'
RUN sed -i Gruntfile.js -e 's/livereload: true,/livereload: false,/'
ADD index.html /revealjs/
ADD title.js /revealjs/plugin/

EXPOSE 8000
VOLUME ["/revealjs/md/"]
ADD slides.md /revealjs/md/
CMD ["grunt", "serve"]
