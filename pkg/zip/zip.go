package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Unzip(name string) error {
	// 解压zip文件
	zipReader, err := zip.OpenReader(name + ".zip")
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		err := saveFile(file, name)
		if err != nil {
			return err
		}
	}

	return nil
}

func saveFile(file *zip.File, name string) error {

	filePath := filepath.Join(name, file.Name)

	if file.FileInfo().IsDir() {
		os.MkdirAll(filePath, os.ModePerm)
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	writer, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer writer.Close()

	reader, err := file.Open()
	if err != nil {
		return err
	}
	defer reader.Close()

	if _, err = io.Copy(writer, reader); err != nil {
		return err
	}
	return nil
}
