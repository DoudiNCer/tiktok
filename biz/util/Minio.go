package util

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/DodiNCer/tiktok/biz/mw"
	"github.com/minio/minio-go/v7"
	"io"
	"time"
)

func MinioUploadVideo(reader io.Reader, size int64) (objName string, err error) {
	m5 := md5.New()
	m5.Write([]byte(time.Now().String()))
	m5.Write([]byte("video"))
	objname := hex.EncodeToString(m5.Sum(nil)) + ".mp4"
	_, err = mw.Minio.PutObject(context.Background(), mw.BucketName, objName, reader, size, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return objname, nil
}

func MinioUploadPhoto(reader io.Reader, size int64) (objName string, err error) {
	m5 := md5.New()
	m5.Write([]byte(time.Now().String()))
	m5.Write([]byte("photo"))
	objname := hex.EncodeToString(m5.Sum(nil)) + ".jpg"
	_, err = mw.Minio.PutObject(context.Background(), mw.BucketName, objName, reader, size, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return objname, nil
}
