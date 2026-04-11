# ProvisionedThroughputDescription

Represents the provisioned throughput settings for the table, consisting of read and
write capacity units, along with data about increases and decreases.

## Contents

###### Note

In the following list, the required parameters are described first.

**LastDecreaseDateTime**

The date and time of the last provisioned throughput decrease for this table.

Type: Timestamp

Required: No

**LastIncreaseDateTime**

The date and time of the last provisioned throughput increase for this table.

Type: Timestamp

Required: No

**NumberOfDecreasesToday**

The number of provisioned throughput decreases for this table during this UTC calendar
day. For current maximums on provisioned throughput decreases, see [Service,\
Account, and Table Quotas](../../../../services/dynamodb/latest/developerguide/limits.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: Long

Valid Range: Minimum value of 1.

Required: No

**ReadCapacityUnits**

The maximum number of strongly consistent reads consumed per second before DynamoDB
returns a `ThrottlingException`. Eventually consistent reads require less
effort than strongly consistent reads, so a setting of 50 `ReadCapacityUnits`
per second provides 100 eventually consistent `ReadCapacityUnits` per
second.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**WriteCapacityUnits**

The maximum number of writes consumed per second before DynamoDB returns a
`ThrottlingException`.

Type: Long

Valid Range: Minimum value of 0.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/provisionedthroughputdescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/provisionedthroughputdescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/provisionedthroughputdescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProvisionedThroughput

ProvisionedThroughputOverride

All content copied from https://docs.aws.amazon.com/.
