# Stream

Represents all of the data describing a particular stream.

## Contents

###### Note

In the following list, the required parameters are described first.

**StreamArn**

The Amazon Resource Name (ARN) for the stream.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

Required: No

**StreamLabel**

A timestamp, in ISO 8601 format, for this stream.

Note that `LatestStreamLabel` is not a unique identifier for the stream, because it is
possible that a stream from another table might have the same timestamp. However, the
combination of the following three elements is guaranteed to be unique:

- the AWS customer ID.

- the table name

- the `StreamLabel`

Type: String

Required: No

**TableName**

The DynamoDB table with which the stream is associated.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/streams-dynamodb-2012-08-10/stream.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/streams-dynamodb-2012-08-10/stream.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/streams-dynamodb-2012-08-10/stream.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ShardFilter

StreamDescription

All content copied from https://docs.aws.amazon.com/.
