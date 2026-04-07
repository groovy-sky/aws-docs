# GetCapacityManagerMetricData

Retrieves capacity usage metrics for your EC2 resources. Returns time-series data for metrics like unused capacity, utilization rates, and costs
across On-Demand, Spot, and Capacity Reservations. Data can be grouped and filtered by various dimensions such as region, account, and instance family.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have
the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**EndTime**

The end time for the metric data query, in ISO 8601 format. If the end time is beyond the latest ingested data, it will be automatically adjusted to the latest available data point.

Type: Timestamp

Required: Yes

**FilterBy.N**

Conditions to filter the metric data. Each filter specifies a dimension, comparison operator ('equals', 'in'), and values to match against.

Type: Array of [CapacityManagerCondition](api-capacitymanagercondition.md) objects

Array Members: Minimum number of 0 items. Maximum number of 20 items.

Required: No

**GroupBy.N**

The dimensions by which to group the metric data. This determines how the data is aggregated and returned.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 20 items.

Length Constraints: Minimum length of 0. Maximum length of 300.

Valid Values: `resource-region | availability-zone-id | account-id | instance-family | instance-type | instance-platform | reservation-arn | reservation-id | reservation-type | reservation-create-timestamp | reservation-start-timestamp | reservation-end-timestamp | reservation-end-date-type | tenancy | reservation-state | reservation-instance-match-criteria | reservation-unused-financial-owner`

Required: No

**MaxResults**

The maximum number of data points to return. Valid range is 1 to 100,000. Use with NextToken for pagination of large result sets.

Type: Integer

Required: No

**MetricName.N**

The names of the metrics to retrieve. Maximum of 10 metrics per request.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 40 items.

Valid Values: `reservation-total-capacity-hrs-vcpu | reservation-total-capacity-hrs-inst | reservation-max-size-vcpu | reservation-max-size-inst | reservation-min-size-vcpu | reservation-min-size-inst | reservation-unused-total-capacity-hrs-vcpu | reservation-unused-total-capacity-hrs-inst | reservation-unused-total-estimated-cost | reservation-max-unused-size-vcpu | reservation-max-unused-size-inst | reservation-min-unused-size-vcpu | reservation-min-unused-size-inst | reservation-max-utilization | reservation-min-utilization | reservation-avg-utilization-vcpu | reservation-avg-utilization-inst | reservation-total-count | reservation-total-estimated-cost | reservation-avg-future-size-vcpu | reservation-avg-future-size-inst | reservation-min-future-size-vcpu | reservation-min-future-size-inst | reservation-max-future-size-vcpu | reservation-max-future-size-inst | reservation-avg-committed-size-vcpu | reservation-avg-committed-size-inst | reservation-max-committed-size-vcpu | reservation-max-committed-size-inst | reservation-min-committed-size-vcpu | reservation-min-committed-size-inst | reserved-total-usage-hrs-vcpu | reserved-total-usage-hrs-inst | reserved-total-estimated-cost | unreserved-total-usage-hrs-vcpu | unreserved-total-usage-hrs-inst | unreserved-total-estimated-cost | spot-total-usage-hrs-vcpu | spot-total-usage-hrs-inst | spot-total-estimated-cost | spot-avg-run-time-before-interruption-inst | spot-max-run-time-before-interruption-inst | spot-min-run-time-before-interruption-inst`

Required: Yes

**NextToken**

The token for the next page of results. Use this value in a subsequent call to retrieve additional data points.

Type: String

Required: No

**Period**

The granularity, in seconds, of the returned data points.

Type: Integer

Valid Range: Minimum value of 3600.

Required: Yes

**StartTime**

The start time for the metric data query, in ISO 8601 format. The time range (end time - start time) must be a multiple of the specified period.

Type: Timestamp

Required: Yes

## Response Elements

The following elements are returned by the service.

**metricDataResultSet**

The metric data points returned by the query. Each result contains dimension values, timestamp, and metric values with their associated statistics.

Type: Array of [MetricDataResult](api-metricdataresult.md) objects

**nextToken**

The token to use to retrieve the next page of results. This value is null when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetCapacityManagerMetricData)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetCapacityManagerMetricData)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getcapacitymanagermetricdata.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getcapacitymanagermetricdata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getcapacitymanagermetricdata.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getcapacitymanagermetricdata.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getcapacitymanagermetricdata.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getcapacitymanagermetricdata.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetCapacityManagerMetricData)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getcapacitymanagermetricdata.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetCapacityManagerAttributes

GetCapacityManagerMetricDimensions
