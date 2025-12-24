package main

import (
	"fmt"

	"github.com/KindAndPoliteMan/xtools-go/utils"
	"github.com/KindAndPoliteMan/xtools-go/cmd/template"
)

func main() {
	fmt.Println("Hello, world!")
	userName, err := utils.GetGitConfigValue("user.name")
	utils.CheckError(err)
	userEmail, err := utils.GetGitConfigValue("user.email")
	utils.CheckError(err)
	fmt.Printf("%s\n%s\n", userName, userEmail)
	t := template.TemplateData{
		Name: "hi",
		Version: "0.1.0",
		MaintainerName: userName,
		MaintainerEmail: userEmail,
		Homepage: "https://example.com",
		Distfiles: []string{
			"aaa",
			"bbb",
		},
	}
	res, err := t.GeneratePackage()
	utils.CheckError(err)
	fmt.Printf("%s\n", res)
}
