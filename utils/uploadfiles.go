package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func SaveUploadedFiles(files []*multipart.FileHeader) ([]string, error) {
	var imageUrls []string
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		// Generate a unique file name
		fileName := uuid.New().String() + filepath.Ext(file.Filename)
		dst, err := os.Create("assets/uploads/" + fileName)
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		// Copy the file to the destination
		if _, err = io.Copy(dst, src); err != nil {
			return nil, err
		}

		imageUrl := "/uploads/" + fileName
		imageUrls = append(imageUrls, imageUrl)
	}
	return imageUrls, nil
}

func SaveOneUploadedFile(file *multipart.FileHeader) (string, error) {
	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Generate a unique file name using a UUID and retain the original file extension
	fileName := uuid.New().String() + filepath.Ext(file.Filename)

	// Create the destination file path
	dst, err := os.Create("assets/uploads/" + fileName)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the file content to the destination
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	// Return the relative URL of the saved image
	imageUrl := "/uploads/" + fileName
	return imageUrl, nil
}

func SaveUploadedFilesExtentions(files []*multipart.FileHeader) ([]string, []string, error) {
	var imageUrls []string
	var extensions []string

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return nil, nil, err
		}
		defer src.Close()

		// Generate a unique file name
		ext := filepath.Ext(file.Filename)
		fileName := uuid.New().String() + ext
		dst, err := os.Create("assets/uploads/" + fileName)
		if err != nil {
			return nil, nil, err
		}
		defer dst.Close()

		// Copy the file to the destination
		if _, err = io.Copy(dst, src); err != nil {
			return nil, nil, err
		}

		imageUrl := "/uploads/" + fileName
		imageUrls = append(imageUrls, imageUrl)
		extensions = append(extensions, ext)
	}

	return imageUrls, extensions, nil
}
