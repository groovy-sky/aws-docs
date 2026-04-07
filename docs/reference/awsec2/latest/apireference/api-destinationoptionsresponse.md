# DestinationOptionsResponse

Describes the destination options for a flow log.

## Contents

**fileFormat**

The format for the flow log.

Type: String

Valid Values: `plain-text | parquet`

Required: No

**hiveCompatiblePartitions**

Indicates whether to use Hive-compatible prefixes for flow logs stored in Amazon S3.

Type: Boolean

Required: No

**perHourPartition**

Indicates whether to partition the flow log per hour.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/destinationoptionsresponse.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/destinationoptionsresponse.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/destinationoptionsresponse.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DestinationOptionsRequest

DeviceOptions
