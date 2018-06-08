package cdn

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

// DescribeCdnDomainLogs invokes the cdn.DescribeCdnDomainLogs API synchronously
// api document: https://help.aliyun.com/api/cdn/describecdndomainlogs.html
func (client *Client) DescribeCdnDomainLogs(request *DescribeCdnDomainLogsRequest) (response *DescribeCdnDomainLogsResponse, err error) {
	response = CreateDescribeCdnDomainLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCdnDomainLogsWithChan invokes the cdn.DescribeCdnDomainLogs API asynchronously
// api document: https://help.aliyun.com/api/cdn/describecdndomainlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCdnDomainLogsWithChan(request *DescribeCdnDomainLogsRequest) (<-chan *DescribeCdnDomainLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnDomainLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnDomainLogs(request)
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

// DescribeCdnDomainLogsWithCallback invokes the cdn.DescribeCdnDomainLogs API asynchronously
// api document: https://help.aliyun.com/api/cdn/describecdndomainlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCdnDomainLogsWithCallback(request *DescribeCdnDomainLogsRequest, callback func(response *DescribeCdnDomainLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnDomainLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnDomainLogs(request)
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

// DescribeCdnDomainLogsRequest is the request struct for api DescribeCdnDomainLogs
type DescribeCdnDomainLogsRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	LogDay        string           `position:"Query" name:"LogDay"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	StartTime     string           `position:"Query" name:"StartTime"`
	EndTime       string           `position:"Query" name:"EndTime"`
}

// DescribeCdnDomainLogsResponse is the response struct for api DescribeCdnDomainLogs
type DescribeCdnDomainLogsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	PageNumber     int            `json:"PageNumber" xml:"PageNumber"`
	PageSize       int            `json:"PageSize" xml:"PageSize"`
	TotalCount     int            `json:"TotalCount" xml:"TotalCount"`
	DomainLogModel DomainLogModel `json:"DomainLogModel" xml:"DomainLogModel"`
}

// CreateDescribeCdnDomainLogsRequest creates a request to invoke DescribeCdnDomainLogs API
func CreateDescribeCdnDomainLogsRequest() (request *DescribeCdnDomainLogsRequest) {
	request = &DescribeCdnDomainLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnDomainLogs", "", "")
	return
}

// CreateDescribeCdnDomainLogsResponse creates a response to parse from DescribeCdnDomainLogs response
func CreateDescribeCdnDomainLogsResponse() (response *DescribeCdnDomainLogsResponse) {
	response = &DescribeCdnDomainLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
