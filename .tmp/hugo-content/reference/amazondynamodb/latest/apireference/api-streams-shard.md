# Shard

A uniquely identified group of stream records within a stream.

## Contents

###### Note

In the following list, the required parameters are described first.

**ParentShardId**

The shard ID of the current shard's parent.

Type: String

Length Constraints: Minimum length of 28. Maximum length of 65.

Required: No

**SequenceNumberRange**

The range of possible sequence numbers for the shard.

Type: [SequenceNumberRange](api-streams-sequencenumberrange.md) object

Required: No

**ShardId**

The system-generated identifier for this shard.

Type: String

Length Constraints: Minimum length of 28. Maximum length of 65.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/streams-dynamodb-2012-08-10/shard.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/streams-dynamodb-2012-08-10/shard.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/streams-dynamodb-2012-08-10/shard.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SequenceNumberRange

ShardFilter

All content copied from https://docs.aws.amazon.com/.
