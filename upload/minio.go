package upload

import (
	"fmt"
	"github.com/minio/minio-go/v6"
	"log"
)

func MinioUpload(fileName,path string) string {

	domainPoint := "domainPoint"
	accessKeyID := "accessKeyID"
	secretAccessKey := "secretAccessKey"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(domainPoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	bucketName := "image"
	location := "us-east-1"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the zip file

	contentType := "image/jpg"

	// Upload the zip file with FPutObject
	n, err := minioClient.FPutObject(bucketName, fileName, path, minio.PutObjectOptions{ContentType:contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", fileName, n)

	return fmt.Sprintf("http://%s/%s/%s",domainPoint,bucketName,fileName)
}
