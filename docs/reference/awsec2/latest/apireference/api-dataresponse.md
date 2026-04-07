# DataResponse

The response to a `DataQuery`.

## Contents

**destination**

The Region or Availability Zone that's the destination for the data query. For example, `eu-west-1`.

Type: String

Required: No

**id**

The ID passed in the `DataQuery`.

Type: String

Required: No

**metric**

The metric used for the network performance request.

Type: String

Valid Values: `aggregate-latency`

Required: No

**MetricPointSet.N**

A list of `MetricPoint` objects.

Type: Array of [MetricPoint](api-metricpoint.md) objects

Required: No

**period**

The period used for the network performance request.

Type: String

Valid Values: `five-minutes | fifteen-minutes | one-hour | three-hours | one-day | one-week`

Required: No

**source**

The Region or Availability Zone that's the source for the data query. For example, `us-east-1`.

Type: String

Required: No

**statistic**

The statistic used for the network performance request.

Type: String

Valid Values: `p50`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/dataresponse.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/dataresponse.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/dataresponse.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DataQuery

DeclarativePoliciesReport
