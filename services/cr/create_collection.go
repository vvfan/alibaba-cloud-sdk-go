package cr

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

// CreateCollection invokes the cr.CreateCollection API synchronously
// api document: https://help.aliyun.com/api/cr/createcollection.html
func (client *Client) CreateCollection(request *CreateCollectionRequest) (response *CreateCollectionResponse, err error) {
	response = CreateCreateCollectionResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCollectionWithChan invokes the cr.CreateCollection API asynchronously
// api document: https://help.aliyun.com/api/cr/createcollection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCollectionWithChan(request *CreateCollectionRequest) (<-chan *CreateCollectionResponse, <-chan error) {
	responseChan := make(chan *CreateCollectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCollection(request)
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

// CreateCollectionWithCallback invokes the cr.CreateCollection API asynchronously
// api document: https://help.aliyun.com/api/cr/createcollection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCollectionWithCallback(request *CreateCollectionRequest, callback func(response *CreateCollectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCollectionResponse
		var err error
		defer close(result)
		response, err = client.CreateCollection(request)
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

// CreateCollectionRequest is the request struct for api CreateCollection
type CreateCollectionRequest struct {
	*requests.RoaRequest
}

// CreateCollectionResponse is the response struct for api CreateCollection
type CreateCollectionResponse struct {
	*responses.BaseResponse
}

// CreateCreateCollectionRequest creates a request to invoke CreateCollection API
func CreateCreateCollectionRequest() (request *CreateCollectionRequest) {
	request = &CreateCollectionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "CreateCollection", "/collections", "", "")
	request.Method = requests.PUT
	return
}

// CreateCreateCollectionResponse creates a response to parse from CreateCollection response
func CreateCreateCollectionResponse() (response *CreateCollectionResponse) {
	response = &CreateCollectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
