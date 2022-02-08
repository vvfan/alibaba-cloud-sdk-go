package arms

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

// GetPrometheusApiToken invokes the arms.GetPrometheusApiToken API synchronously
func (client *Client) GetPrometheusApiToken(request *GetPrometheusApiTokenRequest) (response *GetPrometheusApiTokenResponse, err error) {
	response = CreateGetPrometheusApiTokenResponse()
	err = client.DoAction(request, response)
	return
}

// GetPrometheusApiTokenWithChan invokes the arms.GetPrometheusApiToken API asynchronously
func (client *Client) GetPrometheusApiTokenWithChan(request *GetPrometheusApiTokenRequest) (<-chan *GetPrometheusApiTokenResponse, <-chan error) {
	responseChan := make(chan *GetPrometheusApiTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPrometheusApiToken(request)
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

// GetPrometheusApiTokenWithCallback invokes the arms.GetPrometheusApiToken API asynchronously
func (client *Client) GetPrometheusApiTokenWithCallback(request *GetPrometheusApiTokenRequest, callback func(response *GetPrometheusApiTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPrometheusApiTokenResponse
		var err error
		defer close(result)
		response, err = client.GetPrometheusApiToken(request)
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

// GetPrometheusApiTokenRequest is the request struct for api GetPrometheusApiToken
type GetPrometheusApiTokenRequest struct {
	*requests.RpcRequest
}

// GetPrometheusApiTokenResponse is the response struct for api GetPrometheusApiToken
type GetPrometheusApiTokenResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Token     string `json:"Token" xml:"Token"`
}

// CreateGetPrometheusApiTokenRequest creates a request to invoke GetPrometheusApiToken API
func CreateGetPrometheusApiTokenRequest() (request *GetPrometheusApiTokenRequest) {
	request = &GetPrometheusApiTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "GetPrometheusApiToken", "", "")
	request.Method = requests.POST
	return
}

// CreateGetPrometheusApiTokenResponse creates a response to parse from GetPrometheusApiToken response
func CreateGetPrometheusApiTokenResponse() (response *GetPrometheusApiTokenResponse) {
	response = &GetPrometheusApiTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
