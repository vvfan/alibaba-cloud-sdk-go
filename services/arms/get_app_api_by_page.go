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

// GetAppApiByPage invokes the arms.GetAppApiByPage API synchronously
func (client *Client) GetAppApiByPage(request *GetAppApiByPageRequest) (response *GetAppApiByPageResponse, err error) {
	response = CreateGetAppApiByPageResponse()
	err = client.DoAction(request, response)
	return
}

// GetAppApiByPageWithChan invokes the arms.GetAppApiByPage API asynchronously
func (client *Client) GetAppApiByPageWithChan(request *GetAppApiByPageRequest) (<-chan *GetAppApiByPageResponse, <-chan error) {
	responseChan := make(chan *GetAppApiByPageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAppApiByPage(request)
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

// GetAppApiByPageWithCallback invokes the arms.GetAppApiByPage API asynchronously
func (client *Client) GetAppApiByPageWithCallback(request *GetAppApiByPageRequest, callback func(response *GetAppApiByPageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAppApiByPageResponse
		var err error
		defer close(result)
		response, err = client.GetAppApiByPage(request)
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

// GetAppApiByPageRequest is the request struct for api GetAppApiByPage
type GetAppApiByPageRequest struct {
	*requests.RpcRequest
	EndTime       requests.Integer `position:"Query" name:"EndTime"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	PId           string           `position:"Query" name:"PId"`
	StartTime     requests.Integer `position:"Query" name:"StartTime"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	IntervalMills requests.Integer `position:"Query" name:"IntervalMills"`
}

// GetAppApiByPageResponse is the response struct for api GetAppApiByPage
type GetAppApiByPageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetAppApiByPageRequest creates a request to invoke GetAppApiByPage API
func CreateGetAppApiByPageRequest() (request *GetAppApiByPageRequest) {
	request = &GetAppApiByPageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "GetAppApiByPage", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAppApiByPageResponse creates a response to parse from GetAppApiByPage response
func CreateGetAppApiByPageResponse() (response *GetAppApiByPageResponse) {
	response = &GetAppApiByPageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
