package download_git_repo

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const USER = "wangyi12358"

func Download(repo string, name string) {
	zipUrl := fmt.Sprintf("https://github.com/%s/%s/archive/refs/heads/main.zip", USER, repo)
	fmt.Printf("zipUrl: %s\n", zipUrl)
	output := fmt.Sprintf("%s.zip", name)

	// 发起HTTP GET请求
	response, err := http.Get(zipUrl)
	if err != nil {
		fmt.Printf("HTTP GET请求错误: %s\n", err)
		os.Exit(0)
	}
	defer response.Body.Close()

	// 创建保存文件
	outputFile, err := os.Create(output)
	if err != nil {
		fmt.Printf("创建文件错误: %s\n", err)
		os.Exit(0)
	}
	defer outputFile.Close()

	// 将HTTP响应的内容写入文件
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		fmt.Printf("写入文件错误: %s\n", err)
		os.Exit(0)
	}
}
