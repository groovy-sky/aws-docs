# UpdateReplicationGroupMemberAction

Represents a replica to be modified.

## Contents

###### Note

In the following list, the required parameters are described first.

**RegionName**

The Region where the replica exists.

Type: String

Required: Yes

**GlobalSecondaryIndexes**

Replica-specific global secondary index settings.

Type: Array of [ReplicaGlobalSecondaryIndex](api-replicaglobalsecondaryindex.md) objects

Array Members: Minimum number of 1 item.

Required: No

**KMSMasterKeyId**

The AWS KMS key of the replica that should be used for AWS KMS
encryption. To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or
alias ARN. Note that you should only provide this parameter if the key is different from
the default DynamoDB KMS key `alias/aws/dynamodb`.

Type: String

Required: No

**OnDemandThroughputOverride**

Overrides the maximum on-demand throughput for the replica table.

Type: [OnDemandThroughputOverride](api-ondemandthroughputoverride.md) object

Required: No

**ProvisionedThroughputOverride**

Replica-specific provisioned throughput. If not specified, uses the source table's
provisioned throughput settings.

Type: [ProvisionedThroughputOverride](api-provisionedthroughputoverride.md) object

Required: No

**TableClassOverride**

Replica-specific table class. If not specified, uses the source table's table
class.

Type: String

Valid Values: `STANDARD | STANDARD_INFREQUENT_ACCESS`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/updatereplicationgroupmemberaction.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/updatereplicationgroupmemberaction.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/updatereplicationgroupmemberaction.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateKinesisStreamingConfiguration

WarmThroughput

All content copied from https://docs.aws.amazon.com/.
