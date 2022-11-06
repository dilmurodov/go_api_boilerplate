FROM golang:1.18.5

ENV LOG_LEVEL=info
ENV ENV=production

ENV DB_PASSWORD="123test"

# Configure the repo url so we can configure our work directory:
ENV REPO_URL=github.com/yhagio/go_api_boilerplate

# Setup out $GOPATH
ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL

# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY ./ $WORKPATH
WORKDIR $WORKPATH

RUN go build -o exe .

CMD ["./exe"]