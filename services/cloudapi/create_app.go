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

// CreateApp invokes the cloudapi.CreateApp API synchronously
func (client *Client) CreateApp(request *CreateAppRequest) (response *CreateAppResponse, err error) {
	response = CreateCreateAppResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAppWithChan invokes the cloudapi.CreateApp API asynchronously
func (client *Client) CreateAppWithChan(request *CreateAppRequest) (<-chan *CreateAppResponse, <-chan error) {
	responseChan := make(chan *CreateAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateApp(request)
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

// CreateAppWithCallback invokes the cloudapi.CreateApp API asynchronously
func (client *Client) CreateAppWithCallback(request *CreateAppRequest, callback func(response *CreateAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAppResponse
		var err error
		defer close(result)
		response, err = client.CreateApp(request)
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

// CreateAppRequest is the request struct for api CreateApp
type CreateAppRequest struct {
	*requests.RpcRequest
	Description   string          `position:"Query" name:"Description"`
	Source        string          `position:"Query" name:"Source"`
	AppName       string          `position:"Query" name:"AppName"`
	SecurityToken string          `position:"Query" name:"SecurityToken"`
	AppSecret     string          `position:"Query" name:"AppSecret"`
	AppKey        string          `position:"Query" name:"AppKey"`
	AppCode       string          `position:"Query" name:"AppCode"`
	Tag           *[]CreateAppTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// CreateAppTag is a repeated param struct in CreateAppRequest
type CreateAppTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateAppResponse is the response struct for api CreateApp
type CreateAppResponse struct {
	*responses.BaseResponse
	TagStatus bool   `json:"TagStatus" xml:"TagStatus"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	AppId     int64  `json:"AppId" xml:"AppId"`
}

// CreateCreateAppRequest creates a request to invoke CreateApp API
func CreateCreateAppRequest() (request *CreateAppRequest) {
	request = &CreateAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateApp", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAppResponse creates a response to parse from CreateApp response
func CreateCreateAppResponse() (response *CreateAppResponse) {
	response = &CreateAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
