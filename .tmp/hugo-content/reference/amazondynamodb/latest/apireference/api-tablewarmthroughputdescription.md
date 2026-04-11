# TableWarmThroughputDescription

Represents the warm throughput value (in read units per second and write units per
second) of the table. Warm throughput is applicable for DynamoDB Standard-IA tables and
specifies the minimum provisioned capacity maintained for immediate data access.

## Contents

###### Note

In the following list, the required parameters are described first.

**ReadUnitsPerSecond**

Represents the base table's warm throughput value in read units per second.

Type: Long

Valid Range: Minimum value of 1.

Required: No

**Status**

Represents warm throughput value of the base table.

Type: String

Valid Values: `CREATING | UPDATING | DELETING | ACTIVE | INACCESSIBLE_ENCRYPTION_CREDENTIALS | ARCHIVING | ARCHIVED | REPLICATION_NOT_AUTHORIZED`

Required: No

**WriteUnitsPerSecond**

Represents the base table's warm throughput value in write units per second.

Type: Long

Valid Range: Minimum value of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/tablewarmthroughputdescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/tablewarmthroughputdescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/tablewarmthroughputdescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableDescription

Tag

All content copied from https://docs.aws.amazon.com/.
