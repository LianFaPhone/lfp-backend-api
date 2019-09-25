package controllers

//import (
//	"LianFaPhone/lfp-backend-api/tools"
//	//s3 "LianFaPhone/bas-tools/sdk.aws.s3"
//	l4g "github.com/alecthomas/log4go"
//	"github.com/kataras/iris"
//
//	"crypto/md5"
//	"encoding/json"
//	"fmt"
//	"LianFaPhone/lfp-backend-api/utils"
//	"io/ioutil"
//	"strings"
//	apibackend "LianFaPhone/lfp-api/errdef"
//	"net/http"
//)
//
//func NewUploadFile(config *tools.Config) *UploadFile {
//	return &UploadFile{
//		mConfig: config,
//	}
//}
//
//type UploadFile struct {
//	mConfig *tools.Config
//}
//
//type LogoFileAddr struct {
//	File string `json:"file"`
//	Addr string `json:"addr"`
//}
//
//func (this *UploadFile) HandleLogoFiles(ctx iris.Context) {
//	l4g.Debug("start deal HandleLogoFiles username[%s]", utils.GetValueUserName(ctx))
//	//设置内存大小
//	err := ctx.Request().ParseMultipartForm(32 << 21)
//	if err != nil {
//		l4g.Error("ParseMultipartForm username[%s] err[%s]", utils.GetValueUserName(ctx), err.Error())
//		ctx.JSON(&Response{Code: apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), Message: "ParseMultipartForm_ERRORS:" + err.Error()})
//		return
//	}
//	files := ctx.Request().MultipartForm.File["file"]
//	if len(files) == 0 {
//		l4g.Error("NoFile username[%s] err", utils.GetValueUserName(ctx))
//		ctx.JSON(&Response{Code: apibackend.BASERR_INVALID_PARAMETER.Code(), Message: "NoFile_ERRORS"})
//		return
//	}
//	l4g.Debug("username[%s] fileNum[%d] ", utils.GetValueUserName(ctx), len(files))
//	//暂时s3Sdk作为一个临时的，因为压根很少用到，何必占内存呢
//	s3Sdk := s3.NewS3Sdk(this.mConfig.Aws.LogoRegion, this.mConfig.Aws.AccessKeyId, this.mConfig.Aws.AccessKey, this.mConfig.Aws.AccessToken)
//	if s3Sdk == nil {
//		l4g.Error("NewS3Sdk username[%s] err[return nil]", utils.GetValueUserName(ctx))
//		ctx.JSON(&Response{Code: apibackend.BASERR_UNKNOWN_BUG.Code(), Message: "NewS3Sdk_ERRORS"})
//		return
//	}
//	defer s3Sdk.Close()
//	timeout := this.mConfig.Aws.LogoTimeout
//	if timeout < 10 {
//		timeout = 10
//	}
//	results := make([]LogoFileAddr, 0)
//	for i := 0; i < len(files); i++ {
//		filename := files[i].Filename
//		l4g.Debug("username[%s] fileName[%s] ", utils.GetValueUserName(ctx), filename)
//		file, err := files[i].Open()
//		if err != nil {
//			l4g.Error("FileOpen[%s] username[%s] err[%s]", filename, utils.GetValueUserName(ctx), err.Error())
//			ctx.JSON(&Response{Code: apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), Message: "FileOpen_ERRORS:" + filename})
//			return
//		}
//		defer file.Close()
//		addr, err := s3Sdk.UpLoad(this.mConfig.Aws.LogoBucket, filename, file, timeout,nil)
//		if err != nil {
//			l4g.Error("S3 UpLoad[%s] username[%s] err[%s]", filename, utils.GetValueUserName(ctx), err.Error())
//			ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "S3_UpLoad_ERRORS:" + filename})
//			return
//		}
//		results = append(results, LogoFileAddr{filename, addr})
//	}
//	content, err := json.Marshal(results)
//	if err != nil {
//		l4g.Error("Marshal username[%s] err[%s]", utils.GetValueUserName(ctx), err.Error())
//		ctx.JSON(&Response{Code:apibackend.BASERR_DATA_PACK_ERROR.Code(), Message: "Marshal_LogFileAddr_ERROR:" + err.Error()})
//		return
//	}
//	ctx.JSON(&Response{Code: 0, Data: string(content)})
//	l4g.Debug("deal HandleLogoFiles username[%s] ok, result[%s]", utils.GetValueUserName(ctx), string(content))
//}
//
//func (this *UploadFile) HandleLogoFiles2(ctx iris.Context) {
//	this.handleFilesToAwsS3(ctx, this.mConfig.Aws.LogoRegion, this.mConfig.Aws.LogoBucket, "", "", this.mConfig.Aws.LogoTimeout)
//}
//
//func (this *UploadFile) HandleNoticeFiles(ctx iris.Context) {
//	//	filePrefix := fmt.Sprintf("%d", time.Now().Unix())
//	this.handleFilesToAwsS3(ctx, this.mConfig.Aws.NoticeRegion, this.mConfig.Aws.NoticeBucket, "image", "", this.mConfig.Aws.NoticeTimeout)
//}
//
//func (this *UploadFile) HandleNotifyFiles(ctx iris.Context) {
//	//	filePrefix := fmt.Sprintf("%d", time.Now().Unix())
//	tp := ctx.URLParam("type")
//	subpath := "image"
//	if(tp == "video") {
//		subpath = "video"
//	}
//	this.handleFilesToAwsS3(ctx, this.mConfig.Aws.NotifyRegion, this.mConfig.Aws.NotifyBucket, subpath, "", this.mConfig.Aws.NotifyTimeout)
//}
//
//func (this *UploadFile) handleFilesToAwsS3(ctx iris.Context, region, bucket, bucketPath, filePrefix string, timeout int) {
//	l4g.Debug("start deal handleFilesToAwsS3 username[%s]", utils.GetValueUserName(ctx))
//	//设置内存大小
//	err := ctx.Request().ParseMultipartForm(32 << 23)
//	if err != nil {
//		l4g.Error("ParseMultipartForm username[%s] err[%s]", utils.GetValueUserName(ctx), err.Error())
//		ctx.JSON(&Response{Code: apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), Message: "ParseMultipartForm_ERRORS:" + err.Error()})
//		return
//	}
//	files := ctx.Request().MultipartForm.File["file"]
//	if len(files) == 0 {
//		l4g.Error("NoFile username[%s] err", utils.GetValueUserName(ctx))
//		ctx.JSON(&Response{Code: apibackend.BASERR_INVALID_PARAMETER.Code(), Message: "NoFile_ERRORS"})
//		return
//	}
//	l4g.Debug("username[%s] fileNum[%d] ", utils.GetValueUserName(ctx), len(files))
//	//暂时s3Sdk作为一个临时的，因为压根很少用到，何必占内存呢
//	s3Sdk := s3.NewS3Sdk(region, this.mConfig.Aws.AccessKeyId, this.mConfig.Aws.AccessKey, this.mConfig.Aws.AccessToken)
//	if s3Sdk == nil {
//		l4g.Error("NewS3Sdk username[%s] err[return nil]", utils.GetValueUserName(ctx))
//		ctx.JSON(&Response{Code: apibackend.BASERR_UNKNOWN_BUG.Code(), Message: "NewS3Sdk_ERRORS"})
//		return
//	}
//	defer s3Sdk.Close()
//	if timeout < 10 {
//		timeout = 10
//	}
//	results := make([]LogoFileAddr, 0)
//	for i := 0; i < len(files); i++ {
//		filename := files[i].Filename
//		l4g.Debug("username[%s] fileName[%s] ", utils.GetValueUserName(ctx), filename)
//		file, err := files[i].Open()
//		if err != nil {
//			l4g.Error("FileOpen[%s] username[%s] err[%s]", filename, utils.GetValueUserName(ctx), err.Error())
//			ctx.JSON(&Response{Code: apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), Message: "FileOpen_ERRORS:" + filename})
//			return
//		}
//		defer file.Close()
//		content, err := ioutil.ReadAll(file)
//		if err != nil {
//			l4g.Error("FileReadAll[%s] username[%s] err[%s]", filename, utils.GetValueUserName(ctx), err.Error())
//			ctx.JSON(&Response{Code: apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), Message: "File_ReadAll_ERRORS:" + filename})
//			return
//		}
//		_, err = file.Seek(0, 0)
//		if err != nil {
//			l4g.Error("FileSeek[%s] username[%s] err[%s]", filename, utils.GetValueUserName(ctx), err.Error())
//			ctx.JSON(&Response{Code: apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), Message: "File_Seek_ERRORS:" + filename})
//			return
//		}
//		filenameMd5 := fmt.Sprintf("%X", md5.Sum(content))
//		if len(bucketPath) != 0 {
//			bucketPath = strings.TrimRight(bucketPath, "/")
//			filenameMd5 = bucketPath + "/" + filePrefix + filenameMd5
//		}
//		l4g.Info("S3 start UpLoad[%s]md5[%s] username[%s]", filename, filenameMd5, utils.GetValueUserName(ctx))
//		addr, err := s3Sdk.UpLoad(bucket, filenameMd5, file, timeout,nil)
//		if err != nil {
//			l4g.Error("S3 UpLoad[%s][%s] username[%s] err[%s]", filename, filenameMd5, utils.GetValueUserName(ctx), err.Error())
//			ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "S3_UpLoad_ERRORS:" + filename})
//			return
//		}
//		addr = strings.TrimRight(addr, "/")
//		results = append(results, LogoFileAddr{files[i].Filename, addr})
//		l4g.Info("%s=%s", files[i].Filename, addr)
//	}
//	content, err := json.Marshal(results)
//	if err != nil {
//		l4g.Error("Marshal username[%s] err[%s]", utils.GetValueUserName(ctx), err.Error())
//		ctx.JSON(&Response{Code: apibackend.BASERR_DATA_PACK_ERROR.Code(), Message: "Marshal_LogFileAddr_ERROR:" + err.Error()})
//		return
//	}
//	ctx.JSON(&Response{Code: 0, Data: string(content)})
//	l4g.Debug("deal handleFilesToAwsS3 username[%s] ok, result[%s]", utils.GetValueUserName(ctx), string(content))
//}
//
//func (this *UploadFile) AddStatus(ctx iris.Context) {
//	if len(this.mConfig.BasQuote.Addr) < 5 {//暂时把这个功能放在 quote-collect中
//		l4g.Error("host[%s] config not set", this.mConfig.BasQuote.Addr)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_CONFIG_ERROR.Code(), Message: "BasQuoteAddr_ERROR"})
//		return
//	}
//
//	url := this.mConfig.BasQuote.Addr + "/v1/filetransfer/status/add" +"?" + ctx.Request().URL.RawQuery
//	req, err := http.NewRequest("POST", url, ctx.Request().Body)
//	if err != nil {
//		l4g.Error("http.NewRequest url[%s] err[%v]", url, err)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_CONFIG_ERROR.Code(), Message: "NewRequest err "+err.Error()})
//		return
//	}
//
//	req.Header.Set("Content-Type", "application/json")
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		l4g.Error("http Post[%s] err[%v]", url, err)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: err.Error()})
//		return
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != 200 {
//		l4g.Error("http[%s] response err, status[%d][%s]", url, resp.StatusCode, resp.Status)
//		ctx.JSON(&Response{Code:  apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: resp.Status})
//		return
//	}
//
//	content, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		l4g.Error("Body readAll  err[%v]", err)
//		ctx.JSON(&Response{Code:  apibackend.BASERR_SERVICE_UNKNOWN_ERROR.Code(), Message: err.Error()})
//		return
//	}
//
//	if len(content) == 0 {
//		l4g.Error("upload response is null")
//		ctx.JSON(&Response{Code:  apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "null data"})
//		return
//	}
//	//if string(content) == "Not Found" {
//	//	ZapLog().With(zap.String("url", url)).Error("download response is Not Found")
//	//	ctx.JSON(common.NewErrorResponse(ctx, nil, "", apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code()))
//	//	return
//	//}
//	resMsg := &struct{
//		Code int `json:"code"`
//		Message string `json:"message"`
//	}{}
//
//	if err := json.Unmarshal(content, resMsg); err != nil {
//		l4g.Error("json Unmarshal err[%s]content[%s]", err.Error(), string(content))
//		ctx.JSON(&Response{Code:apibackend.BASERR_DATA_UNPACK_ERROR.Code(), Message:"json unpack err"})
//		return
//	}
//	if resMsg.Code != 0 {
//		l4g.Error("upload[%s] response is %d", url, resMsg.Code)
//		ctx.JSON(&Response{Message: fmt.Sprintf("%d-%s",resMsg.Code, resMsg.Message), Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code()})
//		return
//	}
//
//
//	ctx.JSON(&Response{Message:resMsg.Message})
//}
//
//
//func (this *UploadFile) UpdateStatus(ctx iris.Context) {
//	if len(this.mConfig.BasQuote.Addr) < 5 {//暂时把这个功能放在 quote-collect中
//		l4g.Error("host[%s] config not set", this.mConfig.BasQuote.Addr)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_CONFIG_ERROR.Code(), Message: "BasQuoteAddr_ERROR"})
//		return
//	}
//
//	url := this.mConfig.BasQuote.Addr + "/v1/filetransfer/status/update" +"?" + ctx.Request().URL.RawQuery
//	req, err := http.NewRequest("POST", url, ctx.Request().Body)
//	if err != nil {
//		l4g.Error("http.NewRequest url[%s] err[%v]", url, err)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_CONFIG_ERROR.Code(), Message: "NewRequest err "+err.Error()})
//		return
//	}
//
//	req.Header.Set("Content-Type", "application/json")
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		l4g.Error("http Post[%s] err[%v]", url, err)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: err.Error()})
//		return
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != 200 {
//		l4g.Error("http response err, status[%d][%s]", resp.StatusCode, resp.Status)
//		ctx.JSON(&Response{Code:  apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: resp.Status})
//		return
//	}
//
//	content, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		l4g.Error("Body readAll  err[%v]", err)
//		ctx.JSON(&Response{Code:  apibackend.BASERR_SERVICE_UNKNOWN_ERROR.Code(), Message: err.Error()})
//		return
//	}
//
//	if len(content) == 0 {
//		l4g.Error("upload response is null")
//		ctx.JSON(&Response{Code:  apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "null data"})
//		return
//	}
//	//if string(content) == "Not Found" {
//	//	ZapLog().With(zap.String("url", url)).Error("download response is Not Found")
//	//	ctx.JSON(common.NewErrorResponse(ctx, nil, "", apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code()))
//	//	return
//	//}
//	resMsg := &struct{
//		Code int `json:"code"`
//		Message string `json:"message"`
//	}{}
//
//	if err := json.Unmarshal(content, resMsg); err != nil {
//		l4g.Error("json Unmarshal err[%s]content[%s]", err.Error(), string(content))
//		ctx.JSON(&Response{Code:apibackend.BASERR_DATA_UNPACK_ERROR.Code(), Message:"json unpack err"})
//		return
//	}
//	if resMsg.Code != 0 {
//		l4g.Error("upload[%s] response is %d", url, resMsg.Code)
//		ctx.JSON(&Response{Message: fmt.Sprintf("%d-%s",resMsg.Code, resMsg.Message), Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code()})
//		return
//	}
//
//
//	ctx.JSON(&Response{Message:resMsg.Message})
//}
//
