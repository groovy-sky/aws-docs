# GlobalSecondaryIndexUpdate

Represents one of the following:

- A new global secondary index to be added to an existing table.

- New provisioned throughput parameters for an existing global secondary
index.

- An existing global secondary index to be removed from an existing
table.

## Contents

###### Note

In the following list, the required parameters are described first.

**Create**

The parameters required for creating a global secondary index on an existing
table:

- `IndexName `

- `KeySchema `

- `AttributeDefinitions `

- `Projection `

- `ProvisionedThroughput `

Type: [CreateGlobalSecondaryIndexAction](api-createglobalsecondaryindexaction.md) object

Required: No

**Delete**

The name of an existing global secondary index to be removed.

Type: [DeleteGlobalSecondaryIndexAction](api-deleteglobalsecondaryindexaction.md) object

Required: No

**Update**

The name of an existing global secondary index, along with new provisioned throughput
settings to be applied to that index.

Type: [UpdateGlobalSecondaryIndexAction](api-updateglobalsecondaryindexaction.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/globalsecondaryindexupdate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/globalsecondaryindexupdate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/globalsecondaryindexupdate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlobalSecondaryIndexInfo

GlobalSecondaryIndexWarmThroughputDescription

All content copied from https://docs.aws.amazon.com/.
