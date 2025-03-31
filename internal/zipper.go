package internal

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Zip(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		relPath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		
		if relPath == "." {
			return nil // Skip root directory
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = relPath
		if info.IsDir() {
			header.Name += "/"
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})
}

func Unzip(source, destination string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	os.MkdirAll(destination, 0755)

	for _, file := range reader.File {
		path := filepath.Join(destination, file.Name)
		
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, 0755)
			continue
		}

		os.MkdirAll(filepath.Dir(path), 0755)
		
		destFile, err := os.Create(path)
		if err != nil {
			return err
		}
		
		srcFile, err := file.Open()
		if err != nil {
			destFile.Close()
			return err
		}
		
		_, err = io.Copy(destFile, srcFile)
		srcFile.Close()
		destFile.Close()
		
		if err != nil {
			return err
		}
	}

	return nil
}

