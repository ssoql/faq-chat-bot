FROM golang:1.18

ENV REPO_URL=github.com/ssoql/faq-chat-bot
# Setup out $GOPATH
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL
ENV WORKPATH=$APP_PATH/src
ENV TMPL_PATH=$WORKPATH/api/templates/
COPY src $WORKPATH
WORKDIR $WORKPATH/api

RUN go mod init faq-chat-bot
RUN go mod tidy
RUN go build -o faq-chat-bot-app .

# Expose port 8081 to the world:
EXPOSE 8084

#CMD ["./faq-chat-bot-app"]
ENTRYPOINT ./start_docker.sh