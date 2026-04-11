# DestinationOptionsRequest

Describes the destination options for a flow log.

## Contents

**FileFormat**

The format for the flow log. The default is `plain-text`.

Type: String

Valid Values: `plain-text | parquet`

Required: No

**HiveCompatiblePartitions**

Indicates whether to use Hive-compatible prefixes for flow logs stored in Amazon S3.
The default is `false`.

Type: Boolean

Required: No

**PerHourPartition**

Indicates whether to partition the flow log per hour. This reduces the cost and response
time for queries. The default is `false`.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/destinationoptionsrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/destinationoptionsrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/destinationoptionsrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeFleetsInstances

DestinationOptionsResponse
