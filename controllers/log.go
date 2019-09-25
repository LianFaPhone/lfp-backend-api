package controllers

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	"LianFaPhone/lfp-backend-api/api-common"
	"LianFaPhone/lfp-backend-api/models"
	"LianFaPhone/lfp-backend-api/tools"
	"LianFaPhone/lfp-backend-api/utils"
	l4g "github.com/alecthomas/log4go"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/i18n"
	"time"
	"strings"
)

var GLogController LogController

type LogController struct {
	db       *gorm.DB
	logModel *models.LogModel
	config   *tools.Config
}

func NewLogController(config *tools.Config) *LogController {

	GLogController.logModel = models.NewLogModel()
	GLogController.config = config

	//删除操作日志
	go GLogController.Start()

	return &GLogController
}

func (l *LogController) Start() {
	defer utils.PanicPrint()
	CleanIntvl := l.config.OperateLog.CleanIntvl
	if CleanIntvl < 3600 {
		CleanIntvl = 3600
	}
    for {
		l.logModel.DeleteLoginLog(l.config.OperateLog.RemainDays)
		l.logModel.DeleteOperatLog(l.config.OperateLog.RemainDays)
		time.Sleep(time.Second * time.Duration(CleanIntvl))
	}

}




func (l *LogController) RecodeLog(ctx iris.Context) {
	var (
		path   string
		userId uint
		ip     string
		method string
	)

	path = ctx.Path()
	ip = common.GetRealIp(ctx)
	method = ctx.Method()

	if strings.HasSuffix(path, "/account/login") {
		body := ctx.Values().Get("user")
		if body == nil {
			return
		}
		userId := uint(body.(*models.Account).Id)
		if userId <= 0 {
			return
		}
		go l.RecodeLoginLog(userId, ip, "web")
	}else{
		userId = uint(utils.NewUtils().GetValueUserId(ctx))
		if userId <= 0 {
			return
		}
		go l.RecodeOperationLog(userId, ip, method+path)
	}
	return
}

func (l *LogController) RecodeLoginLog(userId uint, ip string, device string) {
	defer utils.PanicPrint()
	var (
		country string
		city    string
	)
	ipInfo, err := utils.IpLocation(l.config.IpFind.Auth, ip)
	if err != nil {
		l4g.Error("IpLocation err", err.Error())
		return
	}

	if ipInfo.Country == "" {
		country = "Unknown"
	} else {
		country = ipInfo.Country
	}

	if ipInfo.City == "" {
		city = "Unknown"
	} else {
		city = ipInfo.City
	}

	if device == "" {
		device = "web"
	}

	_, err = l.logModel.CreateLoginLog(userId, ip, country, city, device)
	if err != nil {
		l4g.Error("CreateLoginLog err", err.Error())
		//		glog.Error(err.Error())
		return
	}
}

func (l *LogController) RecodeOperationLog(userId uint, ip string, operation string) {
	defer utils.PanicPrint()
	var (
		country string
		city    string
	)
	ipInfo, err := utils.IpLocation(l.config.IpFind.Auth, ip)
	if err != nil {
		l4g.Error("IpLocation err", err.Error())
		return
	}

	if ipInfo.Country == "" {
		country = "Unknown"
	} else {
		country = ipInfo.Country
	}

	if ipInfo.City == "" {
		city = "Unknown"
	} else {
		city = ipInfo.City
	}

	recordLog, ok := l.config.RecordLogsMap[operation]
	if !ok {
		recordLog = new(tools.RecordLog)
	}

	_, err = l.logModel.CreateOperationLog(userId, operation, ip, country, city, recordLog.ZhName, recordLog.EnName)
	if err != nil {
		l4g.Error("CreateOperationLog err", err.Error())
		//		glog.Error(err.Error())
		return
	}
}

func (l *LogController) GetLoginLog(ctx iris.Context) {
	var (
		//userId uint
		err    error
		params struct {
			common.PageParams
			UserId uint `params:"user_id"`
		}

		respData []interface{}
	)

	common.GetParams(ctx, &params)
	common.GetParams(ctx, &params.PageParams)

	if params.Limit <= 0 || params.Limit >= 100 {
		params.Limit = 10
	}

	if params.Page <= 1 {
		params.Page = 1
	}

	//userId = common.GetUserIdFromCtx(ctx)

	data, count, err := l.logModel.GetLoginLog(params.UserId, params.Limit, params.Limit*(params.Page-1))
	if err != nil {
		l4g.Error("GetLoginLog err", err.Error())
		//		glog.Error(err.Error())
		ctx.JSON(Response{Code: apibackend.BASERR_DATABASE_ERROR.Code(), Message: err.Error()})
		return
	}

	respData = make([]interface{}, len(data))

	for k, v := range data {
		respData[k] = &struct {
			Id        uint   `json:"id"`
			Ip        string `json:"ip"`
			Country   string `json:"country"`
			City      string `json:"city"`
			Device    string `json:"device"`
			CreatedAt int64  `json:"created_at"`
		}{v.ID, v.Ip, v.Country, v.City, v.Device, v.CreatedAt}
	}

	ctx.JSON(NewResponse(0, "").SetLimitResult(respData, count, params.Page))
}

func (l *LogController) GetOperationLog(ctx iris.Context) {
	var (
		//userId uint
		err    error
		params struct {
			common.PageParams
			UserId uint `params:"user_id"`
		}

		respData []interface{}
	)

	common.GetParams(ctx, &params)
	common.GetParams(ctx, &params.PageParams)

	if params.Limit <= 0 || params.Limit >= 100 {
		params.Limit = 10
	}

	if params.Page <= 1 {
		params.Page = 1
	}

	//userId = common.GetUserIdFromCtx(ctx)

	data, count, err := l.logModel.GetOperationLog(params.UserId, params.Limit, params.Limit*(params.Page-1))
	if err != nil {
		l4g.Error("GetOperationLog err", err.Error())
		//		glog.Error(err.Error())
		ctx.JSON(Response{Code: apibackend.BASERR_DATABASE_ERROR.Code(), Message: err.Error()})
		return
	}

	respData = make([]interface{}, len(data))

	for k, v := range data {
		operation := i18n.Translate(ctx, v.Operation)
		if operation == "" {
			operation = v.Operation
		}
		respData[k] = &struct {
			Id        uint   `json:"id"`
			Operation string `json:"operation"`
			Ip        string `json:"ip"`
			Country   string `json:"country"`
			City      string `json:"city"`
			CreatedAt int64  `json:"created_at"`
		}{v.ID, operation, v.Ip, v.Country, v.City, v.CreatedAt}
	}

	ctx.JSON(NewResponse(0, "").SetLimitResult(respData, count, params.Page))
}
