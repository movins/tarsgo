package tars

import (
	"os"

	"github.com/rexshan/tarsgo/tars/protocol/res/nodef"
)

//NodeFHelper is helper struct.
type NodeFHelper struct {
	comm *Communicator
	si   nodef.ServerInfo
	sf   *nodef.ServerF
}

//SetNodeInfo sets the node location for the given Communicator.
func (n *NodeFHelper) SetNodeInfo(comm *Communicator, node string, app string, server string) {
	n.comm = comm
	n.sf = new(nodef.ServerF)
	n.sf.SetServant(comm.GetServantProxy(node))
	n.sf.TarsSetTimeout(15000)
	n.si = nodef.ServerInfo{
		app,
		server,
		int32(os.Getpid()),
		"",
		//"tars",
		//container,
	}
}

//KeepAlive sends the keepalive pacakage to the node.
func (n *NodeFHelper) KeepAlive(adapter string) {
	TLOG.Infof("KeepAlive :%s",adapter)
	n.si.Adapter = adapter
	_, err := n.sf.KeepAlive(&n.si)
	if err != nil {
		TLOG.Error("keepalive fail:", adapter)
	}
}

//ReportVersion report the tars version to the node.
func (n *NodeFHelper) ReportVersion(version string) {
	TLOG.Infof("ReportVersion :%s",version)
	_, err := n.sf.ReportVersion(n.si.Application, n.si.ServerName, version)
	if err != nil {
		TLOG.Error("report Version fail:")
	}
}
