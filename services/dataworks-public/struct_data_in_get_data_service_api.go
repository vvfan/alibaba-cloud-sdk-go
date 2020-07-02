package dataworks_public

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

// DataInGetDataServiceApi is a nested struct in dataworks_public response
type DataInGetDataServiceApi struct {
	ApiId               int64               `json:"ApiId" xml:"ApiId"`
	ApiMode             int                 `json:"ApiMode" xml:"ApiMode"`
	ApiName             string              `json:"ApiName" xml:"ApiName"`
	ApiPath             string              `json:"ApiPath" xml:"ApiPath"`
	CreatedTime         string              `json:"CreatedTime" xml:"CreatedTime"`
	CreatorId           string              `json:"CreatorId" xml:"CreatorId"`
	Description         string              `json:"Description" xml:"Description"`
	GroupId             string              `json:"GroupId" xml:"GroupId"`
	ModifiedTime        string              `json:"ModifiedTime" xml:"ModifiedTime"`
	OperatorId          string              `json:"OperatorId" xml:"OperatorId"`
	ProjectId           int64               `json:"ProjectId" xml:"ProjectId"`
	RequestMethod       int                 `json:"RequestMethod" xml:"RequestMethod"`
	ResponseContentType int                 `json:"ResponseContentType" xml:"ResponseContentType"`
	Status              int                 `json:"Status" xml:"Status"`
	TenantId            int64               `json:"TenantId" xml:"TenantId"`
	Timeout             int                 `json:"Timeout" xml:"Timeout"`
	VisibleRange        int                 `json:"VisibleRange" xml:"VisibleRange"`
	Protocols           []int               `json:"Protocols" xml:"Protocols"`
	RegistrationDetails RegistrationDetails `json:"RegistrationDetails" xml:"RegistrationDetails"`
	ScriptDetails       ScriptDetails       `json:"ScriptDetails" xml:"ScriptDetails"`
	WizardDetails       WizardDetails       `json:"WizardDetails" xml:"WizardDetails"`
}
