package download_git_repo

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const USER = "wangyi12358"

func Download(repo string, name string) {
	zipUrl := fmt.Sprintf("https://github.com/%s/%s/archive/refs/heads/main.zip", USER, repo)
	output := fmt.Sprintf("%s.zip", name)

	// 发起HTTP GET请求
	response, err := http.Get(zipUrl)
	if err != nil {
		fmt.Printf("HTTP GET请求错误: %s\n", err)
		os.Exit(0)
	}
	defer response.Body.Close()

	// 创建保存文件
	tmpFile, err := os.CreateTemp("", output)
	if err != nil {
		fmt.Printf("创建临时文件错误: %s\n", err)
		os.Exit(0)
	}
	defer func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}()

	// 将HTTP响应的内容写入文件
	_, err = io.Copy(tmpFile, response.Body)
	if err != nil {
		fmt.Printf("写入文件错误: %s\n", err)
		os.Exit(0)
	}

	unzip(tmpFile, name)
}

func unzip(tmpFile *os.File, name string) {

	// 打开ZIP文件
	zipReader, err := zip.OpenReader(tmpFile.Name())
	if err != nil {
		fmt.Errorf("打开ZIP文件错误: %s", err)
		os.Exit(0)
	}
	defer zipReader.Close()

	// 解压缩ZIP文件内容
	for _, file := range zipReader.File {
		if err := saveFile(file, name); err != nil {
			fmt.Errorf("保存文件错误: %s", err)
			os.Exit(0)
		}
	}
}

func saveFile(file *zip.File, name string) error {
	filePath := filepath.Join(name, file.Name)

	if file.FileInfo().IsDir() {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return fmt.Errorf("创建目录错误: %s", err)
	}

	writer, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("创建文件错误: %s", err)
	}
	defer writer.Close()

	reader, err := file.Open()
	if err != nil {
		return fmt.Errorf("打开ZIP文件中的文件错误: %s", err)
	}
	defer reader.Close()

	if _, err = io.Copy(writer, reader); err != nil {
		return fmt.Errorf("写入文件错误: %s", err)
	}

	return nil
}
