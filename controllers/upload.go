package controllers

import (
	"UtsuruConcept/db"
	"UtsuruConcept/models"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// UploadRequest handles upload information like the images to upload.
type UploadRequest struct {
	Token  string                  `form:"token" binding:"required"`
	Images []*multipart.FileHeader `form:"images" binding:"required"`
}

// Upload uploads one or more images to Google Cloud and creates associated metadata entries in the database when using
// the "/upload" endpoint.
func Upload(c *gin.Context) {
	var requestBody UploadRequest
	if err := c.ShouldBindWith(&requestBody, binding.FormMultipart); err != nil {
		showResponseError(c, http.StatusBadRequest, err)
		return
	}

	userID, err := validateToken(requestBody.Token)
	if err != nil {
		showResponseError(c, http.StatusBadRequest, err)
		return
	}

	dbObj := db.GetDb()
	user, bucket, err := prepareUpload(c, userID, dbObj)
	if err != nil {
		// The body has been defined in prepareUpload()
		return
	}

	uploadedFiles, notUploadedFiles, err := processUpload(user, requestBody.Images, bucket, dbObj, c)
	if err != nil {
		showResponseError(c, http.StatusInternalServerError, err)
		return
	}

	dbObj.Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"uploaded_files":     uploadedFiles,
		"not_uploaded_files": notUploadedFiles,
	})
}

// prepareUpload retrieves user information and initializes related libraries for processing an image upload.
// It returns the user uploading the pictures, the Google Cloud bucket used for storing pictures and an error.
func prepareUpload(c *gin.Context, userID string, dbObj *gorm.DB) (*models.User, *storage.BucketHandle, error) {
	user, err := models.GetUserById(userID, dbObj)

	if err != nil {
		// The user does not exist
		showResponseError(c, http.StatusBadRequest, err)
		return nil, nil, err
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		// The client has not been configured correctly
		showResponseError(c, http.StatusInternalServerError, err)
		return nil, nil, err
	}

	bucketName := os.Getenv("CLOUD_STORAGE_BUCKET_NAME")
	bucket := client.Bucket(bucketName)
	return user, bucket, nil
}

// processUpload handles the main processing of uploading files to Google Cloud Storage.
// It returns an array of uploaded pictures, pictures that has not been uploaded and an error if something happened.
//
// Note: Two pictures cannot have the same filename in the repository. As such if one is tried to be uploaded, it will
// not be uploaded.
func processUpload(user *models.User, files []*multipart.FileHeader, bucket *storage.BucketHandle, dbObj *gorm.DB, c *gin.Context) ([]string, []string, error) {
	imagesFolderName := os.Getenv("CLOUD_STORAGE_IMAGES_FOLDER")
	uploadedFiles := make([]string, 0)
	notUploadedFiles := make([]string, 0)
	for _, fileInfo := range files {
		userID := user.ID.String()

		if models.IsValidImageExtension(fileInfo.Filename) && !models.IsUserImageExists(userID, fileInfo.Filename, dbObj) {
			userFolder := bucket.Object(fmt.Sprintf("%s/%s/", imagesFolderName, userID))
			ctx := context.Background()
			if _, err := userFolder.Attrs(ctx); err != nil {
				if err = createBucketUserFolder(userFolder); err != nil {
					return nil, nil, err
				}
			}

			if err := saveUploadedFile(fileInfo, c); err != nil {
				return nil, nil, err
			}

			if err := uploadToGoogleCloud(bucket, imagesFolderName, userID, fileInfo.Filename); err != nil {
				return nil, nil, err
			}

			uploadFolder := os.Getenv("UPLOAD_FOLDER")
			imagePath := fmt.Sprintf("%s/%s", uploadFolder, fileInfo.Filename)
			imageData, err := generateImageData(imagePath)
			_ = os.Remove(imagePath)

			if err != nil {
				return nil, nil, err
			}

			imageModel := models.CreateImage(fileInfo.Filename, fileInfo.Size, *imageData)
			user.Images = append(user.Images, *imageModel)
			uploadedFiles = append(uploadedFiles, fileInfo.Filename)
		} else {
			notUploadedFiles = append(notUploadedFiles, fileInfo.Filename)
		}
	}
	return uploadedFiles, notUploadedFiles, nil
}

// createBucketUserFolder creates a new folder for the user inside the Google Cloud Storage bucket.
// It returns an error if the folder could not be created.
func createBucketUserFolder(userFolderHandle *storage.ObjectHandle) error {
	folderWriter := userFolderHandle.NewWriter(context.Background())
	_, err := folderWriter.Write(make([]byte, 0))
	_ = folderWriter.Close()
	if err != nil {
		return err
	}
	return nil
}

// uploadToGoogleCloud uploads the file to the user's folder in Google Cloud Storage.
// It returns an error if the file could not be uploaded to Google Cloud.
func uploadToGoogleCloud(bucket *storage.BucketHandle, imagesFolderName string, userID string, fileName string) error {
	uploadFolder := os.Getenv("UPLOAD_FOLDER")
	imageObject := bucket.Object(fmt.Sprintf("%s/%s/%s", imagesFolderName, userID, fileName))
	storageWriter := imageObject.NewWriter(context.Background())
	savedFileReader, err := os.Open(fmt.Sprintf("%s/%s", uploadFolder, fileName))
	if err != nil {
		return err
	}
	if _, err = io.Copy(storageWriter, savedFileReader); err != nil {
		return err
	}
	defer storageWriter.Close()
	defer savedFileReader.Close()
	return nil
}
