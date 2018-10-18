// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type DescribeAlarmHistoryRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 报警规则的Id (Optional) */
    Id *string `json:"id"`

    /* 产品名称 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 资源Id (Optional) */
    ResourceId *string `json:"resourceId"`

    /* 查询数据开始时间，默认24小时前，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间  */
    StartTime string `json:"startTime"`

    /* 查询数据结束时间，默认当前时间，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间  */
    EndTime string `json:"endTime"`

    /* 页码, 默认为1, 取值范围：[1,∞) (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域 Id (Required)
 * param startTime: 查询数据开始时间，默认24小时前，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间 (Required)
 * param endTime: 查询数据结束时间，默认当前时间，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAlarmHistoryRequest(
    regionId string,
    startTime string,
    endTime string,
) *DescribeAlarmHistoryRequest {

	return &DescribeAlarmHistoryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/alarmHistory",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param id: 报警规则的Id (Optional)
 * param serviceCode: 产品名称 (Optional)
 * param resourceId: 资源Id (Optional)
 * param startTime: 查询数据开始时间，默认24小时前，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间 (Required)
 * param endTime: 查询数据结束时间，默认当前时间，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间 (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞) (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 */
func NewDescribeAlarmHistoryRequestWithAllParams(
    regionId string,
    id *string,
    serviceCode *string,
    resourceId *string,
    startTime string,
    endTime string,
    pageNumber *int,
    pageSize *int,
) *DescribeAlarmHistoryRequest {

    return &DescribeAlarmHistoryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/alarmHistory",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
        ServiceCode: serviceCode,
        ResourceId: resourceId,
        StartTime: startTime,
        EndTime: endTime,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAlarmHistoryRequestWithoutParam() *DescribeAlarmHistoryRequest {

    return &DescribeAlarmHistoryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/alarmHistory",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeAlarmHistoryRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: 报警规则的Id(Optional) */
func (r *DescribeAlarmHistoryRequest) SetId(id string) {
    r.Id = &id
}

/* param serviceCode: 产品名称(Optional) */
func (r *DescribeAlarmHistoryRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param resourceId: 资源Id(Optional) */
func (r *DescribeAlarmHistoryRequest) SetResourceId(resourceId string) {
    r.ResourceId = &resourceId
}

/* param startTime: 查询数据开始时间，默认24小时前，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间(Required) */
func (r *DescribeAlarmHistoryRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 查询数据结束时间，默认当前时间，可以输入long型时间，也可以输入yyyy-MM-dd'T'HH:mm:ssZ类型时间(Required) */
func (r *DescribeAlarmHistoryRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞)(Optional) */
func (r *DescribeAlarmHistoryRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *DescribeAlarmHistoryRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAlarmHistoryRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAlarmHistoryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAlarmHistoryResult `json:"result"`
}

type DescribeAlarmHistoryResult struct {
    AlarmHistoryList []monitor.AlarmHistory `json:"alarmHistoryList"`
    PageNumber int `json:"pageNumber"`
    NumberPages int `json:"numberPages"`
    NumberRecords int `json:"numberRecords"`
    PageSize int `json:"pageSize"`
}