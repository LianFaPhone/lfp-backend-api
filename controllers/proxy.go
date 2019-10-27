package controllers

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	"LianFaPhone/lfp-backend-api/tools"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/kataras/iris"
	"net/http"
	"net/http/httputil"
	"net/url"
	"LianFaPhone/lfp-backend-api/utils"
)

type ProxyController struct {
	CfgProxy *tools.Proxy
	Controllers
}

func (this *ProxyController) Proxy(ctx iris.Context) {
	toPath := this.CfgProxy.ToPrefix + "/" + ctx.Params().Get("param")

	remote, err := url.Parse(this.CfgProxy.ToHost + toPath)
	if err != nil {
		l4g.Error(fmt.Errorf("url[%s] parse err: %v", this.CfgProxy.ToHost+toPath, err))
		ctx.JSON(Response{Code: apibackend.BASERR_INTERNAL_CONFIG_ERROR.Code(), Message: err.Error()})
		return
	}

	targetQuery := remote.RawQuery
	director := func(req *http.Request) {
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = toPath
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		acc,_ := utils.NewUtils().GetValueUserInfo(ctx)
		if acc != nil && acc.Extend != nil{
			req.URL.RawQuery +=  "&extend=" + *acc.Extend
		}

		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}

	}
	proxy := &httputil.ReverseProxy{Director: director}

	w := ctx.ResponseWriter()
	r := ctx.Request()

	// important! https://stackoverflow.com/questions/23164547/golang-reverseproxy-not-working
	r.Host = remote.Host
	proxy.ServeHTTP(w, r)
	if _, ok := this.Config.RecordLogsMap[ctx.Method()+ctx.Path()]; ok {
		ctx.Next()
	}
}
