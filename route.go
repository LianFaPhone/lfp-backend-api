package main

import (
	"LianFaPhone/lfp-backend-api/controllers"
	"LianFaPhone/lfp-backend-api/tools"
	"github.com/alecthomas/log4go"
	"github.com/go-redis/redis"
	"github.com/kataras/iris"
	"strings"
)

type (
	Service struct {
		App         *iris.Application
		Config      *tools.Config
		RedisClient *redis.Client
	}
)

func (this *Service) routes() {
	this.App = iris.New()

	this.App.Any("/", new(controllers.Index).Index)

	verify := controllers.Verify{}
	verify.Config = this.Config

	redirectC := controllers.NewRedirectController(this.Config)

	var logCtrl = controllers.NewLogController(this.Config)

	v1 := this.App.Party("/v1/bk", verify.VerifyAccess)
	v1.Done(logCtrl.RecodeLog)

	{
		accountParty := v1.Party("/account")
		{
			accounts := controllers.Account{}

			accountParty.Post("/register", accounts.Register)
			accountParty.Any("/search", accounts.Search)
			accountParty.Get("/user-info", accounts.GetUserInfo)
			accountParty.Get("/batch-user-by-ids", accounts.BatchUserByIds)
			accountParty.Any("/update", accounts.Update)
			accountParty.Put("/disabled", accounts.Disabled)
			accountParty.Any("/set-admin", accounts.SetAdmin)
			accountParty.Put("/before-change-password", accounts.ChangeBeforePassword)
			accountParty.Put("/after-change-password", accounts.ChangeAfterPassword)
			accountParty.Put("/change-user-password", accounts.ChangeUserPassword)
			accountParty.Any("/delete", accounts.Delete)

			login := controllers.Login{}
			login.Config = this.Config

			accountParty.Post("/login", login.Login)
			accountParty.Any("/logout", login.Logout)
		}

		accessParty := v1.Party("/access")
		{
			access := controllers.Access{}

			accessParty.Post("/add-access", access.AddAccess)
			accessParty.Get("/search", access.Search)
			accessParty.Delete("/delete", access.Delete)
			accessParty.Put("/update", access.Update)
			accessParty.Get("/search-user-pertain-access", access.SearchUserPertainAccess)
		}

		roleParty := v1.Party("/role")
		{
			role := controllers.Role{}

			roleParty.Post("/add-role", role.AddRule)
			roleParty.Any("/search", role.Search)
			roleParty.Any("/delete", role.Delete)
			roleParty.Any("/update", role.Update)
			roleParty.Put("/disabled", role.Disabled)
		}

		userRoleParty := v1.Party("/user-role")
		{
			userRole := controllers.UserRole{}

			userRoleParty.Post("/set-user-role", userRole.SetUserRule)
			userRoleParty.Get("/search-user-role", userRole.SearchUserRole)
		}

		roleAccessParty := v1.Party("/role-access")
		{
			roleAccess := controllers.RoleAccess{}

			roleAccessParty.Post("/set-role-access", roleAccess.SetRuleAccess)
			roleAccessParty.Get("/search", roleAccess.Search)
		}

		gaParty := v1.Party("/ga")
		{
			ga := controllers.GA{}
			ga.Config = this.Config

			gaParty.Get("/bind", ga.Bind)
			gaParty.Post("/verify", ga.Verify)
			gaParty.Post("/bind-verify", ga.BindVerify)
		}

		logParty := v1.Party("/log")
		{
			logParty.Get("/login", logCtrl.GetLoginLog)
			logParty.Get("/safe", logCtrl.GetOperationLog)
		}

		//这个接口以后优化
		//upFile := controllers.NewUploadFile(this.Config)
		//fileParty := v1.Party("/upload")
		//{
		//	fileParty.Any("/coinlogo", upFile.HandleLogoFiles2)
		//	fileParty.Any("/notice", upFile.HandleNoticeFiles)
		//	fileParty.Any("/notify", upFile.HandleNotifyFiles)
		//	fileParty.Post("/task/add", upFile.AddStatus)
		//	fileParty.Post("/task/update", upFile.UpdateStatus)
		//}

		notifyPy := v1.Party("/notify")
		{ //做代理功能，用Any
			notifyPy.Any("/{param:path}", redirectC.HandlerV1BasNotify)
		}
		//fissioPy := v1.Party("/fissionshare")
		//{ //做代理功能，用Any
		//	fissioPy.Any("/{param:path}", redirectC.HandlerV1Fission)
		//}
		//luckDrawPy := v1.Party("/luckdraw")
		//{ //做代理功能，用Any
		//	luckDrawPy.Any("/{param:path}", redirectC.HandlerV1Fission)
		//}
		//mainPy := v1.Party("/bas-merchant-bk")
		//		//{ //做代理功能，用Any
		//		//	mainPy.Any("/{param:path}", redirectC.HandlerV2Merchant)
		//		//}

		for _, proxy := range this.Config.ProxyList {
			srcPrefix := strings.TrimPrefix(proxy.SrcPrefix, "/v1/bk")
			log4go.Info("proxy[%s] [%v]", srcPrefix, *proxy)
			proxyParty := v1.Party(srcPrefix)
			{
				proxyController := controllers.ProxyController{CfgProxy: proxy}
				proxyController.Config = this.Config
				proxyParty.Any("/{param:path}", proxyController.Proxy)
			}
		}

		//    /v2/bas-merchant-bk/config-device/list
		//proxyParty := v1.Party("/merchant/teammanage")
		//{
		//	proxyController := controllers.ProxyController{}
		//
		//	proxyParty.Any("/{param:path}", proxyController.Proxy2)
		//}
	}

}
