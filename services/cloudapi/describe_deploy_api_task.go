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

// DescribeDeployApiTask invokes the cloudapi.DescribeDeployApiTask API synchronously
func (client *Client) DescribeDeployApiTask(request *DescribeDeployApiTaskRequest) (response *DescribeDeployApiTaskResponse, err error) {
	response = CreateDescribeDeployApiTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDeployApiTaskWithChan invokes the cloudapi.DescribeDeployApiTask API asynchronously
func (client *Client) DescribeDeployApiTaskWithChan(request *DescribeDeployApiTaskRequest) (<-chan *DescribeDeployApiTaskResponse, <-chan error) {
	responseChan := make(chan *DescribeDeployApiTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDeployApiTask(request)
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

// DescribeDeployApiTaskWithCallback invokes the cloudapi.DescribeDeployApiTask API asynchronously
func (client *Client) DescribeDeployApiTaskWithCallback(request *DescribeDeployApiTaskRequest, callback func(response *DescribeDeployApiTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDeployApiTaskResponse
		var err error
		defer close(result)
		response, err = client.DescribeDeployApiTask(request)
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

// DescribeDeployApiTaskRequest is the request struct for api DescribeDeployApiTask
type DescribeDeployApiTaskRequest struct {
	*requests.RpcRequest
	OperationUid  string `position:"Query" name:"OperationUid"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// DescribeDeployApiTaskResponse is the response struct for api DescribeDeployApiTask
type DescribeDeployApiTaskResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	DeployedResults DeployedResults `json:"DeployedResults" xml:"DeployedResults"`
}

// CreateDescribeDeployApiTaskRequest creates a request to invoke DescribeDeployApiTask API
func CreateDescribeDeployApiTaskRequest() (request *DescribeDeployApiTaskRequest) {
	request = &DescribeDeployApiTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeDeployApiTask", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDeployApiTaskResponse creates a response to parse from DescribeDeployApiTask response
func CreateDescribeDeployApiTaskResponse() (response *DescribeDeployApiTaskResponse) {
	response = &DescribeDeployApiTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
