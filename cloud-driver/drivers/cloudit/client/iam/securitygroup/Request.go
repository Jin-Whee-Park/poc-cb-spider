package securitygroup

import (
	"fmt"
	"github.com/cloud-barista/poc-cb-spider/cloud-driver/drivers/cloudit/client"
)

type SecurityGroupRules struct {
	ID         string `json:"id"`
	SecGroupID string `json:"secgroupId"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Port       string `json:"port"`
	Target     string `json:"target"`
	Protocol   string `json:"protocol"`
	Creator    string `json:"creator"`
	CreatedAt  string `json:"createdAt"`
}

type SecurityGroupInfo struct {
	ID          string
	Name        string
	TenantID    string
	Creator     string
	State       string
	CreatedAt   string
	Protection  int
	Rules       []SecurityGroupRules
	RulesCount  int
	Description string
	AsID        string
}

func List(restClient *client.RestClient, requestOpts *client.RequestOpts) (*[]SecurityGroupInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.IAM, "securitygroups")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Get(requestURL, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var securityGroup []SecurityGroupInfo
	if err := result.ExtractInto(&securityGroup); err != nil {
		return nil, err
	}
	return &securityGroup, nil
}

func ListRule(restClient *client.RestClient, securitygroupId string, requestOpts *client.RequestOpts) (*[]SecurityGroupRules, error) {
	requestURL := restClient.CreateRequestBaseURL(client.IAM, "securitygroups", securitygroupId)
	fmt.Println(requestURL)
	
	var result client.Result
	if _, result.Err = restClient.Get(requestURL, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}
	
	var sgRules []SecurityGroupRules
	if err := result.ExtractInto(&sgRules); err != nil {
		return nil, err
	}
	return &sgRules, nil
}

func Get(restClient *client.RestClient, securitygroupId string, requestOpts *client.RequestOpts) (*SecurityGroupInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.IAM, "securitygroups", securitygroupId, "detail")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Get(requestURL, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var securityGroup SecurityGroupInfo
	if err := result.ExtractInto(&securityGroup); err != nil {
		return nil, err
	}
	return &securityGroup, nil
}

func Create(restClient *client.RestClient, requestOpts *client.RequestOpts) (*SecurityGroupInfo, error) {
	requestURL := restClient.CreateRequestBaseURL(client.IAM, "securitygroups")
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Post(requestURL, nil, &result.Body, requestOpts); result.Err != nil {
		return nil, result.Err
	}

	var securityGroup SecurityGroupInfo
	if err := result.ExtractInto(&securityGroup); err != nil {
		return nil, err
	}
	return &securityGroup, nil
}

func Delete(restClient *client.RestClient, securitygroupId string, requestOpts *client.RequestOpts) error {
	requestURL := restClient.CreateRequestBaseURL(client.IAM, "securitygroups", securitygroupId)
	fmt.Println(requestURL)

	var result client.Result
	if _, result.Err = restClient.Delete(requestURL, requestOpts); result.Err != nil {
		return result.Err
	}
	return nil
}
