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

// CreateDISyncTask invokes the dataworks_public.CreateDISyncTask API synchronously
func (client *Client) CreateDISyncTask(request *CreateDISyncTaskRequest) (response *CreateDISyncTaskResponse, err error) {
	response = CreateCreateDISyncTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDISyncTaskWithChan invokes the dataworks_public.CreateDISyncTask API asynchronously
func (client *Client) CreateDISyncTaskWithChan(request *CreateDISyncTaskRequest) (<-chan *CreateDISyncTaskResponse, <-chan error) {
	responseChan := make(chan *CreateDISyncTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDISyncTask(request)
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

// CreateDISyncTaskWithCallback invokes the dataworks_public.CreateDISyncTask API asynchronously
func (client *Client) CreateDISyncTaskWithCallback(request *CreateDISyncTaskRequest, callback func(response *CreateDISyncTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDISyncTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateDISyncTask(request)
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

// CreateDISyncTaskRequest is the request struct for api CreateDISyncTask
type CreateDISyncTaskRequest struct {
	*requests.RpcRequest
	TaskType    string           `position:"Query" name:"TaskType"`
	ClientToken string           `position:"Query" name:"ClientToken"`
	TaskParam   string           `position:"Query" name:"TaskParam"`
	TaskName    string           `position:"Query" name:"TaskName"`
	TaskContent string           `position:"Query" name:"TaskContent"`
	ProjectId   requests.Integer `position:"Query" name:"ProjectId"`
}

// CreateDISyncTaskResponse is the response struct for api CreateDISyncTask
type CreateDISyncTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateDISyncTaskRequest creates a request to invoke CreateDISyncTask API
func CreateCreateDISyncTaskRequest() (request *CreateDISyncTaskRequest) {
	request = &CreateDISyncTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateDISyncTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDISyncTaskResponse creates a response to parse from CreateDISyncTask response
func CreateCreateDISyncTaskResponse() (response *CreateDISyncTaskResponse) {
	response = &CreateDISyncTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
