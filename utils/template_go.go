package utils

import (
	"fmt"
	"os"
)

func CreateGoTemplate(url gitUrl) {
	cwd, err := os.Getwd()
	if err != nil {
		panic("Unable to get the current working directory：" + err.Error())
	}
	repo, err := CloneRepo(Repo{
		URL:         string(url),
		Branch:      "main",
		AccessToken: "",
	}, cwd)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("Template generated in %s", repo)
}
