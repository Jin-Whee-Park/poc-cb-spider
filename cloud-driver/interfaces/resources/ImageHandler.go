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
//package image

type ImageReqInfo struct {
	Name string
	// @todo
}

type ImageInfo struct {
	Name string
	Id string
	// @todo
}

var test = new(ImageInfo)



type ImageHandler interface {
	CreateImage(imageReqInfo ImageReqInfo) (*ImageInfo, error)
	ListImage() ([]*ImageInfo, error)
	GetImage(imageID string) (*ImageInfo, error)
	DeleteImage(imageID string) (bool, error)
}

