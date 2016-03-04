FROM alpine:3.3

ADD ./build/infiltrator /infiltrator

ENTRYPOINT ["/infiltrator"]
