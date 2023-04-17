package aliyun

import (
	"backend-go/config"
	"backend-go/pkg/timex"
	"fmt"
	"mime/multipart"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
)

var AliYunOSS config.AliYunOSS
var OSSBucket *oss.Bucket

func InitAliYunOSS(conf *config.AliYun) {
	client, err := oss.New(conf.OSS.Endpoint, conf.AccessKeyID, conf.AccessSecret)
	if err != nil {
		zap.L().Error(err.Error(), zap.String("Aws", "创建 OSSClient 实例失败"))
		panic(fmt.Sprintf("%s - AliYun OSS Client 连接失败", timex.GetUTCFormatTime()))
	}

	// 获取存储空间。
	bucket, err := client.Bucket(conf.OSS.Bucket)
	if err != nil {
		zap.L().Error(err.Error(), zap.String("AliYun OSS", "获取 OSS 存储空间失败"))
		panic(fmt.Sprintf("%s - AliYun OSS 获取存储空间失败", timex.GetUTCFormatTime()))
	}

	OSSBucket = bucket
	AliYunOSS = conf.OSS
	zap.L().Info("AliYun OSS 已连接!!!")
}

func GetOSSCdnUrl() string {
	return AliYunOSS.OssUrl + "/"
}

// Upload 上传文件。
func Upload(file multipart.File, filename string) error {
	if err := OSSBucket.PutObject(filename, file); err != nil {
		zap.L().Error(err.Error(), zap.String("AliYun OSS", "上传文件失败"))
		return err
	}
	return nil
}

func UploadImage(file multipart.File, filename string, ext string) error {
	if err := OSSBucket.PutObject(filename, file, oss.ContentType("image/"+ext)); err != nil {
		zap.L().Error(err.Error(), zap.String("AliYun OSS", "图片上传失败"))
		return err
	}
	return nil
}
