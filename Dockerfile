FROM golang

ENV CGO_ENABLED=0
ENV GO111MODULE=on

WORKDIR /go/src/mclogs

COPY . ./
RUN go build -o /bin/mclogs cmd/mclogs/*.go

FROM scratch
COPY --from=0 /bin/mclogs /bin/mclogs
ENTRYPOINT ["mclogs"]
