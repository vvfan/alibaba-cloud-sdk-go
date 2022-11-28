package domain

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

// TransferInRefetchWhoisEmail invokes the domain.TransferInRefetchWhoisEmail API synchronously
func (client *Client) TransferInRefetchWhoisEmail(request *TransferInRefetchWhoisEmailRequest) (response *TransferInRefetchWhoisEmailResponse, err error) {
	response = CreateTransferInRefetchWhoisEmailResponse()
	err = client.DoAction(request, response)
	return
}

// TransferInRefetchWhoisEmailWithChan invokes the domain.TransferInRefetchWhoisEmail API asynchronously
func (client *Client) TransferInRefetchWhoisEmailWithChan(request *TransferInRefetchWhoisEmailRequest) (<-chan *TransferInRefetchWhoisEmailResponse, <-chan error) {
	responseChan := make(chan *TransferInRefetchWhoisEmailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TransferInRefetchWhoisEmail(request)
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

// TransferInRefetchWhoisEmailWithCallback invokes the domain.TransferInRefetchWhoisEmail API asynchronously
func (client *Client) TransferInRefetchWhoisEmailWithCallback(request *TransferInRefetchWhoisEmailRequest, callback func(response *TransferInRefetchWhoisEmailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TransferInRefetchWhoisEmailResponse
		var err error
		defer close(result)
		response, err = client.TransferInRefetchWhoisEmail(request)
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

// TransferInRefetchWhoisEmailRequest is the request struct for api TransferInRefetchWhoisEmail
type TransferInRefetchWhoisEmailRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// TransferInRefetchWhoisEmailResponse is the response struct for api TransferInRefetchWhoisEmail
type TransferInRefetchWhoisEmailResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateTransferInRefetchWhoisEmailRequest creates a request to invoke TransferInRefetchWhoisEmail API
func CreateTransferInRefetchWhoisEmailRequest() (request *TransferInRefetchWhoisEmailRequest) {
	request = &TransferInRefetchWhoisEmailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "TransferInRefetchWhoisEmail", "", "")
	request.Method = requests.POST
	return
}

// CreateTransferInRefetchWhoisEmailResponse creates a response to parse from TransferInRefetchWhoisEmail response
func CreateTransferInRefetchWhoisEmailResponse() (response *TransferInRefetchWhoisEmailResponse) {
	response = &TransferInRefetchWhoisEmailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
