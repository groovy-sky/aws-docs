# ReplicaGlobalSecondaryIndex

Represents the properties of a replica global secondary index.

## Contents

###### Note

In the following list, the required parameters are described first.

**IndexName**

The name of the global secondary index.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: Yes

**OnDemandThroughputOverride**

Overrides the maximum on-demand throughput settings for the specified global secondary
index in the specified replica table.

Type: [OnDemandThroughputOverride](api-ondemandthroughputoverride.md) object

Required: No

**ProvisionedThroughputOverride**

Replica table GSI-specific provisioned throughput. If not specified, uses the source
table GSI's read capacity settings.

Type: [ProvisionedThroughputOverride](api-provisionedthroughputoverride.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/replicaglobalsecondaryindex.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/replicaglobalsecondaryindex.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/replicaglobalsecondaryindex.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicaDescription

ReplicaGlobalSecondaryIndexAutoScalingDescription

All content copied from https://docs.aws.amazon.com/.
