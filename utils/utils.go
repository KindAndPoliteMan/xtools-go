package utils

import (
	"strings"
	"os/exec"
)

func CheckError(err error) {
	if (err != nil) {
		panic(err)
	}
	return
}

func GetGitConfigValue(key string) (string, error) {
	var sb strings.Builder
	cmd := exec.Command("git", "config", key)
	cmd.Stdout = &sb
	if err := cmd.Run(); err != nil {
		return "", err
	}
	var res = strings.TrimSpace(sb.String())
	return res, nil
}
