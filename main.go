package main

import (
	"context"
	"fmt"

	operations "github.com/Asad2730/S3_v2_CRUD/opetations"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)
	downloader := manager.NewDownloader(client)

	operations.Create_Update(uploader, "your-bucketName", "Your-object-key", "path/to/your/local/file.txt")
	operations.Read(downloader, "your-bucketName", "Your-object-key", "path/to/your/local/downloaded_file.txt")
	operations.Delete(client, "your-bucketName", "Your-object-key")
}
