package main

import (
	basel4g "LianFaPhone/lfp-base/log/l4g"

	"LianFaPhone/lfp-backend-api/tools"
	l4g "github.com/alecthomas/log4go"
	"github.com/urfave/cli"
	"os"
	"time"
	//"LianFaPhone/bas-monitor/rpc"
	//monitorcommon "LianFaPhone/bas-monitor/common"
	//"LianFaPhone/lfp-backend-api/monitor"
	//"context"
)

var (
	Meta = ""
)

func main() {
	//laxFlag := config.NewLaxFlagDefault()
	//cfgDir := laxFlag.String("conf_path", config.GetBastionPayConfigDir(), "config path")
	//logPath := laxFlag.String("log_path", config.GetBastionPayConfigDir()+"/log.xml", "log conf path")
	//laxFlag.LaxParseDefault()
	//fmt.Printf("commandline param: conf_path=%s, log_path=%s\n", *cfgDir, *logPath)

	command := Command{}
	com := command.NewCli()
	com.Cli.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "run Commands",
			Action: func(c *cli.Context) {
				conf := tools.Analyze(command.ConfigPath)

				svc := Service{
					Config: conf,
				}
				basel4g.LoadConfig(command.LogConfPath)
				l4g.Info("conf[%v]", *conf)

				// new monitor node
				//var rpcNode *rpc.Node
				//if conf.BasMonitor.RpcAddr != "" {
				//	var err error
				//	monitorCfg := monitorcommon.ConfigNode{}
				//	monitorCfg.Version = conf.BasMonitor.Version
				//	monitorCfg.Name = conf.BasMonitor.Name
				//	monitorCfg.Tag = conf.BasMonitor.Tag
				//	monitorCfg.RpcAddr = conf.BasMonitor.RpcAddr
				//	monitorCfg.Env = conf.BasMonitor.Env
				//	rpcNode, err = monitor.NewNode(&monitorCfg, Meta)
				//	if err != nil {
				//		l4g.Warn("bas_monitor err: %v", err)
				//	}
				//}
				//
				//basutils.GlobalMonitor.Start(conf.Monitor.Addr)
				//models.GlobalNotifyMgr.Init(conf)
				//models.GlobalNotifyMgr.Start()

				// start monitor node
				//ctx, cancel := context.WithCancel(context.Background())
				//monitor.BeginNodeMonitor()
				//if rpcNode != nil {
				//	rpc.StartNode(ctx, rpcNode)
				//}

				//			svc.RunLogrus()
				l4g.Info("start svc...")
				svc.Run()

				//cancel()
				//if rpcNode != nil {
				//	rpc.StopNode(rpcNode)
				//}
				l4g.Info("stop svc...")
			},
		},
	}

	defer basel4g.Close()
	//defer models.GlobalNotifyMgr.Close()

	if err := com.Cli.Run(os.Args); err != nil {
		l4g.Error(err, "run server errors")
	}
	l4g.Info("Stoped.....")
	time.Sleep(time.Second * 1)
}
