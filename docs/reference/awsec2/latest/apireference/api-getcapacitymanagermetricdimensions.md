# GetCapacityManagerMetricDimensions

Retrieves the available dimension values for capacity metrics within a specified time range. This is useful for discovering what accounts,
regions, instance families, and other dimensions have data available for filtering and grouping.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides
an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**EndTime**

The end time for the dimension query, in ISO 8601 format. Only dimensions with data in this time range will be returned.

Type: Timestamp

Required: Yes

**FilterBy.N**

Conditions to filter which dimension values are returned. Each filter specifies a dimension, comparison operator, and values to match against.

Type: Array of [CapacityManagerCondition](api-capacitymanagercondition.md) objects

Array Members: Minimum number of 0 items. Maximum number of 20 items.

Required: No

**GroupBy.N**

The dimensions to group by when retrieving available dimension values. This determines which dimension combinations are returned. Required parameter.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 20 items.

Length Constraints: Minimum length of 0. Maximum length of 300.

Valid Values: `resource-region | availability-zone-id | account-id | instance-family | instance-type | instance-platform | reservation-arn | reservation-id | reservation-type | reservation-create-timestamp | reservation-start-timestamp | reservation-end-timestamp | reservation-end-date-type | tenancy | reservation-state | reservation-instance-match-criteria | reservation-unused-financial-owner`

Required: Yes

**MaxResults**

The maximum number of dimension combinations to return. Valid range is 1 to 1000. Use with NextToken for pagination.

Type: Integer

Required: No

**MetricName.N**

The metric names to use as an additional filter when retrieving dimensions. Only dimensions that have data for these
metrics will be returned. Required parameter with maximum size of 1 for v1.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 40 items.

Valid Values: `reservation-total-capacity-hrs-vcpu | reservation-total-capacity-hrs-inst | reservation-max-size-vcpu | reservation-max-size-inst | reservation-min-size-vcpu | reservation-min-size-inst | reservation-unused-total-capacity-hrs-vcpu | reservation-unused-total-capacity-hrs-inst | reservation-unused-total-estimated-cost | reservation-max-unused-size-vcpu | reservation-max-unused-size-inst | reservation-min-unused-size-vcpu | reservation-min-unused-size-inst | reservation-max-utilization | reservation-min-utilization | reservation-avg-utilization-vcpu | reservation-avg-utilization-inst | reservation-total-count | reservation-total-estimated-cost | reservation-avg-future-size-vcpu | reservation-avg-future-size-inst | reservation-min-future-size-vcpu | reservation-min-future-size-inst | reservation-max-future-size-vcpu | reservation-max-future-size-inst | reservation-avg-committed-size-vcpu | reservation-avg-committed-size-inst | reservation-max-committed-size-vcpu | reservation-max-committed-size-inst | reservation-min-committed-size-vcpu | reservation-min-committed-size-inst | reserved-total-usage-hrs-vcpu | reserved-total-usage-hrs-inst | reserved-total-estimated-cost | unreserved-total-usage-hrs-vcpu | unreserved-total-usage-hrs-inst | unreserved-total-estimated-cost | spot-total-usage-hrs-vcpu | spot-total-usage-hrs-inst | spot-total-estimated-cost | spot-avg-run-time-before-interruption-inst | spot-max-run-time-before-interruption-inst | spot-min-run-time-before-interruption-inst`

Required: Yes

**NextToken**

The token for the next page of results. Use this value in a subsequent call to retrieve additional dimension values.

Type: String

Required: No

**StartTime**

The start time for the dimension query, in ISO 8601 format. Only dimensions with data in this time range will be returned.

Type: Timestamp

Required: Yes

## Response Elements

The following elements are returned by the service.

**metricDimensionResultSet**

The available dimension combinations that have data within the specified time range and filters.

Type: Array of [CapacityManagerDimension](api-capacitymanagerdimension.md) objects

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetCapacityManagerMetricDimensions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetCapacityManagerMetricDimensions)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getcapacitymanagermetricdimensions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getcapacitymanagermetricdimensions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getcapacitymanagermetricdimensions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getcapacitymanagermetricdimensions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getcapacitymanagermetricdimensions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getcapacitymanagermetricdimensions.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetCapacityManagerMetricDimensions)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getcapacitymanagermetricdimensions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetCapacityManagerMetricData

GetCapacityReservationUsage
