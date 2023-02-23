package mw

import (
	"context"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	endpoint        string = "127.0.0.1:19000"
	accessKeyID     string = "tiktok"
	secretAccessKey string = "tiktokpass"
	useSSL          bool   = false
	BucketName      string = "tiktok"
	MinioLinkPrefix string = "haorui.xyz:8086/tiktok/"
	policy          string = `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetBucketLocation"],"Resource":["arn:aws:s3:::tiktok"]},{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:ListBucket"],"Resource":["arn:aws:s3:::tiktok"],"Condition":{"StringEquals":{"s3:prefix":["*"]}}},{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetObject"],"Resource":["arn:aws:s3:::tiktok/**"]}]}
`
)

var Minio *minio.Client

func InitMinio() {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL})
	if err != nil {
		panic("Minio Connect Failed: " + err.Error())
		return
	}
	Minio = client
	log.Info("Minio connect success:", client)
	exists, err := client.BucketExists(context.Background(), BucketName)
	if err != nil {
		panic("Minio Check Failed: " + err.Error())
	}
	if !exists {
		err := client.MakeBucket(context.Background(), BucketName, minio.MakeBucketOptions{Region: "cn-south-1", ObjectLocking: false})
		if err != nil {
			panic("Minio Init Failed: " + err.Error())
			return
		}
		err = client.SetBucketPolicy(context.Background(), BucketName, policy)
		if err != nil {
			panic("Set Minio Policy Error: " + err.Error())
			return
		}
	}
}
