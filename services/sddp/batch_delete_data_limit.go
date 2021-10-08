package sddp

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

// BatchDeleteDataLimit invokes the sddp.BatchDeleteDataLimit API synchronously
func (client *Client) BatchDeleteDataLimit(request *BatchDeleteDataLimitRequest) (response *BatchDeleteDataLimitResponse, err error) {
	response = CreateBatchDeleteDataLimitResponse()
	err = client.DoAction(request, response)
	return
}

// BatchDeleteDataLimitWithChan invokes the sddp.BatchDeleteDataLimit API asynchronously
func (client *Client) BatchDeleteDataLimitWithChan(request *BatchDeleteDataLimitRequest) (<-chan *BatchDeleteDataLimitResponse, <-chan error) {
	responseChan := make(chan *BatchDeleteDataLimitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchDeleteDataLimit(request)
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

// BatchDeleteDataLimitWithCallback invokes the sddp.BatchDeleteDataLimit API asynchronously
func (client *Client) BatchDeleteDataLimitWithCallback(request *BatchDeleteDataLimitRequest, callback func(response *BatchDeleteDataLimitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchDeleteDataLimitResponse
		var err error
		defer close(result)
		response, err = client.BatchDeleteDataLimit(request)
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

// BatchDeleteDataLimitRequest is the request struct for api BatchDeleteDataLimit
type BatchDeleteDataLimitRequest struct {
	*requests.RpcRequest
	ResourceType  requests.Integer `position:"Query" name:"ResourceType"`
	SourceIp      string           `position:"Query" name:"SourceIp"`
	DataLimitList string           `position:"Query" name:"DataLimitList"`
}

// BatchDeleteDataLimitResponse is the response struct for api BatchDeleteDataLimit
type BatchDeleteDataLimitResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchDeleteDataLimitRequest creates a request to invoke BatchDeleteDataLimit API
func CreateBatchDeleteDataLimitRequest() (request *BatchDeleteDataLimitRequest) {
	request = &BatchDeleteDataLimitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "BatchDeleteDataLimit", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchDeleteDataLimitResponse creates a response to parse from BatchDeleteDataLimit response
func CreateBatchDeleteDataLimitResponse() (response *BatchDeleteDataLimitResponse) {
	response = &BatchDeleteDataLimitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}