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

// ListWorkspaces invokes the viapi_regen.ListWorkspaces API synchronously
func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (response *ListWorkspacesResponse, err error) {
	response = CreateListWorkspacesResponse()
	err = client.DoAction(request, response)
	return
}

// ListWorkspacesWithChan invokes the viapi_regen.ListWorkspaces API asynchronously
func (client *Client) ListWorkspacesWithChan(request *ListWorkspacesRequest) (<-chan *ListWorkspacesResponse, <-chan error) {
	responseChan := make(chan *ListWorkspacesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListWorkspaces(request)
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

// ListWorkspacesWithCallback invokes the viapi_regen.ListWorkspaces API asynchronously
func (client *Client) ListWorkspacesWithCallback(request *ListWorkspacesRequest, callback func(response *ListWorkspacesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListWorkspacesResponse
		var err error
		defer close(result)
		response, err = client.ListWorkspaces(request)
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

// ListWorkspacesRequest is the request struct for api ListWorkspaces
type ListWorkspacesRequest struct {
	*requests.RpcRequest
	PageSize    requests.Integer `position:"Body" name:"PageSize"`
	CurrentPage requests.Integer `position:"Body" name:"CurrentPage"`
	Name        string           `position:"Body" name:"Name"`
}

// ListWorkspacesResponse is the response struct for api ListWorkspaces
type ListWorkspacesResponse struct {
	*responses.BaseResponse
	Message   string               `json:"Message" xml:"Message"`
	RequestId string               `json:"RequestId" xml:"RequestId"`
	Code      string               `json:"Code" xml:"Code"`
	Data      DataInListWorkspaces `json:"Data" xml:"Data"`
}

// CreateListWorkspacesRequest creates a request to invoke ListWorkspaces API
func CreateListWorkspacesRequest() (request *ListWorkspacesRequest) {
	request = &ListWorkspacesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "ListWorkspaces", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListWorkspacesResponse creates a response to parse from ListWorkspaces response
func CreateListWorkspacesResponse() (response *ListWorkspacesResponse) {
	response = &ListWorkspacesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
