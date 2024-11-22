
# build with settings.

```bash
go build -ldflags="
-X 'main.prebuildS3Bucket=***'
-X 'main.prebuildAwsRegion=***'
-X 'main.prebuildDiscordWebhookURL=https://discord.com/***'
-X 'main.prebuildAwsAccessKeyID=***'
-X 'main.prebuildAwsSecretAccessKey=***' " \
-o dpaste dpaste.go
```