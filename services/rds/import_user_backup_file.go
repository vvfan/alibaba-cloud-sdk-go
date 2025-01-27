package rds

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

// ImportUserBackupFile invokes the rds.ImportUserBackupFile API synchronously
func (client *Client) ImportUserBackupFile(request *ImportUserBackupFileRequest) (response *ImportUserBackupFileResponse, err error) {
	response = CreateImportUserBackupFileResponse()
	err = client.DoAction(request, response)
	return
}

// ImportUserBackupFileWithChan invokes the rds.ImportUserBackupFile API asynchronously
func (client *Client) ImportUserBackupFileWithChan(request *ImportUserBackupFileRequest) (<-chan *ImportUserBackupFileResponse, <-chan error) {
	responseChan := make(chan *ImportUserBackupFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportUserBackupFile(request)
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

// ImportUserBackupFileWithCallback invokes the rds.ImportUserBackupFile API asynchronously
func (client *Client) ImportUserBackupFileWithCallback(request *ImportUserBackupFileRequest, callback func(response *ImportUserBackupFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportUserBackupFileResponse
		var err error
		defer close(result)
		response, err = client.ImportUserBackupFile(request)
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

// ImportUserBackupFileRequest is the request struct for api ImportUserBackupFile
type ImportUserBackupFileRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	Retention            requests.Integer `position:"Query" name:"Retention"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	BackupFile           string           `position:"Query" name:"BackupFile"`
	BucketRegion         string           `position:"Query" name:"BucketRegion"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RestoreSize          requests.Integer `position:"Query" name:"RestoreSize"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	Comment              string           `position:"Query" name:"Comment"`
}

// ImportUserBackupFileResponse is the response struct for api ImportUserBackupFile
type ImportUserBackupFileResponse struct {
	*responses.BaseResponse
	Status    bool   `json:"Status" xml:"Status"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	BackupId  string `json:"BackupId" xml:"BackupId"`
}

// CreateImportUserBackupFileRequest creates a request to invoke ImportUserBackupFile API
func CreateImportUserBackupFileRequest() (request *ImportUserBackupFileRequest) {
	request = &ImportUserBackupFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ImportUserBackupFile", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateImportUserBackupFileResponse creates a response to parse from ImportUserBackupFile response
func CreateImportUserBackupFileResponse() (response *ImportUserBackupFileResponse) {
	response = &ImportUserBackupFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
