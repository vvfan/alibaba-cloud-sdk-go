package viapi_regen

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

// UpdateTrainTask invokes the viapi_regen.UpdateTrainTask API synchronously
func (client *Client) UpdateTrainTask(request *UpdateTrainTaskRequest) (response *UpdateTrainTaskResponse, err error) {
	response = CreateUpdateTrainTaskResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateTrainTaskWithChan invokes the viapi_regen.UpdateTrainTask API asynchronously
func (client *Client) UpdateTrainTaskWithChan(request *UpdateTrainTaskRequest) (<-chan *UpdateTrainTaskResponse, <-chan error) {
	responseChan := make(chan *UpdateTrainTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateTrainTask(request)
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

// UpdateTrainTaskWithCallback invokes the viapi_regen.UpdateTrainTask API asynchronously
func (client *Client) UpdateTrainTaskWithCallback(request *UpdateTrainTaskRequest, callback func(response *UpdateTrainTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateTrainTaskResponse
		var err error
		defer close(result)
		response, err = client.UpdateTrainTask(request)
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

// UpdateTrainTaskRequest is the request struct for api UpdateTrainTask
type UpdateTrainTaskRequest struct {
	*requests.RpcRequest
	Description        string           `position:"Body" name:"Description"`
	TrainMode          string           `position:"Body" name:"TrainMode"`
	Id                 requests.Integer `position:"Body" name:"Id"`
	PreTrainTaskId     requests.Integer `position:"Body" name:"PreTrainTaskId"`
	PreTrainTaskFlag   requests.Boolean `position:"Body" name:"PreTrainTaskFlag"`
	AdvancedParameters string           `position:"Body" name:"AdvancedParameters"`
	LabelId            requests.Integer `position:"Body" name:"LabelId"`
	Name               string           `position:"Body" name:"Name"`
	DatasetId          requests.Integer `position:"Body" name:"DatasetId"`
}

// UpdateTrainTaskResponse is the response struct for api UpdateTrainTask
type UpdateTrainTaskResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateUpdateTrainTaskRequest creates a request to invoke UpdateTrainTask API
func CreateUpdateTrainTaskRequest() (request *UpdateTrainTaskRequest) {
	request = &UpdateTrainTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "UpdateTrainTask", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateTrainTaskResponse creates a response to parse from UpdateTrainTask response
func CreateUpdateTrainTaskResponse() (response *UpdateTrainTaskResponse) {
	response = &UpdateTrainTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
