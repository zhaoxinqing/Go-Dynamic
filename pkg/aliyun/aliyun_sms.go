package aliyun

import (
	"backend-go/config"
	"backend-go/pkg/timex"
	"fmt"

	afs "github.com/alibabacloud-go/afs-20180112/client"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

var AliYunSMS *config.AliYunSMS
var AliYunSMSClient *dysmsapi.Client

func InitAliYunSMS(conf *config.AliYun) {

	client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", conf.AccessKeyID, conf.AccessSecret)
	if err != nil {
		zap.L().Error(err.Error(), zap.String("AliYun SMS", "连接 AliYun SMS 失败"))
		panic(fmt.Sprintf("%s - 连接 AliYun SMS 失败", timex.GetUTCFormatTime()))
	}

	AliYunSMS = &conf.SMS
	AliYunSMSClient = client
	zap.L().Info("AliYun SMS 已连接!!!")
}

func SendMessage(phone, smsCode string) error {
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-beijing"
	request.QueryParams["PhoneNumbers"] = phone                         // 手机号
	request.QueryParams["SignName"] = AliYunSMS.SignName                // 阿里云验证过的项目名 自己设置
	request.QueryParams["TemplateCode"] = AliYunSMS.Template            // 阿里云的短信模板号 自己设置
	request.QueryParams["TemplateParam"] = "{\"code\":" + smsCode + "}" // 短信模板中的验证码内容 自己生成
	response, err := AliYunSMSClient.ProcessCommonRequest(request)
	if err != nil {
		zap.L().Error(err.Error(), zap.String("AliYun SMS", " AliYun SMS 验证码发送失败"+err.Error()))
		return err
	}

	err = AliYunSMSClient.DoAction(request, response)
	if err != nil {
		zap.L().Error(err.Error(), zap.String("AliYun SMS", " AliYun SMS 验证码发送失败"+err.Error()))
		return err
	}
	return nil
}

func AliCheckAfs(conf *config.AliYun, sessionId string, token string, sig string, scene string, ip string) (bool, error) {
	c := new(rpc.Config)
	c.SetAccessKeyId(conf.AccessKeyID).
		SetAccessKeySecret(conf.AccessSecret).
		SetRegionId("cn-hangzhou").
		SetEndpoint("afs.aliyuncs.com")
	client, err := afs.NewClient(c)
	if err != nil {
		return false, err
	}
	request := new(afs.AuthenticateSigRequest)
	request.SetSig(sig)
	request.SetSessionId(sessionId)
	request.SetToken(token)
	request.SetScene(scene)
	request.SetRemoteIp(ip)
	request.SetAppKey(conf.Afs.AppKey)
	var response *afs.AuthenticateSigResponse
	response, err = client.AuthenticateSig(request)
	if err != nil {
		zap.L().Error(err.Error(), zap.String("AliYun AFS", " AliYun AFS 滑动验证失败"+*response.Msg))
		return false, err
	}
	code := *response.Code
	zap.L().Info("AliYun AFS AliYun AFS 滑动验证数据" + cast.ToString(response))
	// response 返回code 100 表示验签通过，900 表示验签失败
	if code != 100 {
		return false, nil
	}
	return true, nil
}
