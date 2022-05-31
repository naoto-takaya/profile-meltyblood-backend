FROM golang:1.18.2-alpine3.15
WORKDIR /var/www
RUN apk add --no-cache \
    alpine-sdk \
    curl \
    git 
# RUN go install github.com/pilu/fresh
# COPY src/go.mod .

## install go tools
RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@v2.1.2 &&\
    go install github.com/ramya-rao-a/go-outline@1.0.0 &&\
    go install github.com/nsf/gocode@v.20170907 &&\
    go install github.com/acroca/go-symbols@v0.1.1 &&\
    go install github.com/fatih/gomodifytags@v1.16.0 &&\
    go install github.com/josharian/impl@v1.1.0 &&\
    go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0 &&\
    go install github.com/go-delve/delve/cmd/dlv@v1.8.3 &&\
    # golang.org/x/lint/golint \
    go install golang.org/x/tools/gopls@v0.8.4


# RUN go install github.com/pilu/fresh
RUN go install github.com/cosmtrek/air@v1.29.0
# RUN go mod download.com/cosmtrek/air@v1.21.2

# RUN go get -u github
COPY ./src /var/www