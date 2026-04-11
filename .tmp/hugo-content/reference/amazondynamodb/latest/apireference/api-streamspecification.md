# StreamSpecification

Represents the DynamoDB Streams configuration for a table in DynamoDB.

## Contents

###### Note

In the following list, the required parameters are described first.

**StreamEnabled**

Indicates whether DynamoDB Streams is enabled (true) or disabled (false) on the
table.

Type: Boolean

Required: Yes

**StreamViewType**

When an item in the table is modified, `StreamViewType` determines what
information is written to the stream for this table. Valid values for
`StreamViewType` are:

- `KEYS_ONLY` \- Only the key attributes of the modified item are
written to the stream.

- `NEW_IMAGE` \- The entire item, as it appears after it was modified,
is written to the stream.

- `OLD_IMAGE` \- The entire item, as it appeared before it was modified,
is written to the stream.

- `NEW_AND_OLD_IMAGES` \- Both the new and the old item images of the
item are written to the stream.

Type: String

Valid Values: `NEW_IMAGE | OLD_IMAGE | NEW_AND_OLD_IMAGES | KEYS_ONLY`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/streamspecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/streamspecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/streamspecification.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SSESpecification

TableAutoScalingDescription

All content copied from https://docs.aws.amazon.com/.
