package monitor

import (
	"LianFaPhone/bas-monitor/common"
	"LianFaPhone/bas-monitor/defaultrpc"
	"LianFaPhone/bas-monitor/logger"
	"LianFaPhone/bas-monitor/rpc"
	"fmt"
	l4g "github.com/alecthomas/log4go"
)

type MyLogger struct {
}

func (l *MyLogger) Debug(arg0 string, args ...interface{}) {
	l4g.Debug(arg0, args...)
}
func (l *MyLogger) Info(arg0 string, args ...interface{}) {
	l4g.Info(arg0, args...)
}

func (l *MyLogger) Trace(arg0 string, args ...interface{}) {
	l4g.Trace(arg0, args...)
}
func (l *MyLogger) Warns(arg0 string, args ...interface{}) {
	l4g.Warn(arg0, args...)
}
func (l *MyLogger) Error(arg0 string, args ...interface{}) {
	l4g.Error(arg0, args...)
}
func (l *MyLogger) Fatal(arg0 string, args ...interface{}) {
	l4g.Critical(arg0, args...)
}

func NewNode(cfg *common.ConfigNode, meta string) (*rpc.Node, error) {
	logger.InitLogger(&MyLogger{})

	if cfg == nil {
		return nil, fmt.Errorf("cfg is nil")
	}

	nodeInst, err := rpc.NewNode(*cfg, meta, connectCenter)
	if err != nil {
		return nil, err
	}

	defaultrpc.SetDefaultNodeInst(nodeInst)

	return nodeInst, nil
}

func BeginNodeMonitor() {
	initCaller()
	initNotifier()
}
