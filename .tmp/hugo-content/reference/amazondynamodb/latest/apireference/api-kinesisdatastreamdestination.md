# KinesisDataStreamDestination

Describes a Kinesis data stream destination.

## Contents

###### Note

In the following list, the required parameters are described first.

**ApproximateCreationDateTimePrecision**

The precision of the Kinesis data stream timestamp. The values are either
`MILLISECOND` or `MICROSECOND`.

Type: String

Valid Values: `MILLISECOND | MICROSECOND`

Required: No

**DestinationStatus**

The current status of replication.

Type: String

Valid Values: `ENABLING | ACTIVE | DISABLING | DISABLED | ENABLE_FAILED | UPDATING`

Required: No

**DestinationStatusDescription**

The human-readable string that corresponds to the replica status.

Type: String

Required: No

**StreamArn**

The ARN for a specific Kinesis data stream.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/kinesisdatastreamdestination.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/kinesisdatastreamdestination.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/kinesisdatastreamdestination.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeySchemaElement

LocalSecondaryIndex

All content copied from https://docs.aws.amazon.com/.
