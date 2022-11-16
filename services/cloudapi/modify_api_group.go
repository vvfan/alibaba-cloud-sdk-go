package cloudapi

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyApiGroup invokes the cloudapi.ModifyApiGroup API synchronously
func (client *Client) ModifyApiGroup(request *ModifyApiGroupRequest) (response *ModifyApiGroupResponse, err error) {
	response = CreateModifyApiGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyApiGroupWithChan invokes the cloudapi.ModifyApiGroup API asynchronously
func (client *Client) ModifyApiGroupWithChan(request *ModifyApiGroupRequest) (<-chan *ModifyApiGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyApiGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyApiGroup(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyApiGroupWithCallback invokes the cloudapi.ModifyApiGroup API asynchronously
func (client *Client) ModifyApiGroupWithCallback(request *ModifyApiGroupRequest, callback func(response *ModifyApiGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyApiGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyApiGroup(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyApiGroupRequest is the request struct for api ModifyApiGroup
type ModifyApiGroupRequest struct {
	*requests.RpcRequest
	DefaultDomain      string               `position:"Query" name:"DefaultDomain"`
	BasePath           string               `position:"Query" name:"BasePath"`
	Description        string               `position:"Query" name:"Description"`
	SecurityToken      string               `position:"Query" name:"SecurityToken"`
	RpcPattern         string               `position:"Query" name:"RpcPattern"`
	UserLogConfig      string               `position:"Query" name:"UserLogConfig"`
	Tag                *[]ModifyApiGroupTag `position:"Query" name:"Tag"  type:"Repeated"`
	CustomerConfigs    string               `position:"Query" name:"CustomerConfigs"`
	GroupId            string               `position:"Query" name:"GroupId"`
	GroupName          string               `position:"Query" name:"GroupName"`
	PassthroughHeaders string               `position:"Query" name:"PassthroughHeaders"`
	CompatibleFlags    string               `position:"Query" name:"CompatibleFlags"`
	CustomTraceConfig  string               `position:"Query" name:"CustomTraceConfig"`
}

// ModifyApiGroupTag is a repeated param struct in ModifyApiGroupRequest
type ModifyApiGroupTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// ModifyApiGroupResponse is the response struct for api ModifyApiGroup
type ModifyApiGroupResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	BasePath    string `json:"BasePath" xml:"BasePath"`
	GroupId     string `json:"GroupId" xml:"GroupId"`
	GroupName   string `json:"GroupName" xml:"GroupName"`
	Description string `json:"Description" xml:"Description"`
	SubDomain   string `json:"SubDomain" xml:"SubDomain"`
}

// CreateModifyApiGroupRequest creates a request to invoke ModifyApiGroup API
func CreateModifyApiGroupRequest() (request *ModifyApiGroupRequest) {
	request = &ModifyApiGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyApiGroup", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyApiGroupResponse creates a response to parse from ModifyApiGroup response
func CreateModifyApiGroupResponse() (response *ModifyApiGroupResponse) {
	response = &ModifyApiGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
