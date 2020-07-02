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

// UpdateConnection invokes the dataworks_public.UpdateConnection API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/updateconnection.html
func (client *Client) UpdateConnection(request *UpdateConnectionRequest) (response *UpdateConnectionResponse, err error) {
	response = CreateUpdateConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateConnectionWithChan invokes the dataworks_public.UpdateConnection API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/updateconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateConnectionWithChan(request *UpdateConnectionRequest) (<-chan *UpdateConnectionResponse, <-chan error) {
	responseChan := make(chan *UpdateConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateConnection(request)
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

// UpdateConnectionWithCallback invokes the dataworks_public.UpdateConnection API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/updateconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateConnectionWithCallback(request *UpdateConnectionRequest, callback func(response *UpdateConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateConnectionResponse
		var err error
		defer close(result)
		response, err = client.UpdateConnection(request)
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

// UpdateConnectionRequest is the request struct for api UpdateConnection
type UpdateConnectionRequest struct {
	*requests.RpcRequest
	EnvType      requests.Integer `position:"Query" name:"EnvType"`
	Description  string           `position:"Query" name:"Description"`
	ConnectionId requests.Integer `position:"Query" name:"ConnectionId"`
	Content      string           `position:"Query" name:"Content"`
	Status       string           `position:"Query" name:"Status"`
}

// UpdateConnectionResponse is the response struct for api UpdateConnection
type UpdateConnectionResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	HttpStatusCode string `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           bool   `json:"Data" xml:"Data"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateConnectionRequest creates a request to invoke UpdateConnection API
func CreateUpdateConnectionRequest() (request *UpdateConnectionRequest) {
	request = &UpdateConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "UpdateConnection", "dide", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateConnectionResponse creates a response to parse from UpdateConnection response
func CreateUpdateConnectionResponse() (response *UpdateConnectionResponse) {
	response = &UpdateConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
