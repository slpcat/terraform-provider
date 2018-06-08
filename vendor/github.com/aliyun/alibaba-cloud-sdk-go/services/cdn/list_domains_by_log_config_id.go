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

// ListDomainsByLogConfigId invokes the cdn.ListDomainsByLogConfigId API synchronously
// api document: https://help.aliyun.com/api/cdn/listdomainsbylogconfigid.html
func (client *Client) ListDomainsByLogConfigId(request *ListDomainsByLogConfigIdRequest) (response *ListDomainsByLogConfigIdResponse, err error) {
	response = CreateListDomainsByLogConfigIdResponse()
	err = client.DoAction(request, response)
	return
}

// ListDomainsByLogConfigIdWithChan invokes the cdn.ListDomainsByLogConfigId API asynchronously
// api document: https://help.aliyun.com/api/cdn/listdomainsbylogconfigid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDomainsByLogConfigIdWithChan(request *ListDomainsByLogConfigIdRequest) (<-chan *ListDomainsByLogConfigIdResponse, <-chan error) {
	responseChan := make(chan *ListDomainsByLogConfigIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDomainsByLogConfigId(request)
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

// ListDomainsByLogConfigIdWithCallback invokes the cdn.ListDomainsByLogConfigId API asynchronously
// api document: https://help.aliyun.com/api/cdn/listdomainsbylogconfigid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDomainsByLogConfigIdWithCallback(request *ListDomainsByLogConfigIdRequest, callback func(response *ListDomainsByLogConfigIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDomainsByLogConfigIdResponse
		var err error
		defer close(result)
		response, err = client.ListDomainsByLogConfigId(request)
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

// ListDomainsByLogConfigIdRequest is the request struct for api ListDomainsByLogConfigId
type ListDomainsByLogConfigIdRequest struct {
	*requests.RpcRequest
}

// ListDomainsByLogConfigIdResponse is the response struct for api ListDomainsByLogConfigId
type ListDomainsByLogConfigIdResponse struct {
	*responses.BaseResponse
	RequestId string                            `json:"RequestId" xml:"RequestId"`
	Domains   DomainsInListDomainsByLogConfigId `json:"Domains" xml:"Domains"`
}

// CreateListDomainsByLogConfigIdRequest creates a request to invoke ListDomainsByLogConfigId API
func CreateListDomainsByLogConfigIdRequest() (request *ListDomainsByLogConfigIdRequest) {
	request = &ListDomainsByLogConfigIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "ListDomainsByLogConfigId", "", "")
	return
}

// CreateListDomainsByLogConfigIdResponse creates a response to parse from ListDomainsByLogConfigId response
func CreateListDomainsByLogConfigIdResponse() (response *ListDomainsByLogConfigIdResponse) {
	response = &ListDomainsByLogConfigIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
