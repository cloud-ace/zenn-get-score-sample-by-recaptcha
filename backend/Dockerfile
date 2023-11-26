FROM golang:1.21.4 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main

FROM alpine:3.18.4

WORKDIR /root/

RUN apk --no-cache add ca-certificates

COPY --from=build /app/main .

# プロジェクトID
ENV PROJECT_ID="{Google CloudのプロジェクトID}"
#reCAPTCHA EnterpriseのキーID
ENV SITE_KEY="{reCAPTCHA EnterpriseのキーID}"

# 検証のため、reCAPTCHA Enterpriseのスコア結果を出力するファイルのパス
ENV LOCAL_FILE_PATH="/root/output"

RUN mkdir -p /root/output

EXPOSE 8081

CMD ["./main"]
