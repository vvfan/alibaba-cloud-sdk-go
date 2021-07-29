package outboundbot

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

// GetSummaryInfo invokes the outboundbot.GetSummaryInfo API synchronously
func (client *Client) GetSummaryInfo(request *GetSummaryInfoRequest) (response *GetSummaryInfoResponse, err error) {
	response = CreateGetSummaryInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetSummaryInfoWithChan invokes the outboundbot.GetSummaryInfo API asynchronously
func (client *Client) GetSummaryInfoWithChan(request *GetSummaryInfoRequest) (<-chan *GetSummaryInfoResponse, <-chan error) {
	responseChan := make(chan *GetSummaryInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSummaryInfo(request)
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

// GetSummaryInfoWithCallback invokes the outboundbot.GetSummaryInfo API asynchronously
func (client *Client) GetSummaryInfoWithCallback(request *GetSummaryInfoRequest, callback func(response *GetSummaryInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSummaryInfoResponse
		var err error
		defer close(result)
		response, err = client.GetSummaryInfo(request)
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

// GetSummaryInfoRequest is the request struct for api GetSummaryInfo
type GetSummaryInfoRequest struct {
	*requests.RpcRequest
	InstanceIdList *[]string `position:"Query" name:"InstanceIdList"  type:"Repeated"`
}

// GetSummaryInfoResponse is the response struct for api GetSummaryInfo
type GetSummaryInfoResponse struct {
	*responses.BaseResponse
	RequestId                   string                    `json:"RequestId" xml:"RequestId"`
	Success                     bool                      `json:"Success" xml:"Success"`
	Code                        string                    `json:"Code" xml:"Code"`
	Message                     string                    `json:"Message" xml:"Message"`
	HttpStatusCode              int                       `json:"HttpStatusCode" xml:"HttpStatusCode"`
	AgentBotInstanceSummaryList []AgentBotInstanceSummary `json:"AgentBotInstanceSummaryList" xml:"AgentBotInstanceSummaryList"`
}

// CreateGetSummaryInfoRequest creates a request to invoke GetSummaryInfo API
func CreateGetSummaryInfoRequest() (request *GetSummaryInfoRequest) {
	request = &GetSummaryInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "GetSummaryInfo", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSummaryInfoResponse creates a response to parse from GetSummaryInfo response
func CreateGetSummaryInfoResponse() (response *GetSummaryInfoResponse) {
	response = &GetSummaryInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}