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

type DescribeMetricDataRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 监控项英文标识(id)  */
    Metric string `json:"metric"`

    /* 资源的类型，取值vm, lb, ip, database 等  */
    ServiceCode string `json:"serviceCode"`

    /* 资源的uuid  */
    ResourceId string `json:"resourceId"`

    /* 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ（默认为当前时间，早于30d时，将被重置为30d） (Optional) */
    StartTime *string `json:"startTime"`

    /* 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，将由startTime与timeInterval计算得出） (Optional) */
    EndTime *string `json:"endTime"`

    /* 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval 与 endTime 至少填一项 (Optional) */
    TimeInterval *string `json:"timeInterval"`

    /* 自定义标签 (Optional) */
    Tags []monitor.TagFilter `json:"tags"`
}

/*
 * param regionId: 地域 Id (Required)
 * param metric: 监控项英文标识(id) (Required)
 * param serviceCode: 资源的类型，取值vm, lb, ip, database 等 (Required)
 * param resourceId: 资源的uuid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeMetricDataRequest(
    regionId string,
    metric string,
    serviceCode string,
    resourceId string,
) *DescribeMetricDataRequest {

	return &DescribeMetricDataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/metrics/{metric}/metricData",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Metric: metric,
        ServiceCode: serviceCode,
        ResourceId: resourceId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param metric: 监控项英文标识(id) (Required)
 * param serviceCode: 资源的类型，取值vm, lb, ip, database 等 (Required)
 * param resourceId: 资源的uuid (Required)
 * param startTime: 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ（默认为当前时间，早于30d时，将被重置为30d） (Optional)
 * param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，将由startTime与timeInterval计算得出） (Optional)
 * param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval 与 endTime 至少填一项 (Optional)
 * param tags: 自定义标签 (Optional)
 */
func NewDescribeMetricDataRequestWithAllParams(
    regionId string,
    metric string,
    serviceCode string,
    resourceId string,
    startTime *string,
    endTime *string,
    timeInterval *string,
    tags []monitor.TagFilter,
) *DescribeMetricDataRequest {

    return &DescribeMetricDataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/metrics/{metric}/metricData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Metric: metric,
        ServiceCode: serviceCode,
        ResourceId: resourceId,
        StartTime: startTime,
        EndTime: endTime,
        TimeInterval: timeInterval,
        Tags: tags,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeMetricDataRequestWithoutParam() *DescribeMetricDataRequest {

    return &DescribeMetricDataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/metrics/{metric}/metricData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeMetricDataRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param metric: 监控项英文标识(id)(Required) */
func (r *DescribeMetricDataRequest) SetMetric(metric string) {
    r.Metric = metric
}

/* param serviceCode: 资源的类型，取值vm, lb, ip, database 等(Required) */
func (r *DescribeMetricDataRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

/* param resourceId: 资源的uuid(Required) */
func (r *DescribeMetricDataRequest) SetResourceId(resourceId string) {
    r.ResourceId = resourceId
}

/* param startTime: 查询时间范围的开始时间， UTC时间，格式：yyyy-MM-dd'T'HH:mm:ssZ（默认为当前时间，早于30d时，将被重置为30d）(Optional) */
func (r *DescribeMetricDataRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 查询时间范围的结束时间， UTC时间，格式：2016-12- yyyy-MM-dd'T'HH:mm:ssZ（为空时，将由startTime与timeInterval计算得出）(Optional) */
func (r *DescribeMetricDataRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param timeInterval: 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval 与 endTime 至少填一项(Optional) */
func (r *DescribeMetricDataRequest) SetTimeInterval(timeInterval string) {
    r.TimeInterval = &timeInterval
}

/* param tags: 自定义标签(Optional) */
func (r *DescribeMetricDataRequest) SetTags(tags []monitor.TagFilter) {
    r.Tags = tags
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeMetricDataRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeMetricDataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeMetricDataResult `json:"result"`
}

type DescribeMetricDataResult struct {
    MetricDatas []monitor.MetricData `json:"metricDatas"`
}