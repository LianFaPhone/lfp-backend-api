package controllers

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	"LianFaPhone/lfp-backend-api/tools"
	"LianFaPhone/lfp-backend-api/utils"
	"encoding/json"
	l4g "github.com/alecthomas/log4go"
	"github.com/kataras/iris"
	"io/ioutil"
	"net/http"
)

type AdminResponse struct {
	ctx iris.Context
	// response status
	Status struct {
		// response code
		Code int `json:"code"`
		// response msg
		Msg string `json:"msg"`
	} `json:"status"`
	// response result
	Result interface{} `json:"result"`
}

func NewRedirectController(config *tools.Config) *RedirectController {
	bp := &RedirectController{
		config: config,
	}
	return bp
}

type RedirectController struct {
	config *tools.Config
}

type ResNotifyMsg struct {
	Err                 *int        `json:"err,omitempty"`
	ErrMsg              *string     `json:"errmsg,omitempty"`
	TemplateGroupList   interface{} `json:"templategrouplist,omitempty"`
	Templates           interface{} `json:"template,omitempty"`
	TemplateHistoryList interface{} `json:"templatehistorylist,omitempty"`
}

func (this *ResNotifyMsg) GetErr() int {
	if this.Err == nil {
		return 0
	}
	return *this.Err
}

func (this *ResNotifyMsg) GetErrMsg() string {
	if this.ErrMsg == nil {
		return ""
	}
	return *this.ErrMsg
}

func (bp *RedirectController) HandlerV1BasNotify(ctx iris.Context) {
	query := ctx.Request().URL.RawQuery
	newUrl := bp.config.Notify.Addr + ctx.Path()
	if len(query) != 0 {
		newUrl = newUrl + "?" + query
	}
	req, err := http.NewRequest(ctx.Method(), newUrl, ctx.Request().Body)
	if err != nil {
		l4g.Error("http NewRequest username[%s] err[%s]", utils.GetValueUserName(ctx), err.Error())
		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_CONFIG_ERROR.Code(), Message: "NewRequest_ERROR:" + err.Error()})
		return
	}

	l4g.Debug("url:%s ", newUrl)

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		l4g.Error("http Do username[%s] url[%s] err[%s]", utils.GetValueUserName(ctx), newUrl, err.Error())
		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BasNotify_HTTP_DO_ERROR:" + err.Error()})
		return
	}
	if resp.StatusCode != 200 {
		l4g.Error("http Do Response username[%s] url[%s] err[%s]", utils.GetValueUserName(ctx), newUrl, resp.Status)
		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BasNotify_HTTP_RESPONSE_ERROR:" + resp.Status})
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l4g.Error("Body readAll username[%s] err[%s]", utils.GetValueUserName(ctx), err.Error())
		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_READ_BODY_ERROR:" + err.Error()})
		return
	}
	defer resp.Body.Close()

	if len(content) == 0 {
		l4g.Error("username[%s] admin.api[%s] response is null", utils.GetValueUserName(ctx), newUrl)
		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response body content null"})
		return
	}
	if string(content) == "Not Found" {
		l4g.Error("username[%s] admin.api[%s] response is Not Found", utils.GetValueUserName(ctx), newUrl)
		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response is Not Found"})
		return
	}
	l4g.Debug("BASNOTIFY response content[%s]", string(content))
	adminRes := new(ResNotifyMsg)
	if err := json.Unmarshal(content, adminRes); err != nil {
		l4g.Error("Unmarshal username[%s] content[%s] err[%s]", utils.GetValueUserName(ctx), string(content), err.Error())
		ctx.JSON(&Response{Code: apibackend.BASERR_DATA_UNPACK_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response cannot Unmarshal, " + err.Error()})
		return
	}

	if adminRes.GetErr() != 0 {
		l4g.Error("BASNOTIFY username[%s] Response.Status.Code[%d] err[%s]", utils.GetValueUserName(ctx), adminRes.GetErr(), adminRes.GetErrMsg())
		ctx.JSON(&Response{Code: adminRes.GetErr(), Message: "BASNOTIFY_REDIRECT_ERROR: " + adminRes.GetErrMsg()})
		return
	}
	if adminRes.TemplateGroupList != nil {
		ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.TemplateGroupList})
		return
	}
	if adminRes.Templates != nil {
		ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.Templates})
		return
	}
	if adminRes.TemplateHistoryList != nil {
		ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.TemplateHistoryList})
		return
	}
	ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg()})
	//	l4g.Debug("deal HandleV1Admin username[%s] ok", utils.GetValueUserName(ctx))

	if _, ok := bp.config.RecordLogsMap[ctx.Method()+ctx.Path()]; ok {
		ctx.Next()
	}
	return
}

//func (bp *RedirectController) HandlerV1Fission(ctx iris.Context) {
//	query := ctx.Request().URL.RawQuery
//	newUrl := bp.config.MarketFissionApi.Addr + ctx.Path()
//	if len(query) != 0 {
//		newUrl =newUrl + "?" + query
//	}
//	req, err := http.NewRequest(ctx.Method(), newUrl, ctx.Request().Body)
//	if err != nil {
//		l4g.Error("http NewRequest username[%s] err[%s]", utils.GetValueUserName(ctx), err.Error())
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_CONFIG_ERROR.Code(), Message: "NewRequest_ERROR:" + err.Error()})
//		return
//	}
//
//	l4g.Debug("url:%s ", newUrl)
//
//	req.Header.Set("Content-Type", "application/json")
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		l4g.Error("http Do username[%s] url[%s] err[%s]", utils.GetValueUserName(ctx), newUrl, err.Error())
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BasNotify_HTTP_DO_ERROR:" + err.Error()})
//		return
//	}
//	if resp.StatusCode != 200 {
//		l4g.Error("http Do Response username[%s] url[%s] err[%s]", utils.GetValueUserName(ctx), newUrl, resp.Status)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BasNotify_HTTP_RESPONSE_ERROR:" + resp.Status})
//		return
//	}
//
//	content, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		l4g.Error("Body readAll username[%s] err[%s]", utils.GetValueUserName(ctx), err.Error())
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_READ_BODY_ERROR:"+ err.Error()})
//		return
//	}
//	defer resp.Body.Close()
//
//	if len(content) == 0 {
//		l4g.Error("username[%s] admin.api[%s] response is null", utils.GetValueUserName(ctx), newUrl)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response body content null"})
//		return
//	}
//	if string(content) == "Not Found" {
//		l4g.Error("username[%s] admin.api[%s] response is Not Found", utils.GetValueUserName(ctx), newUrl)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response is Not Found"})
//		return
//	}
//	l4g.Debug("BASNOTIFY response content[%s]", string(content))
//	ctx.Write(content)
//	//adminRes := new(Response)
//	//if err := json.Unmarshal(content, adminRes); err != nil {
//	//	l4g.Error("Unmarshal username[%s] content[%s] err[%s]", utils.GetValueUserName(ctx), string(content), err.Error())
//	//	ctx.JSON(&Response{Code: apibackend.BASERR_DATA_UNPACK_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response cannot Unmarshal, " + err.Error()})
//	//	return
//	//}
//	//
//	//if adminRes.Code != 0 {
//	//	l4g.Error("BASNOTIFY username[%s] Response.Status.Code[%d] err[%s]", utils.GetValueUserName(ctx), adminRes.Code, adminRes.Message)
//	//	ctx.JSON(&Response{Code: adminRes.Code, Message: "BASNOTIFY_REDIRECT_ERROR: " + adminRes.Message})
//	//	return
//	//}
//	//if adminRes.TemplateGroupList != nil {
//	//	ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.TemplateGroupList})
//	//	return
//	//}
//	//if adminRes.Templates != nil {
//	//	ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.Templates})
//	//	return
//	//}
//	//if adminRes.TemplateHistoryList != nil {
//	//	ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.TemplateHistoryList})
//	//	return
//	//}
//	//ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg()})
//	////	l4g.Debug("deal HandleV1Admin username[%s] ok", utils.GetValueUserName(ctx))
//	ctx.Next()
//	return
//}
//
//func (bp *RedirectController) HandlerV1LuckDraw(ctx iris.Context) {
//	query := ctx.Request().URL.RawQuery
//	newUrl := bp.config.MarketLuckDrawApi.Addr + ctx.Path()
//	if len(query) != 0 {
//		newUrl =newUrl + "?" + query
//	}
//	req, err := http.NewRequest(ctx.Method(), newUrl, ctx.Request().Body)
//	if err != nil {
//		l4g.Error("http NewRequest username[%s] err[%s]", utils.GetValueUserName(ctx), err.Error())
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_CONFIG_ERROR.Code(), Message: "NewRequest_ERROR:" + err.Error()})
//		return
//	}
//
//	l4g.Debug("url:%s ", newUrl)
//
//	req.Header.Set("Content-Type", "application/json")
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		l4g.Error("http Do username[%s] url[%s] err[%s]", utils.GetValueUserName(ctx), newUrl, err.Error())
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BasNotify_HTTP_DO_ERROR:" + err.Error()})
//		return
//	}
//	if resp.StatusCode != 200 {
//		l4g.Error("http Do Response username[%s] url[%s] err[%s]", utils.GetValueUserName(ctx), newUrl, resp.Status)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BasNotify_HTTP_RESPONSE_ERROR:" + resp.Status})
//		return
//	}
//
//	content, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		l4g.Error("Body readAll username[%s] err[%s]", utils.GetValueUserName(ctx), err.Error())
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_READ_BODY_ERROR:"+ err.Error()})
//		return
//	}
//	defer resp.Body.Close()
//
//	if len(content) == 0 {
//		l4g.Error("username[%s] admin.api[%s] response is null", utils.GetValueUserName(ctx), newUrl)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response body content null"})
//		return
//	}
//	if string(content) == "Not Found" {
//		l4g.Error("username[%s] admin.api[%s] response is Not Found", utils.GetValueUserName(ctx), newUrl)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response is Not Found"})
//		return
//	}
//	l4g.Debug("BASNOTIFY response content[%s]", string(content))
//	ctx.Write(content)
//	//adminRes := new(Response)
//	//if err := json.Unmarshal(content, adminRes); err != nil {
//	//	l4g.Error("Unmarshal username[%s] content[%s] err[%s]", utils.GetValueUserName(ctx), string(content), err.Error())
//	//	ctx.JSON(&Response{Code: apibackend.BASERR_DATA_UNPACK_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response cannot Unmarshal, " + err.Error()})
//	//	return
//	//}
//	//
//	//if adminRes.Code != 0 {
//	//	l4g.Error("BASNOTIFY username[%s] Response.Status.Code[%d] err[%s]", utils.GetValueUserName(ctx), adminRes.Code, adminRes.Message)
//	//	ctx.JSON(&Response{Code: adminRes.Code, Message: "BASNOTIFY_REDIRECT_ERROR: " + adminRes.Message})
//	//	return
//	//}
//	//if adminRes.TemplateGroupList != nil {
//	//	ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.TemplateGroupList})
//	//	return
//	//}
//	//if adminRes.Templates != nil {
//	//	ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.Templates})
//	//	return
//	//}
//	//if adminRes.TemplateHistoryList != nil {
//	//	ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.TemplateHistoryList})
//	//	return
//	//}
//	//ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg()})
//	////	l4g.Debug("deal HandleV1Admin username[%s] ok", utils.GetValueUserName(ctx))
//	ctx.Next()
//	return
//}
//
//func (bp *RedirectController) HandlerV2Merchant(ctx iris.Context) {
//	query := ctx.Request().URL.RawQuery
//	newPath := "/v2" + ctx.Path()[6:]
//	newUrl := bp.config.MarketLuckDrawApi.Addr + newPath
//	if len(query) != 0 {
//		newUrl =newUrl + "?" + query
//	}
//	req, err := http.NewRequest(ctx.Method(), newUrl, ctx.Request().Body)
//	if err != nil {
//		l4g.Error("http NewRequest username[%s] err[%s]", utils.GetValueUserName(ctx), err.Error())
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_CONFIG_ERROR.Code(), Message: "NewRequest_ERROR:" + err.Error()})
//		return
//	}
//
//	l4g.Debug("url:%s ", newUrl)
//
//	req.Header.Set("Content-Type", "application/json")
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		l4g.Error("http Do username[%s] url[%s] err[%s]", utils.GetValueUserName(ctx), newUrl, err.Error())
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BasNotify_HTTP_DO_ERROR:" + err.Error()})
//		return
//	}
//	if resp.StatusCode != 200 {
//		l4g.Error("http Do Response username[%s] url[%s] err[%s]", utils.GetValueUserName(ctx), newUrl, resp.Status)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BasNotify_HTTP_RESPONSE_ERROR:" + resp.Status})
//		return
//	}
//
//	content, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		l4g.Error("Body readAll username[%s] err[%s]", utils.GetValueUserName(ctx), err.Error())
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_READ_BODY_ERROR:"+ err.Error()})
//		return
//	}
//	defer resp.Body.Close()
//
//	if len(content) == 0 {
//		l4g.Error("username[%s] admin.api[%s] response is null", utils.GetValueUserName(ctx), newUrl)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response body content null"})
//		return
//	}
//	if string(content) == "Not Found" {
//		l4g.Error("username[%s] admin.api[%s] response is Not Found", utils.GetValueUserName(ctx), newUrl)
//		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response is Not Found"})
//		return
//	}
//	l4g.Debug("BASNOTIFY response content[%s]", string(content))
//	ctx.Write(content)
//	//adminRes := new(Response)
//	//if err := json.Unmarshal(content, adminRes); err != nil {
//	//	l4g.Error("Unmarshal username[%s] content[%s] err[%s]", utils.GetValueUserName(ctx), string(content), err.Error())
//	//	ctx.JSON(&Response{Code: apibackend.BASERR_DATA_UNPACK_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response cannot Unmarshal, " + err.Error()})
//	//	return
//	//}
//	//
//	//if adminRes.Code != 0 {
//	//	l4g.Error("BASNOTIFY username[%s] Response.Status.Code[%d] err[%s]", utils.GetValueUserName(ctx), adminRes.Code, adminRes.Message)
//	//	ctx.JSON(&Response{Code: adminRes.Code, Message: "BASNOTIFY_REDIRECT_ERROR: " + adminRes.Message})
//	//	return
//	//}
//	//if adminRes.TemplateGroupList != nil {
//	//	ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.TemplateGroupList})
//	//	return
//	//}
//	//if adminRes.Templates != nil {
//	//	ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.Templates})
//	//	return
//	//}
//	//if adminRes.TemplateHistoryList != nil {
//	//	ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.TemplateHistoryList})
//	//	return
//	//}
//	//ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg()})
//	////	l4g.Debug("deal HandleV1Admin username[%s] ok", utils.GetValueUserName(ctx))
//	ctx.Next()
//	return
//}
