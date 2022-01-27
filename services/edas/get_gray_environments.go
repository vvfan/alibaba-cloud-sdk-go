package edas

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

// GetGrayEnvironments invokes the edas.GetGrayEnvironments API synchronously
func (client *Client) GetGrayEnvironments(request *GetGrayEnvironmentsRequest) (response *GetGrayEnvironmentsResponse, err error) {
	response = CreateGetGrayEnvironmentsResponse()
	err = client.DoAction(request, response)
	return
}

// GetGrayEnvironmentsWithChan invokes the edas.GetGrayEnvironments API asynchronously
func (client *Client) GetGrayEnvironmentsWithChan(request *GetGrayEnvironmentsRequest) (<-chan *GetGrayEnvironmentsResponse, <-chan error) {
	responseChan := make(chan *GetGrayEnvironmentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGrayEnvironments(request)
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

// GetGrayEnvironmentsWithCallback invokes the edas.GetGrayEnvironments API asynchronously
func (client *Client) GetGrayEnvironmentsWithCallback(request *GetGrayEnvironmentsRequest, callback func(response *GetGrayEnvironmentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetGrayEnvironmentsResponse
		var err error
		defer close(result)
		response, err = client.GetGrayEnvironments(request)
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

// GetGrayEnvironmentsRequest is the request struct for api GetGrayEnvironments
type GetGrayEnvironmentsRequest struct {
	*requests.RoaRequest
	LogicalRegionId string `position:"Query" name:"LogicalRegionId"`
}

// GetGrayEnvironmentsResponse is the response struct for api GetGrayEnvironments
type GetGrayEnvironmentsResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Code      int       `json:"Code" xml:"Code"`
	Message   string    `json:"Message" xml:"Message"`
	Data      []EnvList `json:"Data" xml:"Data"`
}

// CreateGetGrayEnvironmentsRequest creates a request to invoke GetGrayEnvironments API
func CreateGetGrayEnvironmentsRequest() (request *GetGrayEnvironmentsRequest) {
	request = &GetGrayEnvironmentsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "GetGrayEnvironments", "/pop/v5/gray/env_list", "Edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetGrayEnvironmentsResponse creates a response to parse from GetGrayEnvironments response
func CreateGetGrayEnvironmentsResponse() (response *GetGrayEnvironmentsResponse) {
	response = &GetGrayEnvironmentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
