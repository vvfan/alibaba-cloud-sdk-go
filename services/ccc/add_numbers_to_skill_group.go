package ccc

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

// AddNumbersToSkillGroup invokes the ccc.AddNumbersToSkillGroup API synchronously
func (client *Client) AddNumbersToSkillGroup(request *AddNumbersToSkillGroupRequest) (response *AddNumbersToSkillGroupResponse, err error) {
	response = CreateAddNumbersToSkillGroupResponse()
	err = client.DoAction(request, response)
	return
}

// AddNumbersToSkillGroupWithChan invokes the ccc.AddNumbersToSkillGroup API asynchronously
func (client *Client) AddNumbersToSkillGroupWithChan(request *AddNumbersToSkillGroupRequest) (<-chan *AddNumbersToSkillGroupResponse, <-chan error) {
	responseChan := make(chan *AddNumbersToSkillGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddNumbersToSkillGroup(request)
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

// AddNumbersToSkillGroupWithCallback invokes the ccc.AddNumbersToSkillGroup API asynchronously
func (client *Client) AddNumbersToSkillGroupWithCallback(request *AddNumbersToSkillGroupRequest, callback func(response *AddNumbersToSkillGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddNumbersToSkillGroupResponse
		var err error
		defer close(result)
		response, err = client.AddNumbersToSkillGroup(request)
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

// AddNumbersToSkillGroupRequest is the request struct for api AddNumbersToSkillGroup
type AddNumbersToSkillGroupRequest struct {
	*requests.RpcRequest
	NumberList   string `position:"Query" name:"NumberList"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	SkillGroupId string `position:"Query" name:"SkillGroupId"`
}

// AddNumbersToSkillGroupResponse is the response struct for api AddNumbersToSkillGroup
type AddNumbersToSkillGroupResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateAddNumbersToSkillGroupRequest creates a request to invoke AddNumbersToSkillGroup API
func CreateAddNumbersToSkillGroupRequest() (request *AddNumbersToSkillGroupRequest) {
	request = &AddNumbersToSkillGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "AddNumbersToSkillGroup", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddNumbersToSkillGroupResponse creates a response to parse from AddNumbersToSkillGroup response
func CreateAddNumbersToSkillGroupResponse() (response *AddNumbersToSkillGroupResponse) {
	response = &AddNumbersToSkillGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}