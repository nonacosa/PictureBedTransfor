package upload


import (
	"github.com/minio/minio-go/v6"
	"log"
)

func main() {
	endpoint := "www.gitrue.com:9000"
	accessKeyID := "pkwenda"
	secretAccessKey := "886pkxiaojiba"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now setup
}
