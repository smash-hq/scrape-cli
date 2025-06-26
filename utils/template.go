package utils

import (
	"fmt"
	"os"
)

func CreateTemplate(url gitUrl, targetName string) {
	cwd, err := os.Getwd()
	if err != nil {
		panic("Unable to get the current working directory：" + err.Error())
	}
	repo, err := CloneRepo(Repo{
		URL:         string(url),
		Branch:      "main",
		AccessToken: "",
		TargetName:  targetName,
	}, cwd)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("Template generated in %s\n", repo)
}
