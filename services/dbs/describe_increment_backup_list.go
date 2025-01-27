package dbs

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

// DescribeIncrementBackupList invokes the dbs.DescribeIncrementBackupList API synchronously
func (client *Client) DescribeIncrementBackupList(request *DescribeIncrementBackupListRequest) (response *DescribeIncrementBackupListResponse, err error) {
	response = CreateDescribeIncrementBackupListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeIncrementBackupListWithChan invokes the dbs.DescribeIncrementBackupList API asynchronously
func (client *Client) DescribeIncrementBackupListWithChan(request *DescribeIncrementBackupListRequest) (<-chan *DescribeIncrementBackupListResponse, <-chan error) {
	responseChan := make(chan *DescribeIncrementBackupListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeIncrementBackupList(request)
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

// DescribeIncrementBackupListWithCallback invokes the dbs.DescribeIncrementBackupList API asynchronously
func (client *Client) DescribeIncrementBackupListWithCallback(request *DescribeIncrementBackupListRequest, callback func(response *DescribeIncrementBackupListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeIncrementBackupListResponse
		var err error
		defer close(result)
		response, err = client.DescribeIncrementBackupList(request)
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

// DescribeIncrementBackupListRequest is the request struct for api DescribeIncrementBackupList
type DescribeIncrementBackupListRequest struct {
	*requests.RpcRequest
	ClientToken     string           `position:"Query" name:"ClientToken"`
	BackupPlanId    string           `position:"Query" name:"BackupPlanId"`
	PageNum         requests.Integer `position:"Query" name:"PageNum"`
	OwnerId         string           `position:"Query" name:"OwnerId"`
	StartTimestamp  requests.Integer `position:"Query" name:"StartTimestamp"`
	EndTimestamp    requests.Integer `position:"Query" name:"EndTimestamp"`
	ShowStorageType requests.Boolean `position:"Query" name:"ShowStorageType"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeIncrementBackupListResponse is the response struct for api DescribeIncrementBackupList
type DescribeIncrementBackupListResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PageNum        int                                `json:"PageNum" xml:"PageNum"`
	RequestId      string                             `json:"RequestId" xml:"RequestId"`
	ErrCode        string                             `json:"ErrCode" xml:"ErrCode"`
	Success        bool                               `json:"Success" xml:"Success"`
	ErrMessage     string                             `json:"ErrMessage" xml:"ErrMessage"`
	TotalPages     int                                `json:"TotalPages" xml:"TotalPages"`
	TotalElements  int                                `json:"TotalElements" xml:"TotalElements"`
	PageSize       int                                `json:"PageSize" xml:"PageSize"`
	Items          ItemsInDescribeIncrementBackupList `json:"Items" xml:"Items"`
}

// CreateDescribeIncrementBackupListRequest creates a request to invoke DescribeIncrementBackupList API
func CreateDescribeIncrementBackupListRequest() (request *DescribeIncrementBackupListRequest) {
	request = &DescribeIncrementBackupListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "DescribeIncrementBackupList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeIncrementBackupListResponse creates a response to parse from DescribeIncrementBackupList response
func CreateDescribeIncrementBackupListResponse() (response *DescribeIncrementBackupListResponse) {
	response = &DescribeIncrementBackupListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
