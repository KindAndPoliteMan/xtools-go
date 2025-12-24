package template

import (
	"strings"
	"text/template"
)

type TemplateData struct {
	Name string
	Version string
	MaintainerName string
	MaintainerEmail string
	Homepage string
	Distfiles []string
}

var (
	packageTemplate = `
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
distfiles="{{ join .Distfiles "\n\t" }}"
checksum=badbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadb`
	tFuncs = map[string]interface{}{
		"join": strings.Join,
	}
)

func (t TemplateData) GeneratePackage() (string, error) {
	var sb strings.Builder
	tmpl, err := template.New("void-package").Funcs(tFuncs).Parse(packageTemplate)
	if err != nil {
		return "", err
	}
	if err := tmpl.Execute(&sb, t); err != nil {
		return "", err
	}
	res := strings.TrimSpace(sb.String())
	return res, nil
}
