FROM golang:1.19

# Copy source code.
RUN mkdir /indexer
RUN mkdir /indexer/src
COPY ./src /indexer/src

# Copy static files.
RUN mkdir /indexer/static
RUN touch /indexer/static/standard_index_structure.json
COPY ./static/standard_index_structure.json /indexer/static/standard_index_structure.json

COPY go.sum /indexer/
RUN touch /indexer/go.work
RUN touch /indexer/go.mod
RUN touch /indexer/main.go
RUN touch /indexer/.env

COPY go.work /indexer/go.work
COPY go.mod /indexer/go.mod
COPY main.go /indexer/main.go
COPY .env /indexer/.env

RUN mkdir /indexer/maildir
ADD ./static/enron_mail_20110402/maildir/allen-p/_sent_mail /indexer/maildir

WORKDIR /indexer
RUN go build -o indexer main.go
RUN chmod +x indexer
CMD ["./indexer", "/indexer/maildir"]

