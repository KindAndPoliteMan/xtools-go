package template

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/KindAndPoliteMan/xtools-go/utils"
)

var template = `
# Template file for '{{ .Name }}'
pkgname={{ .Name }}
version={{ .Version }}
revision=1
#archs="i686 x86_64"
#build_wrksrc=
build_style=gnu-configure
#configure_args=""
#make_build_args=""
#make_install_args=""
#conf_files=""
#make_dirs="/var/log/dir 0755 root root"
hostmakedepends=""
makedepends=""
depends=""
short_desc=""
maintainer="{{ .MaintainerName }} <{{ .MaintainerEmail }}>"
license="GPL-3.0-or-later"
homepage="{{ .Homepage }}"
#changelog=""
distfiles="{{ .Distfiles }}"
checksum=badbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadb
`

func getGitConfigValue(key string) ([]byte, error) {
	cmd := exec.Command(fmt.Sprintf("git config user.name %s", key))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	return data, nil
}

func generateTemplate(name string, version string, userName string, userEmail string, homepage string, distfiles string) error {
	if userName == "" {
		if data, err := getGitConfigValue("user.name"); err != nil {
			return err
		} else {
			userName = string(data)
		}
	}
	if userEmail == "" {
		if data, err := getGitConfigValue("user.email"); err != nil {
			return err
		} else {
			userEmail = string(data)
		}
	}
	data := map[string]string{
		"Name":            name,
		"Version":         version,
		"MaintainerName":  userName,
		"MaintainerEmail": userEmail,
		"Homepage":        homepage,
		"Distfiles":       distfiles,
	}
}
