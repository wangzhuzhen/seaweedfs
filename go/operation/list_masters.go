package operation

import (
	"github.com/aszxqw/weed-fs/go/glog"
	"github.com/aszxqw/weed-fs/go/util"
	"encoding/json"
)

type ClusterStatusResult struct {
	IsLeader bool     `json:"IsLeader,omitempty"`
	Leader   string   `json:"Leader,omitempty"`
	Peers    []string `json:"Peers,omitempty"`
}

func ListMasters(server string) ([]string, error) {
	jsonBlob, err := util.Get("http://" + server + "/cluster/status")
	glog.V(2).Info("list masters result :", string(jsonBlob))
	if err != nil {
		return nil, err
	}
	var ret ClusterStatusResult
	err = json.Unmarshal(jsonBlob, &ret)
	if err != nil {
		return nil, err
	}
	masters := ret.Peers
	if ret.IsLeader {
		masters = append(masters, ret.Leader)
	}
	return masters, nil
}
