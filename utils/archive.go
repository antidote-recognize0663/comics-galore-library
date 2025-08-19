package utils

import (
	"archive/zip"
	"github.com/nwaples/rardecode"
	"io"
	"log"
	"mime/multipart"
	"strings"
)

func CountImages(fileHeader *multipart.FileHeader) (int, error) {
	if fileHeader != nil {
		openFile, err := fileHeader.Open()
		if err != nil {
			return 0, err
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				log.Printf("Could not close file: %s", err.Error())
			}
		}(openFile)
		if strings.HasSuffix(fileHeader.Filename, ".zip") {
			nbrOfPages, err := countImagesInZip(openFile, fileHeader.Size)
			if err != nil {
				return 0, err
			}
			return nbrOfPages, nil
		} else if strings.HasSuffix(fileHeader.Filename, ".rar") {
			nbrOfPages, err := countImagesInRar(openFile)
			if err != nil {
				return 0, err
			}
			return nbrOfPages, nil
		}
	}
	return 0, nil
}

// CountImagesInZip counts image files from a multipart.FileHeader representing a ZIP archive.
func countImagesInZip(file io.ReaderAt, size int64) (int, error) {
	r, err := zip.NewReader(file, size)
	if err != nil {
		return 0, err
	}
	count := 0
	imageExtensions := []string{".jpg", ".jpeg", ".JPEG", ".png", ".gif", ".bmp", ".webp"}
	for _, f := range r.File {
		for _, ext := range imageExtensions {
			if strings.HasSuffix(f.Name, ext) {
				count++
				break
			}
		}
	}
	return count, nil
}

// CountImagesInRar counts the number of image files in a RAR archive.
func countImagesInRar(file io.Reader) (int, error) {
	r, err := rardecode.NewReader(file, "")
	if err != nil {
		return 0, err
	}
	count := 0
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp"}
	for {
		header, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		if !header.IsDir {
			fileName := strings.ToLower(header.Name)
			log.Printf("[RAR] filename: %s", fileName)
			for _, ext := range imageExtensions {
				if strings.HasSuffix(fileName, ext) {
					count++
					break
				}
			}
		}
	}
	return count, nil
}
