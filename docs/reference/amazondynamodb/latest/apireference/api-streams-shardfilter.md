---
title: "ShardFilter"
---

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/streams-dynamodb-2012-08-10/ShardFilter)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/streams-dynamodb-2012-08-10/ShardFilter)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/streams-dynamodb-2012-08-10/ShardFilter)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Shard

Stream

All content copied from https://docs.aws.amazon.com/.
