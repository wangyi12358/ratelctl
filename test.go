package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func extractZipFile(zipPath, destination string) error {
	// 打开ZIP文件
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("打开ZIP文件错误: %s", err)
	}
	defer r.Close()

	// 解压缩ZIP文件内容
	for _, file := range r.File {
		fmt.Printf("解压文件: %s\n", file.Name)
		// 构造解压后的文件路径
		targetPath := filepath.Join(destination, file.Name)

		// 检查文件路径中的目录部分
		dir := filepath.Dir(targetPath)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("创建目录错误: %s", err)
		}

		// 如果不是目录，则解压文件
		if !file.FileInfo().IsDir() {
			// 创建目标文件
			dst, err := os.Create(targetPath)
			if err != nil {
				return fmt.Errorf("创建目标文件错误: %s", err)
			}
			defer dst.Close()

			// 打开ZIP文件中的文件
			src, err := file.Open()
			if err != nil {
				return fmt.Errorf("打开ZIP文件中的文件错误: %s", err)
			}
			defer src.Close()

			// 将ZIP文件中的内容复制到目标文件中
			_, err = io.Copy(dst, src)
			if err != nil {
				return fmt.Errorf("复制文件内容错误: %s", err)
			}
		}
	}

	return nil
}

func main() {
	zipPath := "test.zip"      // 要解压的ZIP文件路径
	destination := "extracted" // 解压后的文件保存路径

	err := extractZipFile(zipPath, destination)
	if err != nil {
		fmt.Printf("解压ZIP文件错误: %s\n", err)
		return
	}

	fmt.Println("ZIP文件解压完成！")
}
