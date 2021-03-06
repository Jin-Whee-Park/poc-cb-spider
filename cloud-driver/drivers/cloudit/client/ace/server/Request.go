package server

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client/iam/securitygroup"
)

type ServerInfo struct {
	VolumeInfoList interface{}
	VmNicInfoList  interface{}
	NicMapInfo     []struct {
		Name    string
		Count   int
		Address string `json:"addr"`
	}
	PoolMapInfo []struct {
		Name       string
		Count      int
		PoolID     string `json:"pool_id"`
		FileSystem string
	}
	AdaptiveIpMapInfo []struct {
		IP        string
		Count     int
		PrivateIP string `json:"private_ip"`
	}
	ID                string
	TenantID          string
	CpuNum            float32
	MemSize           float32
	VncPort           int
	RepeaterPort      int
	State             string
	NodeIp            string
	NodeHostName      string
	Name              string
	Protection        int
	CreatedAt         string
	IsoId             string
	IsoPath           string
	Iso               string
	Template          string
	TemplateID        string
	OsType            string
	RootPassword      string
	HostName          string
	Creator           string
	VolumeId          string
	VolumeSize        int
	VolumeMode        string
	MacAddr           string
	Spec              string
	SpecId            string
	Pool              string
	PoolId            string
	Cycle             int
	Metric            int
	MigrationPort     int
	MigrationIp       string
	Cloudinit         bool
	DeleteVolume      bool
	ServerCount       int
	PrivateIp         string
	AdaptiveIp        string
	InitCloud         int
	ClusterId         string
	ClusterName       string
	NicType           string
	Secgroups         []securitygroup.SecurityGroupRules
	Ip                string
	SubnetAddr        string
	DeviceId          string
	Gpu               string
	GpuCount          int
	GpuId             string
	Description       string
	DiskSize          int
	DiskCount         int
	IsoInsertedAt     string
	Puppet            int
	SshKeyName        string
	SshPublicKey      string
	TemplateOwnership string
	TemplateType      string
	VmStatInfo        string
}

func List(restClient *client.RestClient, requestOpts *client.RequestOpts) (*[]ServerInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers")
	fmt.Println(requestURL)
	
	var result client.Result
	if _, result.Err = restClient.Get(requestURL, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}
	
	var server []ServerInfo
	if err := result.ExtractInto(&server); err != nil {
		return nil, err
	}
	return &server, nil
}

func Get(restClient *client.RestClient, id string, requestOpts *client.RequestOpts) (*ServerInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers", id)
	fmt.Println(requestURL)
	
	var result client.Result
	_, result.Err = restClient.Get(requestURL, &result.Body, requestOpts)
	
	var server ServerInfo
	if err := result.ExtractInto(&server); err != nil {
		return nil, err
	}
	return &server, nil
}

// create
func Start(restClient *client.RestClient, requestOpts *client.RequestOpts) (*ServerInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Post(requestURL, nil, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var server ServerInfo
	if err := result.ExtractInto(&server); err != nil {
		return nil, err
	}
	
	return &server, nil
}

//shutdown
func Suspend(restClient *client.RestClient, id string, requestOpts *client.RequestOpts) error {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers", id, "shutdown")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Post(requestURL, nil, nil, requestOpts); result.Err != nil {
		return result.Err
	}
	return nil
}

//start
func Resume(restClient *client.RestClient, id string, requestOpts *client.RequestOpts) error {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers", id, "start")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Post(requestURL, nil, nil, requestOpts); result.Err != nil {
		return result.Err
	}
	return nil
}

//reboot
func Reboot(restClient *client.RestClient, id string, requestOpts *client.RequestOpts) error {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers", id, "reboot")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Post(requestURL, nil, nil, requestOpts); result.Err != nil {
		return result.Err
	}
	return nil
}

//delete
func Terminate(restClient *client.RestClient, id string, requestOpts *client.RequestOpts) error {
	requestURL := restClient.CreateRequestBaseURL(client.ACE, "servers", id)
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Delete(requestURL, requestOpts); result.Err != nil {
		return result.Err
	}
	return nil
}
