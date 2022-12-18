FROM alpine:latest

# Dockerfile build cache 
ENV REFRESHED_AT 2022-12-04

LABEL "com.github.actions.name"="Github Action for Feishu Notification"
LABEL "com.github.actions.description"="Github Action for Feishu Notification."
LABEL "com.github.actions.icon"="home"
LABEL "com.github.actions.color"="green"
LABEL "repository"="http://github.com/x-actions/feishu"
LABEL "homepage"="http://github.com/x-actions/feishu"
LABEL "maintainer"="xiexianbin<me@xiexianbin.cn>"

LABEL "Name"="Github Action for Feishu Notification"
LABEL "Version"="1.0.0"

ENV LC_ALL C.UTF-8
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US.UTF-8

RUN apk update && apk add --no-cache curl bash && rm -rf /var/cache/apk/*

RUN curl -Lfs -o /usr/local/bin/feishu https://github.com/x-actions/feishu/releases/latest/download/feishu-linux && \
    chmod +x /usr/local/bin/feishu

ADD entrypoint.sh /
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
