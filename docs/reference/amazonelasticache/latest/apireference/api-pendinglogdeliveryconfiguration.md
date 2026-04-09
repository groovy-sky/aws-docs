# PendingLogDeliveryConfiguration

The log delivery configurations being modified

## Contents

###### Note

In the following list, the required parameters are described first.

**DestinationDetails**

Configuration details of either a CloudWatch Logs destination or Kinesis Data Firehose
destination.

Type: [DestinationDetails](api-destinationdetails.md) object

Required: No

**DestinationType**

Returns the destination type, either CloudWatch Logs or Kinesis Data Firehose.

Type: String

Valid Values: `cloudwatch-logs | kinesis-firehose`

Required: No

**LogFormat**

Returns the log format, either JSON or TEXT

Type: String

Valid Values: `text | json`

Required: No

**LogType**

Refers to [slow-log](https://redis.io/commands/slowlog) or
engine-log..

Type: String

Valid Values: `slow-log | engine-log`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/pendinglogdeliveryconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/pendinglogdeliveryconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/pendinglogdeliveryconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterNameValue

PendingModifiedValues

All content copied from https://docs.aws.amazon.com/.
