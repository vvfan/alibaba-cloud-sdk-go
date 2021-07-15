package alb

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

// ListListenerCertificates invokes the alb.ListListenerCertificates API synchronously
func (client *Client) ListListenerCertificates(request *ListListenerCertificatesRequest) (response *ListListenerCertificatesResponse, err error) {
	response = CreateListListenerCertificatesResponse()
	err = client.DoAction(request, response)
	return
}

// ListListenerCertificatesWithChan invokes the alb.ListListenerCertificates API asynchronously
func (client *Client) ListListenerCertificatesWithChan(request *ListListenerCertificatesRequest) (<-chan *ListListenerCertificatesResponse, <-chan error) {
	responseChan := make(chan *ListListenerCertificatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListListenerCertificates(request)
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

// ListListenerCertificatesWithCallback invokes the alb.ListListenerCertificates API asynchronously
func (client *Client) ListListenerCertificatesWithCallback(request *ListListenerCertificatesRequest, callback func(response *ListListenerCertificatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListListenerCertificatesResponse
		var err error
		defer close(result)
		response, err = client.ListListenerCertificates(request)
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

// ListListenerCertificatesRequest is the request struct for api ListListenerCertificates
type ListListenerCertificatesRequest struct {
	*requests.RpcRequest
	CertificateType string           `position:"Query" name:"CertificateType"`
	ListenerId      string           `position:"Query" name:"ListenerId"`
	NextToken       string           `position:"Query" name:"NextToken"`
	MaxResults      requests.Integer `position:"Query" name:"MaxResults"`
}

// ListListenerCertificatesResponse is the response struct for api ListListenerCertificates
type ListListenerCertificatesResponse struct {
	*responses.BaseResponse
	MaxResults   int                `json:"MaxResults" xml:"MaxResults"`
	NextToken    string             `json:"NextToken" xml:"NextToken"`
	RequestId    string             `json:"RequestId" xml:"RequestId"`
	TotalCount   int                `json:"TotalCount" xml:"TotalCount"`
	Certificates []CertificateModel `json:"Certificates" xml:"Certificates"`
}

// CreateListListenerCertificatesRequest creates a request to invoke ListListenerCertificates API
func CreateListListenerCertificatesRequest() (request *ListListenerCertificatesRequest) {
	request = &ListListenerCertificatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alb", "2020-06-16", "ListListenerCertificates", "alb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListListenerCertificatesResponse creates a response to parse from ListListenerCertificates response
func CreateListListenerCertificatesResponse() (response *ListListenerCertificatesResponse) {
	response = &ListListenerCertificatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
