# 開発用イメージ
FROM golang:1.17.7-alpine as dev
ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk --no-cache add git

RUN go install github.com/cosmtrek/air@v1.29.0 \
  && go install golang.org/x/tools/cmd/goimports@latest \
  && go get -u github.com/go-sql-driver/mysql

COPY go.mod go.sum ./

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]


# 本番用中間イメージ
FROM golang:1.17.7-alpine as builder
ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN go get -u github.com/go-sql-driver/mysql

RUN apk --no-cache add git
COPY go.mod go.sum ./

COPY ./cmd ${ROOT}
RUN GOOS=linux GOARCH=amd64 go build -o main {$ROOT}/main.go

# 本番用ファイナルイメージ
# バイナリファイルが実行できればいいので無駄は全て削ぎ落とした
# 今回は不要だが今後のイメージのために作成した
FROM scratch as prod
ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT} .
EXPOSE 8080
CMD ["./main"]