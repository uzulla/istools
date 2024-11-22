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
	"regexp"
	"strconv"
)

var (
	prebuildGitHubToken    string
	prebuildGitHubIssueUrl string
)

func main() {
	// 必要な環境変数を取得
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		githubToken = prebuildGitHubToken
	}

	// GitHub Issue URLを解析
	var issueURL string
	if len(os.Args) >= 2 {
		issueURL = os.Args[1]
	} else {
		issueURL = prebuildGitHubIssueUrl
	}
	repoOwner, repoName, issueNumber, err := parseIssueURL(issueURL)
	if err != nil {
		log.Fatalf("Issue URLが正しくありません: %v", err)
	}

	// 標準入力からデータを読み込む
	inputData, err := readStdin()
	if err != nil {
		log.Fatalf("標準入力の読み込みに失敗しました: %v", err)
	}

	// GitHub Issueにコメントを投稿
	err = postCommentToIssue(githubToken, repoOwner, repoName, issueNumber, string(inputData))
	if err != nil {
		log.Fatalf("Issueへのコメント投稿に失敗しました: %v", err)
	}

	fmt.Println("正常にIssueにコメントを投稿しました")
}

func readStdin() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	_, err := io.Copy(&buf, reader)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func postCommentToIssue(token, owner, repo string, issueNumber int, comment string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d/comments", owner, repo, issueNumber)

	payload := map[string]string{
		"body": "```\n" + comment + "\n```",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("GitHub API がステータスコード %d を返しました: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

func parseIssueURL(url string) (string, string, int, error) {
	re := regexp.MustCompile(`https://github.com/([^/]+)/([^/]+)/issues/(\d+)`)
	matches := re.FindStringSubmatch(url)
	if len(matches) != 4 {
		return "", "", 0, fmt.Errorf("URLの形式が正しくありません")
	}

	issueNumber, err := strconv.Atoi(matches[3])
	if err != nil {
		return "", "", 0, fmt.Errorf("Issue番号が正しくありません: %v", err)
	}

	return matches[1], matches[2], issueNumber, nil
}
