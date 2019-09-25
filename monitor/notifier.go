package monitor

import (
	"LianFaPhone/bas-monitor/defaultrpc"
	l4g "github.com/alecthomas/log4go"
)

func connectCenter(status int)  {
	l4g.Info("bas_monitor connectCenter = %d", status)
}

func initNotifier()  {
	nodeInst := defaultrpc.DefaultNodeInst()
	if nodeInst == nil {
		return
	}
}