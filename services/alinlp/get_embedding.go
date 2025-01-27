package alinlp

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

// GetEmbedding invokes the alinlp.GetEmbedding API synchronously
func (client *Client) GetEmbedding(request *GetEmbeddingRequest) (response *GetEmbeddingResponse, err error) {
	response = CreateGetEmbeddingResponse()
	err = client.DoAction(request, response)
	return
}

// GetEmbeddingWithChan invokes the alinlp.GetEmbedding API asynchronously
func (client *Client) GetEmbeddingWithChan(request *GetEmbeddingRequest) (<-chan *GetEmbeddingResponse, <-chan error) {
	responseChan := make(chan *GetEmbeddingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetEmbedding(request)
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

// GetEmbeddingWithCallback invokes the alinlp.GetEmbedding API asynchronously
func (client *Client) GetEmbeddingWithCallback(request *GetEmbeddingRequest, callback func(response *GetEmbeddingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetEmbeddingResponse
		var err error
		defer close(result)
		response, err = client.GetEmbedding(request)
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

// GetEmbeddingRequest is the request struct for api GetEmbedding
type GetEmbeddingRequest struct {
	*requests.RpcRequest
	Business    string `position:"Query" name:"Business"`
	ServiceCode string `position:"Body" name:"ServiceCode"`
	Text        string `position:"Body" name:"Text"`
}

// GetEmbeddingResponse is the response struct for api GetEmbedding
type GetEmbeddingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetEmbeddingRequest creates a request to invoke GetEmbedding API
func CreateGetEmbeddingRequest() (request *GetEmbeddingRequest) {
	request = &GetEmbeddingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alinlp", "2020-06-29", "GetEmbedding", "alinlp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetEmbeddingResponse creates a response to parse from GetEmbedding response
func CreateGetEmbeddingResponse() (response *GetEmbeddingResponse) {
	response = &GetEmbeddingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
