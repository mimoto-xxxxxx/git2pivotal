package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func getGitConfig(key string) (string, error) {
	data, err := exec.Command("git", "config", key).Output()
	if err != nil {
		return "", fmt.Errorf(`cannot execute "git config %s": %v`, key, err)
	}
	return strings.TrimSpace(string(data)), nil
}

//PivotalTracker API トークンを取得する
func getToken() (string, error) {
	return getGitConfig("git2pivotal.token")
}

//PivotalTracker の ProjectID を取得する
func getProject() (string, error) {
	return getGitConfig("git2pivotal.project")
}

//前に追記するテンプレートを求める
func getPreScriptTemplate() (string, error) {
	return getGitConfig("git2pivotal.prescript")
}

//後ろに追記するテンプレートを求める
func getPostScriptTemplate() (string, error) {
	return getGitConfig("git2pivotal.postscript")
}

//一番最新のコミットを取得する
func getLatestCommit(format string) (string, error) {
	commit, err := exec.Command("git", "log", "-1", "--pretty=format:"+format).Output()
	if err != nil {
		return "", fmt.Errorf(`cannot execute "git log -1": %v`, err)
	}
	return string(commit), nil
}
