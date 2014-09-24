package topology

import (
	"github.com/aszxqw/weed-fs/go/glog"
	"github.com/aszxqw/weed-fs/go/storage"
	"github.com/goraft/raft"
)

type MaxVolumeIdCommand struct {
	MaxVolumeId storage.VolumeId `json:"maxVolumeId"`
}

func NewMaxVolumeIdCommand(value storage.VolumeId) *MaxVolumeIdCommand {
	return &MaxVolumeIdCommand{
		MaxVolumeId: value,
	}
}

func (c *MaxVolumeIdCommand) CommandName() string {
	return "MaxVolumeId"
}

func (c *MaxVolumeIdCommand) Apply(server raft.Server) (interface{}, error) {
	topo := server.Context().(*Topology)
	before := topo.GetMaxVolumeId()
	topo.UpAdjustMaxVolumeId(c.MaxVolumeId)

	glog.V(4).Infoln("max volume id", before, "==>", topo.GetMaxVolumeId())

	return nil, nil
}
