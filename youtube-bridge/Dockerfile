FROM golang

ARG json_config
ARG env 

ENV ENV $env
ENV JSON_CONFIG $json_config

COPY ./*.go /go/src/github.com/matteo-hertel/MusicBridge/youtube-bridge/

WORKDIR /go/src/github.com/matteo-hertel/MusicBridge/youtube-bridge

RUN echo $JSON_CONFIG | tr -d [=\'=]  > client_secret.$ENV.json

RUN go get ./
RUN go build -o youtube-bridge

CMD ./youtube-bridge
	
EXPOSE 8080
