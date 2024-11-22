# build

```
go build -ldflags="
-X 'main.prebuildGitHubToken=ghp_****'
-X 'main.prebuildGitHubIssueUrl=https://github.com/uzulla/istools/issues/1'
" \
-o ghpaste ghpaste.go
```

- Class patで、repo 権限が必要