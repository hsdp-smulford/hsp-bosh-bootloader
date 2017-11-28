package vsphere

import (
	"github.com/cloudfoundry/bosh-bootloader/bosh"
	"github.com/cloudfoundry/bosh-bootloader/storage"
)

type InputGenerator struct {
}

func NewInputGenerator() InputGenerator {
	return InputGenerator{}
}

func (i InputGenerator) Generate(state storage.State) (map[string]interface{}, error) {
	cidr := state.VSphere.Subnet
	parsedCIDR, _ := bosh.ParseCIDRBlock(cidr)
	return map[string]interface{}{
		"vsphere_subnet":            cidr,
		"jumpbox_ip":                parsedCIDR.GetNthIP(5).String(),
		"bosh_director_internal_ip": parsedCIDR.GetNthIP(6).String(),
		"internal_gw":               parsedCIDR.GetNthIP(1).String(),
		"vcenter_cluster":           state.VSphere.Cluster,
		"network_name":              state.VSphere.Network,
		"vcenter_user":              state.VSphere.VCenterUser,
		"vcenter_password":          state.VSphere.VCenterPassword,
		"vcenter_ip":                state.VSphere.VCenterIP,
		"vcenter_dc":                state.VSphere.VCenterDC,
		"vcenter_rp":                state.VSphere.VCenterRP,
		"vcenter_ds":                state.VSphere.VCenterDS,
	}, nil
}
