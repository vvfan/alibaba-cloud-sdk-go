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

// DescribeCheckEcsWarnings invokes the sas.DescribeCheckEcsWarnings API synchronously
func (client *Client) DescribeCheckEcsWarnings(request *DescribeCheckEcsWarningsRequest) (response *DescribeCheckEcsWarningsResponse, err error) {
	response = CreateDescribeCheckEcsWarningsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCheckEcsWarningsWithChan invokes the sas.DescribeCheckEcsWarnings API asynchronously
func (client *Client) DescribeCheckEcsWarningsWithChan(request *DescribeCheckEcsWarningsRequest) (<-chan *DescribeCheckEcsWarningsResponse, <-chan error) {
	responseChan := make(chan *DescribeCheckEcsWarningsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCheckEcsWarnings(request)
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

// DescribeCheckEcsWarningsWithCallback invokes the sas.DescribeCheckEcsWarnings API asynchronously
func (client *Client) DescribeCheckEcsWarningsWithCallback(request *DescribeCheckEcsWarningsRequest, callback func(response *DescribeCheckEcsWarningsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCheckEcsWarningsResponse
		var err error
		defer close(result)
		response, err = client.DescribeCheckEcsWarnings(request)
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

// DescribeCheckEcsWarningsRequest is the request struct for api DescribeCheckEcsWarnings
type DescribeCheckEcsWarningsRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeCheckEcsWarningsResponse is the response struct for api DescribeCheckEcsWarnings
type DescribeCheckEcsWarningsResponse struct {
	*responses.BaseResponse
	SasVersion        string `json:"SasVersion" xml:"SasVersion"`
	CanTry            string `json:"CanTry" xml:"CanTry"`
	WeakPasswordCount string `json:"WeakPasswordCount" xml:"WeakPasswordCount"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeCheckEcsWarningsRequest creates a request to invoke DescribeCheckEcsWarnings API
func CreateDescribeCheckEcsWarningsRequest() (request *DescribeCheckEcsWarningsRequest) {
	request = &DescribeCheckEcsWarningsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeCheckEcsWarnings", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCheckEcsWarningsResponse creates a response to parse from DescribeCheckEcsWarnings response
func CreateDescribeCheckEcsWarningsResponse() (response *DescribeCheckEcsWarningsResponse) {
	response = &DescribeCheckEcsWarningsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
