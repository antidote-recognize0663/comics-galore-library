package utils

import (
	"archive/zip"
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
)

func createZipFile(imageFilenames ...string) (*bytes.Reader, error) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	for _, filename := range imageFilenames {
		f, err := w.Create(filename)
		if err != nil {
			return nil, err
		}

		_, err = f.Write([]byte("dummy data"))
		if err != nil {
			return nil, err
		}
	}

	err := w.Close()
	if err != nil {
		return nil, err
	}

	file := bytes.NewReader(buf.Bytes())

	return file, nil
}

func TestCountImageFilesFromZip(t *testing.T) {
	var testCases = []struct {
		name           string
		imageFilenames []string
		expectedCount  int
		shouldErr      bool
	}{
		{
			name:           "NoImageFiles",
			imageFilenames: []string{"file1.txt", "file2.doc", "file3.pdf"},
			expectedCount:  0,
		},
		{
			name:           "SingleImageFile",
			imageFilenames: []string{"file1.jpg"},
			expectedCount:  1,
		},
		{
			name:           "MultipleImageFiles",
			imageFilenames: []string{"file1.jpg", "file2.png", "file3.webp"},
			expectedCount:  3,
		},
		{
			name:           "ImageAndNonImageFiles",
			imageFilenames: []string{"file1.jpg", "file2.txt", "file3.mp4"},
			expectedCount:  1,
		},
		{
			name:           "DifferentImageExtensions",
			imageFilenames: []string{"file1.jpg", "file2.jpeg", "file3.JPEG", "file4.png"},
			expectedCount:  4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//file, err := createZipFile(tc.imageFilenames...)
			file, err := os.Open(fmt.Sprintf("../files/zip/%s.zip", tc.name))
			if err != nil {
				t.Fatalf("Error creating zip file: %v", err)
			}
			stat, err := file.Stat()
			if err != nil {
				t.Errorf("Error getting file stat: %v", err)
			}
			result, err := countImagesInZip(file, stat.Size())
			if err != nil && !tc.shouldErr {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tc.expectedCount {
				t.Errorf("CountImagesInZip(%v) = %v; want %v", tc.imageFilenames, result, tc.expectedCount)
			}
		})
	}
}
func TestCountImageFilesFromRar(t *testing.T) {
	/*wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	t.Logf("Current working directory: %s", wd)*/
	var testCases = []struct {
		name           string
		imageFilenames []string
		expectedCount  int
		shouldErr      bool
	}{
		{
			name:           "NoImageFiles",
			imageFilenames: []string{"file1.txt", "file2.docx", "file3.pdf"},
			expectedCount:  0,
		},
		{
			name:           "SingleImageFile",
			imageFilenames: []string{"file1.jpg"},
			expectedCount:  1,
		},
		{
			name:           "MultipleImageFiles",
			imageFilenames: []string{"file1.jpg", "file2.png", "file3.webp"},
			expectedCount:  3,
		},
		{
			name:           "ImageAndNonImageFiles",
			imageFilenames: []string{"file1.jpg", "file2.txt", "file3.mp4"},
			expectedCount:  1,
		},
		{
			name:           "DifferentImageExtensions",
			imageFilenames: []string{"file1.jpg", "file2.jpeg", "file3.JPEG", "file4.png"},
			expectedCount:  4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(fmt.Sprintf("../files/rar/%s.rar", tc.name))
			if err != nil {
				t.Fatalf("Error opening RAR file: %v", err) // Updated error message
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					log.Printf("Error closing file: %v", err)
				}
			}(file)
			result, err := countImagesInRar(file)
			if err != nil && !tc.shouldErr {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tc.expectedCount {
				t.Errorf("CountImagesInRar(%v) = %v; want %v", tc.imageFilenames, result, tc.expectedCount)
			}
		})
	}
}
