package alluxio

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/glog"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
)

const (
	driverName = "alluxio"
	version    = "1.0.0"
)

type driver struct {
	csiDriver        *csicommon.CSIDriver
	nodeId, endpoint string
}

func NewDriver(nodeID, endpoint string) *driver {
	glog.Infof("Driver: %v version: %v", driverName, version)
	csiDriver := csicommon.NewCSIDriver(driverName, version, nodeID)
	csiDriver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME})
	csiDriver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})

	return &driver{
		nodeId:    nodeID,
		endpoint:  endpoint,
		csiDriver: csiDriver,
	}
}

func (d *driver) newControllerServer() *controllerServer {
	return &controllerServer{
		DefaultControllerServer: csicommon.NewDefaultControllerServer(d.csiDriver),
	}
}
func (d *driver) newNodeServer() *nodeServer {
	return &nodeServer{
		nodeId:            d.nodeId,
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d.csiDriver),
	}
}

func (d *driver) Run() {
	s := csicommon.NewNonBlockingGRPCServer()
	s.Start(
		d.endpoint,
		csicommon.NewDefaultIdentityServer(d.csiDriver),
		d.newControllerServer(),
		d.newNodeServer(),
	)
	s.Wait()
}
