# WarmThroughput

Provides visibility into the number of read and write operations your table or
secondary index can instantaneously support. The settings can be modified using the
`UpdateTable` operation to meet the throughput requirements of an
upcoming peak event.

## Contents

###### Note

In the following list, the required parameters are described first.

**ReadUnitsPerSecond**

Represents the number of read operations your base table can instantaneously
support.

Type: Long

Required: No

**WriteUnitsPerSecond**

Represents the number of write operations your base table can instantaneously
support.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/warmthroughput.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/warmthroughput.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/warmthroughput.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateReplicationGroupMemberAction

WriteRequest

All content copied from https://docs.aws.amazon.com/.
