FROM golang:1.18

#ENV ELASTIC_HOSTS=localhost:9200
#ENV LOG_LEVEL=info

# Configure the repo url so we can configure our work directory:
ENV REPO_URL=github.com/ssoql/faq-chat-bot

# Setup out $GOPATH
ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL

# /app/src/github.com/federicoleon/bookstore_items-api/src

# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH/api

RUN go mod init faq-chat-bot
RUN go mod tidy
RUN go build -o faq-chat-bot-app .

# Expose port 8081 to the world:
EXPOSE 8084

CMD ["./faq-chat-bot-app"]