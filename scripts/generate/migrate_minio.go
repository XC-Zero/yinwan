package main

import (
	"context"
	"github.com/XC-Zero/yinwan/pkg/client"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
)

var initBucketList = []string{"images", "documents", "pdfs"}

func GenerateBucket() {
	if client.MinioClient == nil {
		panic(errors.New("未初始化 Minio 客户端!"))
	}
	for _, s := range initBucketList {
		exists, err := client.MinioClient.BucketExists(context.Background(), s)
		if err != nil {
			panic(err)
		}
		if exists {
			err = client.MinioClient.RemoveBucket(context.Background(), s)
			if err != nil {
				panic(err)
			}
		}
	}
	for _, s := range initBucketList {
		err := client.MinioClient.MakeBucket(context.Background(), s, minio.MakeBucketOptions{
			Region: "cn-north-1",
		})
		if err != nil {
			panic(err)
		}
	}
}
