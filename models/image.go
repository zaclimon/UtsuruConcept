package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
)

type Image struct {
	ID        uuid.UUID `gorm:"primary_key"`
	FileName  string    `gorm:"not null"`
	Url       string
	Size      int64
	UserID    uuid.UUID
	ImageData ImageData `json:"-"`
}

func (image *Image) BeforeCreate(scope *gorm.Scope) error {
	generatedUuid, err := uuid.NewRandom()
	if err != nil {
		fmt.Errorf("error while generating UUID for the image")
		return err
	}

	return scope.SetColumn("ID", generatedUuid)
}

func IsValidImageExtension(fileName string) bool {
	fileParts := strings.Split(fileName, ".")
	if len(fileParts) >= 1 {
		fileExtension := strings.ToLower(fileParts[len(fileParts)-1])
		return strings.EqualFold(fileExtension, "jpg") ||
			strings.EqualFold(fileExtension, "jpeg") ||
			strings.EqualFold(fileExtension, "png") ||
			strings.EqualFold(fileExtension, "gif")
	}
	return false
}

func IsUserImageExists(userID string, fileName string, dbObj *gorm.DB) bool {
	var image Image
	dbObj.First(&image, "user_id = ? AND file_name = ?", userID, fileName)
	return image.FileName != ""
}

func CreateImage(fileName string, size int64, imageData ImageData) *Image {

	image := &Image{
		FileName:  fileName,
		Size:      size,
		ImageData: imageData,
	}

	return image
}

func DecodeImage(filePath string) (image.Image, error) {
	imageReader, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	image, _, err := image.Decode(imageReader)

	if err != nil {
		return nil, err
	}
	imageReader.Close()
	return image, nil
}
