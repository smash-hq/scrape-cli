package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func AutoRunProject() {
	cwd, err := os.Getwd()
	if err != nil {
		panic("Unable to get the current working directory：" + err.Error())
	}
	if isGoProject(cwd) {
		RunGolangProject()
	} else if isNodeProject(cwd) {
		RunNodeProject()
	} else {
		fmt.Println("⚠️ Unable to detect project type. No run action performed.")
	}
}

func RunGolangProject() {
	fmt.Println("▶ Checking Go installation...")
	if _, err := exec.LookPath("go"); err != nil {
		fmt.Println("❌ Go is not installed or not in PATH.")
		return
	}

	fmt.Println("🔧 Running `go mod tidy`...")
	if err := runCommand("go", "mod", "tidy"); err != nil {
		fmt.Println("❌ `go mod tidy` failed.")
		return
	}
	fmt.Println("🔧 Building `main.go`...")
	if err := runCommand("go", "build", "-o", outputBinary()); err != nil {
		fmt.Println("❌ Failed to build project.")
		return
	}

	fmt.Printf("🚀 Running `%s`...\n", outputBinary())
	if err := runCommand(outputBinary()); err != nil {
		fmt.Println("❌ Failed to run project.")
	}
}

func RunNodeProject() {
	fmt.Println("▶ Detected Node.js project.")

	if _, err := exec.LookPath("npm"); err != nil {
		fmt.Println("❌ npm not found. Please install Node.js and npm.")
		return
	}

	fmt.Println("📦 Running `npm install`...")
	if err := runCommand("npm", "install"); err != nil {
		fmt.Println("❌ npm install failed.")
		return
	}

	fmt.Println("🚀 Running `node index.js`...")
	if err := runCommand("node", "index.js"); err != nil {
		fmt.Println("❌ Failed to run node project.")
	}
}

// outputBinary 返回构建输出的文件名（带平台判断）
func outputBinary() string {
	if runtime.GOOS == "windows" {
		return "main.exe"
	}
	return "./main"
}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout // 显示标准输出
	cmd.Stderr = os.Stderr // 显示错误输出
	cmd.Stdin = os.Stdin   // 允许用户交互（如按键输入）
	return cmd.Run()
}

func isGoProject(projectDir string) bool {
	goMod := filepath.Join(projectDir, "go.mod")
	mainGo := filepath.Join(projectDir, "main.go")
	return fileExists(goMod) || fileExists(mainGo)
}

func isNodeProject(projectDir string) bool {
	packageJson := filepath.Join(projectDir, "package.json")
	return fileExists(packageJson)
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
