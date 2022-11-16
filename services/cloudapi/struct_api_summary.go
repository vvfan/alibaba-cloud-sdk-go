package cloudapi

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

// ApiSummary is a nested struct in cloudapi response
type ApiSummary struct {
	Visibility   string `json:"Visibility" xml:"Visibility"`
	ApiId        string `json:"ApiId" xml:"ApiId"`
	ModifiedTime string `json:"ModifiedTime" xml:"ModifiedTime"`
	Description  string `json:"Description" xml:"Description"`
	GroupName    string `json:"GroupName" xml:"GroupName"`
	GroupId      string `json:"GroupId" xml:"GroupId"`
	CreatedTime  string `json:"CreatedTime" xml:"CreatedTime"`
	ApiName      string `json:"ApiName" xml:"ApiName"`
	RegionId     string `json:"RegionId" xml:"RegionId"`
	ApiPath      string `json:"ApiPath" xml:"ApiPath"`
	ApiMethod    string `json:"ApiMethod" xml:"ApiMethod"`
}
