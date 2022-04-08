FROM golang:1.18
# Setup path with repo
ENV REPO_URL=github.com/ssoql/faq-chat-bot
# Setup $GOPATH
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/src/$REPO_URL
ENV WORKPATH=$APP_PATH/src
# Setup path for templates
ENV TMPL_PATH=$WORKPATH/templates/

COPY src $WORKPATH
WORKDIR $WORKPATH

RUN go mod init faq-chat-bot
RUN go mod tidy
RUN go build -o faq-chat-bot-app .

# Expose port
EXPOSE 8083

CMD ["./faq-chat-bot-app"]