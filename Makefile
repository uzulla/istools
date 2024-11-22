#export PREBUILD_S3_BUCKET=***
#export PREBUILD_AWS_REGION=ap-northeast-1
#export PREBUILD_DISCORD_WEBHOOK_URL=https://discord.com/api/webhooks/***
#export PREBUILD_AWS_ACCESS_KEY_ID=***
#export PREBUILD_AWS_SECRET_ACCESS_KEY=***
#export PREBUILD_GITHUB_TOKEN=***
#export PREBUILD_GITHUB_ISSUE_URL=https://github.com/***

.PHONY:
build: mkdir build-linux-amd64 build-linux-arm64 build-mac-arm64

mkdir:
	mkdir -p build_linux_amd64
	mkdir -p build_linux_arm64
	mkdir -p build_mac_arm64

.PHONY:
build-linux-amd64:
	cd dupload && \
	GOARCH="amd64" \
    GOOS="linux" \
    go build -ldflags=" \
        -X 'main.prebuildS3Bucket=$$PREBUILD_S3_BUCKET' \
        -X 'main.prebuildAwsRegion=$$PREBUILD_AWS_REGION' \
        -X 'main.prebuildDiscordWebhookURL=$$PREBUILD_DISCORD_WEBHOOK_URL' \
        -X 'main.prebuildAwsAccessKeyID=$$PREBUILD_AWS_ACCESS_KEY_ID' \
        -X 'main.prebuildAwsSecretAccessKey=$$PREBUILD_AWS_SECRET_ACCESS_KEY' \
		" \
		-o ../build_linux_amd64/dupload dupload.go
	cd dpaste && \
	GOARCH="amd64" \
    GOOS="linux" \
    go build -ldflags=" \
        -X 'main.prebuildS3Bucket=$$PREBUILD_S3_BUCKET' \
        -X 'main.prebuildAwsRegion=$$PREBUILD_AWS_REGION' \
        -X 'main.prebuildDiscordWebhookURL=$$PREBUILD_DISCORD_WEBHOOK_URL' \
        -X 'main.prebuildAwsAccessKeyID=$$PREBUILD_AWS_ACCESS_KEY_ID' \
        -X 'main.prebuildAwsSecretAccessKey=$$PREBUILD_AWS_SECRET_ACCESS_KEY' \
		" \
		-o ../build_linux_amd64/dpaste dpaste.go
	cd ghpaste && \
	GOARCH="amd64" \
    GOOS="linux" \
    go build -ldflags=" \
        -X 'main.prebuildGitHubToken=$$PREBUILD_GITHUB_TOKEN' \
        -X 'main.prebuildGitHubIssueUrl=$$PREBUILD_GITHUB_ISSUE_URL' \
		" \
		-o ../build_linux_amd64/ghpaste ghpaste.go

.PHONY:
build-linux-arm64:
	cd dupload && \
	GOARCH="arm64" \
   GOOS="linux" \
   go build -ldflags=" \
       -X 'main.prebuildS3Bucket=$$PREBUILD_S3_BUCKET' \
       -X 'main.prebuildAwsRegion=$$PREBUILD_AWS_REGION' \
       -X 'main.prebuildDiscordWebhookURL=$$PREBUILD_DISCORD_WEBHOOK_URL' \
       -X 'main.prebuildAwsAccessKeyID=$$PREBUILD_AWS_ACCESS_KEY_ID' \
       -X 'main.prebuildAwsSecretAccessKey=$$PREBUILD_AWS_SECRET_ACCESS_KEY' \
	" \
	-o ../build_linux_arm64/dupload dupload.go
	cd dpaste && \
	GOARCH="arm64" \
	   GOOS="linux" \
	   go build -ldflags=" \
		   -X 'main.prebuildS3Bucket=$$PREBUILD_S3_BUCKET' \
		   -X 'main.prebuildAwsRegion=$$PREBUILD_AWS_REGION' \
		   -X 'main.prebuildDiscordWebhookURL=$$PREBUILD_DISCORD_WEBHOOK_URL' \
		   -X 'main.prebuildAwsAccessKeyID=$$PREBUILD_AWS_ACCESS_KEY_ID' \
		   -X 'main.prebuildAwsSecretAccessKey=$$PREBUILD_AWS_SECRET_ACCESS_KEY' \
		" \
		-o ../build_linux_arm64/dpaste dpaste.go
	cd ghpaste && \
	GOARCH="arm64" \
	   GOOS="linux" \
	   go build -ldflags=" \
		   -X 'main.prebuildGitHubToken=$$PREBUILD_GITHUB_TOKEN' \
		   -X 'main.prebuildGitHubIssueUrl=$$PREBUILD_GITHUB_ISSUE_URL' \
		" \
		-o ../build_linux_arm64/ghpaste ghpaste.go

.PHONY:
build-mac-arm64:
	cd dupload && \
	GOARCH="arm64" \
	GOOS="darwin" \
	go build -ldflags=" \
		-X 'main.prebuildS3Bucket=$$PREBUILD_S3_BUCKET' \
		-X 'main.prebuildAwsRegion=$$PREBUILD_AWS_REGION' \
		-X 'main.prebuildDiscordWebhookURL=$$PREBUILD_DISCORD_WEBHOOK_URL' \
		-X 'main.prebuildAwsAccessKeyID=$$PREBUILD_AWS_ACCESS_KEY_ID' \
		-X 'main.prebuildAwsSecretAccessKey=$$PREBUILD_AWS_SECRET_ACCESS_KEY' \
	" \
	-o ../build_mac_arm64/dupload dupload.go
	cd dpaste && \
	GOARCH="arm64" \
	GOOS="darwin" \
	go build -ldflags=" \
		-X 'main.prebuildS3Bucket=$$PREBUILD_S3_BUCKET' \
		-X 'main.prebuildAwsRegion=$$PREBUILD_AWS_REGION' \
		-X 'main.prebuildDiscordWebhookURL=$$PREBUILD_DISCORD_WEBHOOK_URL' \
		-X 'main.prebuildAwsAccessKeyID=$$PREBUILD_AWS_ACCESS_KEY_ID' \
		-X 'main.prebuildAwsSecretAccessKey=$$PREBUILD_AWS_SECRET_ACCESS_KEY' \
	" \
	-o ../build_mac_arm64/dpaste dpaste.go
	cd ghpaste && \
	GOARCH="arm64" \
	GOOS="darwin" \
	go build -ldflags=" \
		-X 'main.prebuildGitHubToken=$$PREBUILD_GITHUB_TOKEN' \
		-X 'main.prebuildGitHubIssueUrl=$$PREBUILD_GITHUB_ISSUE_URL' \
	" \
	-o ../build_mac_arm64/ghpaste ghpaste.go
