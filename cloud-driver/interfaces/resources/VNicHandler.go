// Cloud Driver Interface of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is Resouces interfaces of Cloud Driver.
//
// by powerkim@etri.re.kr, 2019.06.

package resources

type VNicReqInfo struct {
	Name string
	Id   string
	// @todo
}

type VNicInfo struct {
	Name string
	Id   string
	// @todo
}

type VNicHandler interface {
	CreateVNic(vNicReqInfo VNicReqInfo) (VNicInfo, error)
	ListVNic() ([]*VNicInfo, error)
	GetVNic(vNicID string) (VNicInfo, error)
	DeleteVNic(vNicID string) (bool, error)
}
