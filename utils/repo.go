package utils

import (
	"archive/zip"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	nethttp "net/http"
	"os"
	"path/filepath"
	"strings"
)

type Repo struct {
	URL         string // 仓库克隆地址，例如：https://github.com/org/repo.git
	Branch      string // 需要拉取的分支名，例如：main、dev
	AccessToken string // GitHub Token（用于认证，公开仓库不需要认证）
}

func CloneRepo(src Repo, localPath string) (string, error) {
	fmt.Printf("Template Source: %s\n", src.URL)
	localCodePath := filepath.Join(localPath, "code")

	owner, repoName := ParseGitURL(src.URL)
	ref := src.Branch // 可以是分支或 tag

	zipURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/zipball/%s", owner, repoName, ref)
	req, err := nethttp.NewRequest("GET", zipURL, nil)
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
		return "", fmt.Errorf("create zip request failed: %v", err)
	}

	token := src.AccessToken
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := nethttp.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf(" Request failed: %v\n", err)
		return "", fmt.Errorf("download git code failed: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != nethttp.StatusOK {
		log.Info("Request failed: %v", resp.Status)
		return "", fmt.Errorf("download git code failed: status code %d", resp.StatusCode)
	}

	zipFilePath := filepath.Join(localPath, "repo.zip")
	if err := os.MkdirAll(localPath, os.ModePerm); err != nil {
		return "", err
	}

	out, err := os.Create(zipFilePath)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(out, resp.Body)
	err = out.Close()
	if err != nil {
		return "", err
	}
	if err := UnzipToTargetDir(zipFilePath, localCodePath); err != nil {
		return "", err
	}
	defer func() {
		err := os.Remove("repo.zip")
		if err != nil {
			return
		}
	}()
	return localCodePath, nil
}

func ParseGitURL(url string) (owner, repo string) {
	url = strings.TrimSuffix(url, ".git")
	parts := strings.Split(url, "/")
	if len(parts) < 2 {
		return "", ""
	}
	repo = parts[len(parts)-1]
	owner = parts[len(parts)-2]
	return owner, repo
}

// UnzipToTargetDir 解压 zipPath 到 targetDir（剥离最外层目录）
func UnzipToTargetDir(zipPath, targetDir string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer func(r *zip.ReadCloser) {
		err := r.Close()
		if err != nil {
			return
		}
	}(r)

	var rootPrefix string

	// 遍历压缩文件条目
	for _, f := range r.File {
		// 获取剥离后的路径（去掉 zip 中的首层目录）
		if rootPrefix == "" {
			parts := strings.SplitN(f.Name, "/", 2)
			if len(parts) == 2 {
				rootPrefix = parts[0] + "/"
			}
		}

		relPath := strings.TrimPrefix(f.Name, rootPrefix)
		if relPath == "" {
			continue
		}

		absPath := filepath.Join(targetDir, relPath)

		// 如果是目录，创建目录
		if f.FileInfo().IsDir() {
			err := os.MkdirAll(absPath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		// 创建所需目录
		if err := os.MkdirAll(filepath.Dir(absPath), os.ModePerm); err != nil {
			return err
		}

		// 打开 zip 中的文件
		rc, err := f.Open()
		if err != nil {
			return err
		}

		// 创建目标文件
		outFile, err := os.Create(absPath)
		if err != nil {
			err := rc.Close()
			if err != nil {
				return err
			}
			return err
		}

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}
		err = outFile.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
