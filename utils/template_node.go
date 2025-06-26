package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateNodeTemplate() {
	dirName := "go"
	if _, err := os.Stat(dirName); !os.IsNotExist(err) {
		fmt.Printf("Folder '%s' already exists.\n", dirName)
		return
	}
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}
	// 创建 main.go 文件
	mainFilePath := filepath.Join(dirName, "main.go")
	mainFile := `package main

import "fmt"

func main() {
	fmt.Println("Hello, Go template!")
}
`
	err = os.WriteFile(mainFilePath, []byte(mainFile), 0644)
	if err != nil {
		fmt.Println("Error writing main.go:", err)
		return
	}

	fmt.Println("Go template generated in ./go/")
}
