FROM golang:1.9.2-alpine3.7

RUN apk update \
 && apk add git


#RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 \
# && chmod +x /usr/local/bin/dep


WORKDIR /go/src/beego-ripple
COPY . .


RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/beego/bee

RUN dep ensure -vendor-only

#RUN go-wrapper download
#RUN go-wrapper install
#CMD ["go-wrapper", "run"]
CMD ["bee", "run"]