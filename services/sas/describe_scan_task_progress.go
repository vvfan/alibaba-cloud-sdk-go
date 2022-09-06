package sas

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

// DescribeScanTaskProgress invokes the sas.DescribeScanTaskProgress API synchronously
func (client *Client) DescribeScanTaskProgress(request *DescribeScanTaskProgressRequest) (response *DescribeScanTaskProgressResponse, err error) {
	response = CreateDescribeScanTaskProgressResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScanTaskProgressWithChan invokes the sas.DescribeScanTaskProgress API asynchronously
func (client *Client) DescribeScanTaskProgressWithChan(request *DescribeScanTaskProgressRequest) (<-chan *DescribeScanTaskProgressResponse, <-chan error) {
	responseChan := make(chan *DescribeScanTaskProgressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScanTaskProgress(request)
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

// DescribeScanTaskProgressWithCallback invokes the sas.DescribeScanTaskProgress API asynchronously
func (client *Client) DescribeScanTaskProgressWithCallback(request *DescribeScanTaskProgressRequest, callback func(response *DescribeScanTaskProgressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScanTaskProgressResponse
		var err error
		defer close(result)
		response, err = client.DescribeScanTaskProgress(request)
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

// DescribeScanTaskProgressRequest is the request struct for api DescribeScanTaskProgress
type DescribeScanTaskProgressRequest struct {
	*requests.RpcRequest
	SourceIp string           `position:"Query" name:"SourceIp"`
	TaskId   requests.Integer `position:"Query" name:"TaskId"`
}

// DescribeScanTaskProgressResponse is the response struct for api DescribeScanTaskProgress
type DescribeScanTaskProgressResponse struct {
	*responses.BaseResponse
	TargetInfo       string `json:"TargetInfo" xml:"TargetInfo"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	ScanTaskProgress string `json:"ScanTaskProgress" xml:"ScanTaskProgress"`
}

// CreateDescribeScanTaskProgressRequest creates a request to invoke DescribeScanTaskProgress API
func CreateDescribeScanTaskProgressRequest() (request *DescribeScanTaskProgressRequest) {
	request = &DescribeScanTaskProgressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeScanTaskProgress", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeScanTaskProgressResponse creates a response to parse from DescribeScanTaskProgress response
func CreateDescribeScanTaskProgressResponse() (response *DescribeScanTaskProgressResponse) {
	response = &DescribeScanTaskProgressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
