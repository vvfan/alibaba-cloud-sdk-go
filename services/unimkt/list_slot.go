package unimkt

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

// ListSlot invokes the unimkt.ListSlot API synchronously
func (client *Client) ListSlot(request *ListSlotRequest) (response *ListSlotResponse, err error) {
	response = CreateListSlotResponse()
	err = client.DoAction(request, response)
	return
}

// ListSlotWithChan invokes the unimkt.ListSlot API asynchronously
func (client *Client) ListSlotWithChan(request *ListSlotRequest) (<-chan *ListSlotResponse, <-chan error) {
	responseChan := make(chan *ListSlotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSlot(request)
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

// ListSlotWithCallback invokes the unimkt.ListSlot API asynchronously
func (client *Client) ListSlotWithCallback(request *ListSlotRequest, callback func(response *ListSlotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSlotResponse
		var err error
		defer close(result)
		response, err = client.ListSlot(request)
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

// ListSlotRequest is the request struct for api ListSlot
type ListSlotRequest struct {
	*requests.RpcRequest
	AdSlotType            string           `position:"Query" name:"AdSlotType"`
	UserId                string           `position:"Query" name:"UserId"`
	OriginSiteUserId      string           `position:"Query" name:"OriginSiteUserId"`
	PageNumber            requests.Integer `position:"Query" name:"PageNumber"`
	MediaName             string           `position:"Query" name:"MediaName"`
	AppName               string           `position:"Query" name:"AppName"`
	AdSlotStatus          string           `position:"Query" name:"AdSlotStatus"`
	TenantId              string           `position:"Query" name:"TenantId"`
	AdSlotId              string           `position:"Query" name:"AdSlotId"`
	PageSize              requests.Integer `position:"Query" name:"PageSize"`
	AdSlotCorporateStatus string           `position:"Query" name:"AdSlotCorporateStatus"`
	EndCreateTime         requests.Integer `position:"Query" name:"EndCreateTime"`
	Business              string           `position:"Query" name:"Business"`
	MediaId               string           `position:"Query" name:"MediaId"`
	Environment           string           `position:"Query" name:"Environment"`
	StartCreateTime       requests.Integer `position:"Query" name:"StartCreateTime"`
	UserSite              string           `position:"Query" name:"UserSite"`
	AdSlotName            string           `position:"Query" name:"AdSlotName"`
}

// ListSlotResponse is the response struct for api ListSlot
type ListSlotResponse struct {
	*responses.BaseResponse
	Code       string   `json:"Code" xml:"Code"`
	Success    bool     `json:"Success" xml:"Success"`
	Message    string   `json:"Message" xml:"Message"`
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	PageNumber int      `json:"PageNumber" xml:"PageNumber"`
	PageSize   int      `json:"PageSize" xml:"PageSize"`
	Total      int64    `json:"Total" xml:"Total"`
	Model      []AdSlot `json:"Model" xml:"Model"`
}

// CreateListSlotRequest creates a request to invoke ListSlot API
func CreateListSlotRequest() (request *ListSlotRequest) {
	request = &ListSlotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "ListSlot", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListSlotResponse creates a response to parse from ListSlot response
func CreateListSlotResponse() (response *ListSlotResponse) {
	response = &ListSlotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}