# CreateGlobalSecondaryIndexAction

Represents a new global secondary index to be added to an existing table.

## Contents

###### Note

In the following list, the required parameters are described first.

**IndexName**

The name of the global secondary index to be created.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: Yes

**KeySchema**

The key schema for the global secondary index.
Global secondary index supports up to 4 partition and up to 4 sort keys.

Type: Array of [KeySchemaElement](api-keyschemaelement.md) objects

Array Members: Minimum number of 1 item.

Required: Yes

**Projection**

Represents attributes that are copied (projected) from the table into an index. These
are in addition to the primary key attributes and index key attributes, which are
automatically projected.

Type: [Projection](api-projection.md) object

Required: Yes

**OnDemandThroughput**

The maximum number of read and write units for the global secondary index being
created. If you use this parameter, you must specify `MaxReadRequestUnits`,
`MaxWriteRequestUnits`, or both. You must use either `OnDemand
                Throughput` or `ProvisionedThroughput` based on your table's
capacity mode.

Type: [OnDemandThroughput](api-ondemandthroughput.md) object

Required: No

**ProvisionedThroughput**

Represents the provisioned throughput settings for the specified global secondary
index.

For current minimum and maximum provisioned throughput values, see [Service,\
Account, and Table Quotas](../../../../services/dynamodb/latest/developerguide/limits.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: [ProvisionedThroughput](api-provisionedthroughput.md) object

Required: No

**WarmThroughput**

Represents the warm throughput value (in read units per second and write units per
second) when creating a secondary index.

Type: [WarmThroughput](api-warmthroughput.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/createglobalsecondaryindexaction.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/createglobalsecondaryindexaction.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/createglobalsecondaryindexaction.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContributorInsightsSummary

CreateGlobalTableWitnessGroupMemberAction

All content copied from https://docs.aws.amazon.com/.
