package model

import (
	"context"

	"github.com/rexshan/tarsgo/tars/protocol/res/requestf"
)

//Servant is interface for call the remote server.
type Servant interface {
	Tars_invoke(ctx context.Context, ctype byte,
		sFuncName string,
		buf []byte,
		status map[string]string,
		context map[string]string,
		Resp *requestf.ResponsePacket) error

	TarsSetTimeout(t int)
	GetProxyEndPoints()[]string
	ProxyInvoke(ctx context.Context, cType byte, sFuncName string, buf []byte, ipPort string,Resp *requestf.ResponsePacket) error
}
