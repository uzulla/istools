package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ビルド時に設定する変数
var (
	prebuildS3Bucket           string
	prebuildAwsRegion          string
	prebuildDiscordWebhookURL  string
	prebuildAwsAccessKeyID     string
	prebuildAwsSecretAccessKey string
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("使用法: dupload <ファイル名> <アップロードキー(オプション)>")
		os.Exit(1)
	}

	var (
		S3Bucket           string
		S3Region           string
		DiscordWebhookURL  string
		AwsAccessKeyId     string
		AwsSecretAccessKey string
	)

	// 必要な環境変数を取得
	if prebuildS3Bucket != "" {
		S3Bucket = prebuildS3Bucket
		S3Region = prebuildAwsRegion
		DiscordWebhookURL = prebuildDiscordWebhookURL
		AwsAccessKeyId = prebuildAwsAccessKeyID
		AwsSecretAccessKey = prebuildAwsSecretAccessKey
	} else {
		S3Region = os.Getenv("AWS_REGION")
		S3Bucket = os.Getenv("S3_BUCKET")
		DiscordWebhookURL = os.Getenv("DISCORD_WEBHOOK_URL")
		AwsAccessKeyId = os.Getenv("AWS_ACCESS_KEY_ID")
		AwsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	}

	if S3Region == "" || S3Bucket == "" || DiscordWebhookURL == "" {
		log.Fatal("環境変数 AWS_REGION, S3_BUCKET, DISCORD_WEBHOOK_URL を設定してください")
	}

	fileName := os.Args[1]

	// ファイルからデータを読み込む
	inputData, err := readFile(fileName)
	if err != nil {
		log.Fatalf("ファイルの読み込みに失敗しました: %v", err)
	}

	// ファイルを読み込んでmimeTypeを取得
	mimeType := http.DetectContentType(inputData)

	// uploadKeyが指定されていない場合はファイル名を使用
	var uploadKey string
	if len(os.Args) > 2 {
		uploadKey = os.Args[2]
	} else {
		uploadKey = fileName
	}

	// S3にアップロード
	s3URL, err := uploadDataToS3(inputData, uploadKey, S3Region, S3Bucket, mimeType, AwsAccessKeyId, AwsSecretAccessKey)
	if err != nil {
		log.Fatalf("S3へのアップロードに失敗しました: %v", err)
	}

	// Discordに投稿
	err = postToDiscord(fileName, s3URL, DiscordWebhookURL)
	if err != nil {
		log.Fatalf("Discordへの投稿に失敗しました: %v", err)
	}

	fmt.Println("正常にアップロードし、Discordに投稿しました")
}

func readFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var buf bytes.Buffer
	reader := bufio.NewReader(file)
	_, err = io.Copy(&buf, reader)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func uploadDataToS3(data []byte, key, region, bucket string, mimeType string, awsAccessKeyId string, awsSecretAccessKey string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			awsAccessKeyId,
			awsSecretAccessKey,
			"",
		),
	})
	if err != nil {
		return "", err
	}

	svc := s3.New(sess)

	// S3にアップロード
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ACL:         aws.String("public-read"), // 公開アクセスを設定
		ContentType: aws.String(mimeType),      // Content-Typeを設定
	})
	if err != nil {
		return "", err
	}

	// アップロードされたファイルのURLを生成
	s3URL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, key)
	return s3URL, nil
}

func postToDiscord(fileName, s3URL, webhookURL string) error {
	message := fmt.Sprintf("アップロードされたファイル: %s\nURL: %s", fileName, s3URL)
	payload := map[string]string{
		"content": message,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Discord webhook がステータスコード %d を返しました", resp.StatusCode)
	}

	return nil
}
