package minio

import (
	"context"
	"fmt"
	"log"
	"offer_tiktok/pkg/constants"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	Client *minio.Client
	err    error
)

func MakeBucket(ctx context.Context, bucketName string) {
	exists, err := Client.BucketExists(ctx, bucketName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !exists {
		err = Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Successfully created mybucket %v\n", bucketName)
	}
}

func init() {
	ctx := context.Background()
	Client, err = minio.New(constants.MinioEndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinioAccessKeyID, constants.MinioSecretAccessKey, ""),
		Secure: constants.MiniouseSSL})
	if err != nil {
		log.Fatalln("minio连接错误: ", err)
	}

	log.Printf("%#v\n", Client)

	MakeBucket(ctx, constants.MinioVideoBucketName)
	MakeBucket(ctx, constants.MinioImgBucketName)
}
