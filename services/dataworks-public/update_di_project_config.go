package dataworks_public

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

// UpdateDIProjectConfig invokes the dataworks_public.UpdateDIProjectConfig API synchronously
func (client *Client) UpdateDIProjectConfig(request *UpdateDIProjectConfigRequest) (response *UpdateDIProjectConfigResponse, err error) {
	response = CreateUpdateDIProjectConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDIProjectConfigWithChan invokes the dataworks_public.UpdateDIProjectConfig API asynchronously
func (client *Client) UpdateDIProjectConfigWithChan(request *UpdateDIProjectConfigRequest) (<-chan *UpdateDIProjectConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateDIProjectConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDIProjectConfig(request)
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

// UpdateDIProjectConfigWithCallback invokes the dataworks_public.UpdateDIProjectConfig API asynchronously
func (client *Client) UpdateDIProjectConfigWithCallback(request *UpdateDIProjectConfigRequest, callback func(response *UpdateDIProjectConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDIProjectConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateDIProjectConfig(request)
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

// UpdateDIProjectConfigRequest is the request struct for api UpdateDIProjectConfig
type UpdateDIProjectConfigRequest struct {
	*requests.RpcRequest
	DestinationType string           `position:"Query" name:"DestinationType"`
	SourceType      string           `position:"Query" name:"SourceType"`
	ProjectConfig   string           `position:"Query" name:"ProjectConfig"`
	ProjectId       requests.Integer `position:"Query" name:"ProjectId"`
}

// UpdateDIProjectConfigResponse is the response struct for api UpdateDIProjectConfig
type UpdateDIProjectConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateUpdateDIProjectConfigRequest creates a request to invoke UpdateDIProjectConfig API
func CreateUpdateDIProjectConfigRequest() (request *UpdateDIProjectConfigRequest) {
	request = &UpdateDIProjectConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "UpdateDIProjectConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateDIProjectConfigResponse creates a response to parse from UpdateDIProjectConfig response
func CreateUpdateDIProjectConfigResponse() (response *UpdateDIProjectConfigResponse) {
	response = &UpdateDIProjectConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
