// Proof of Concepts of CB-Spider.
// The CB-Spider is a sub-Framework of the Cloud-Barista Multi-Cloud Project.
// The CB-Spider Mission is to connect all the clouds with a single interface.
//
//      * Cloud-Barista: https://github.com/cloud-barista
//
// This is a Cloud Driver Example for PoC Test.
//
// by powerkim@etri.re.kr, 2019.06.

//package main
package aws

import (
	"C"

	acon "github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/aws/connect"
	idrv "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces"
	icon "github.com/cloud-barista/poc-cb-spider/cloud-driver/interfaces/connect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	//"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/session"
	//"github.com/aws/aws-sdk-go/service/ec2"
)
import "fmt"

type AwsDriver struct {
}

func (AwsDriver) GetDriverVersion() string {
	return "TEST AWS DRIVER Version 0.1"
}

func (AwsDriver) GetDriverCapability() idrv.DriverCapabilityInfo {
	var drvCapabilityInfo idrv.DriverCapabilityInfo

	drvCapabilityInfo.ImageHandler = false
	drvCapabilityInfo.VNetworkHandler = false
	drvCapabilityInfo.SecurityHandler = false
	drvCapabilityInfo.KeyPairHandler = false
	drvCapabilityInfo.VNicHandler = false
	drvCapabilityInfo.PublicIPHandler = false
	drvCapabilityInfo.VMHandler = true

	return drvCapabilityInfo
}

func getVMClient(regionInfo idrv.RegionInfo) (*ec2.EC2, error) {
	// setup Region
	fmt.Println("AwsDriver : getVMClient() - Region : [" + regionInfo.Region + "]")

	//sess, err := ec2session.NewSession(&aws.Config{
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(regionInfo.Region)},
	)
	if err != nil {
		fmt.Println("Could not create aws New Session", err)
		return nil, err
	}

	// Create EC2 service client
	svc := ec2.New(sess)
	if err != nil {
		fmt.Println("Could not create EC2 service client", err)
		return nil, err
	}

	return svc, nil
}

func (driver *AwsDriver) ConnectCloud(connectionInfo idrv.ConnectionInfo) (icon.CloudConnection, error) {
	// 1. get info of credential and region for Test A Cloud from connectionInfo.
	// 2. create a client object(or service  object) of Test A Cloud with credential info.
	// 3. create CloudConnection Instance of "connect/TDA_CloudConnection".
	// 4. return CloudConnection Interface of TDA_CloudConnection.

	// sample code, do not user like this^^
	//var iConn icon.CloudConnection
	//VMClient, err := getVMClient(connectionInfo.CredentialInfo)
	vmClient, err := getVMClient(connectionInfo.RegionInfo)
	if err != nil {
		return nil, err
	}

	//iConn = acon.AwsCloudConnection{}
	iConn := acon.AwsCloudConnection{
		Region:   connectionInfo.RegionInfo,
		VMClient: vmClient,
	}
	return &iConn, nil // return type: (icon.CloudConnection, error)
}

/*
func (AwsDriver) ConnectCloud(connectionInfo idrv.ConnectionInfo) (icon.CloudConnection, error) {
	// 1. get info of credential and region for Test A Cloud from connectionInfo.
	// 2. create a client object(or service  object) of Test A Cloud with credential info.
	// 3. create CloudConnection Instance of "connect/TDA_CloudConnection".
	// 4. return CloudConnection Interface of TDA_CloudConnection.

	// sample code, do not user like this^^
	var iConn icon.CloudConnection
	iConn = acon.AwsCloudConnection{}
	return iConn, nil // return type: (icon.CloudConnection, error)
}
*/
var TestDriver AwsDriver
