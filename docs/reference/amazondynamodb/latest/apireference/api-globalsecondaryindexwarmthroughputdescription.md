# GlobalSecondaryIndexWarmThroughputDescription

The description of the warm throughput value on a global secondary index.

## Contents

###### Note

In the following list, the required parameters are described first.

**ReadUnitsPerSecond**

Represents warm throughput read units per second value for a global secondary
index.

Type: Long

Valid Range: Minimum value of 1.

Required: No

**Status**

Represents the warm throughput status being created or updated on a global secondary
index. The status can only be `UPDATING` or `ACTIVE`.

Type: String

Valid Values: `CREATING | UPDATING | DELETING | ACTIVE`

Required: No

**WriteUnitsPerSecond**

Represents warm throughput write units per second value for a global secondary
index.

Type: Long

Valid Range: Minimum value of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/globalsecondaryindexwarmthroughputdescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/globalsecondaryindexwarmthroughputdescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/globalsecondaryindexwarmthroughputdescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlobalSecondaryIndexUpdate

GlobalTable

All content copied from https://docs.aws.amazon.com/.
