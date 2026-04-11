# ShardFilter

This optional field contains the filter
definition for the `DescribeStream` API.

## Contents

###### Note

In the following list, the required parameters are described first.

**ShardId**

Contains the `shardId` of the parent shard for which you are requesting child shards.

_Sample request:_

Type: String

Length Constraints: Minimum length of 28. Maximum length of 65.

Required: No

**Type**

Contains the type of filter to be applied on the `DescribeStream` API.
Currently, the only value this parameter accepts is `CHILD_SHARDS`.

Type: String

Valid Values: `CHILD_SHARDS`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/streams-dynamodb-2012-08-10/shardfilter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/streams-dynamodb-2012-08-10/shardfilter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/streams-dynamodb-2012-08-10/shardfilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Shard

Stream

All content copied from https://docs.aws.amazon.com/.
