package cloudauth

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

// DescribeFaceConfig invokes the cloudauth.DescribeFaceConfig API synchronously
func (client *Client) DescribeFaceConfig(request *DescribeFaceConfigRequest) (response *DescribeFaceConfigResponse, err error) {
	response = CreateDescribeFaceConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFaceConfigWithChan invokes the cloudauth.DescribeFaceConfig API asynchronously
func (client *Client) DescribeFaceConfigWithChan(request *DescribeFaceConfigRequest) (<-chan *DescribeFaceConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeFaceConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFaceConfig(request)
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

// DescribeFaceConfigWithCallback invokes the cloudauth.DescribeFaceConfig API asynchronously
func (client *Client) DescribeFaceConfigWithCallback(request *DescribeFaceConfigRequest, callback func(response *DescribeFaceConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFaceConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeFaceConfig(request)
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

// DescribeFaceConfigRequest is the request struct for api DescribeFaceConfig
type DescribeFaceConfigRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Lang     string `position:"Query" name:"Lang"`
}

// DescribeFaceConfigResponse is the response struct for api DescribeFaceConfig
type DescribeFaceConfigResponse struct {
	*responses.BaseResponse
	RequestId string      `json:"RequestId" xml:"RequestId"`
	Items     []ItemsItem `json:"Items" xml:"Items"`
}

// CreateDescribeFaceConfigRequest creates a request to invoke DescribeFaceConfig API
func CreateDescribeFaceConfigRequest() (request *DescribeFaceConfigRequest) {
	request = &DescribeFaceConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "DescribeFaceConfig", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeFaceConfigResponse creates a response to parse from DescribeFaceConfig response
func CreateDescribeFaceConfigResponse() (response *DescribeFaceConfigResponse) {
	response = &DescribeFaceConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}