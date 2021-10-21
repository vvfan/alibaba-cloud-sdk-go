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

// GetDISyncInstanceInfo invokes the dataworks_public.GetDISyncInstanceInfo API synchronously
func (client *Client) GetDISyncInstanceInfo(request *GetDISyncInstanceInfoRequest) (response *GetDISyncInstanceInfoResponse, err error) {
	response = CreateGetDISyncInstanceInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetDISyncInstanceInfoWithChan invokes the dataworks_public.GetDISyncInstanceInfo API asynchronously
func (client *Client) GetDISyncInstanceInfoWithChan(request *GetDISyncInstanceInfoRequest) (<-chan *GetDISyncInstanceInfoResponse, <-chan error) {
	responseChan := make(chan *GetDISyncInstanceInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDISyncInstanceInfo(request)
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

// GetDISyncInstanceInfoWithCallback invokes the dataworks_public.GetDISyncInstanceInfo API asynchronously
func (client *Client) GetDISyncInstanceInfoWithCallback(request *GetDISyncInstanceInfoRequest, callback func(response *GetDISyncInstanceInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDISyncInstanceInfoResponse
		var err error
		defer close(result)
		response, err = client.GetDISyncInstanceInfo(request)
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

// GetDISyncInstanceInfoRequest is the request struct for api GetDISyncInstanceInfo
type GetDISyncInstanceInfoRequest struct {
	*requests.RpcRequest
	TaskType  string           `position:"Query" name:"TaskType"`
	ProjectId requests.Integer `position:"Query" name:"ProjectId"`
	FileId    requests.Integer `position:"Query" name:"FileId"`
}

// GetDISyncInstanceInfoResponse is the response struct for api GetDISyncInstanceInfo
type GetDISyncInstanceInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetDISyncInstanceInfoRequest creates a request to invoke GetDISyncInstanceInfo API
func CreateGetDISyncInstanceInfoRequest() (request *GetDISyncInstanceInfoRequest) {
	request = &GetDISyncInstanceInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetDISyncInstanceInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateGetDISyncInstanceInfoResponse creates a response to parse from GetDISyncInstanceInfo response
func CreateGetDISyncInstanceInfoResponse() (response *GetDISyncInstanceInfoResponse) {
	response = &GetDISyncInstanceInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
