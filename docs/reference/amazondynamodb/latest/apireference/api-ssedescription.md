# SSEDescription

The description of the server-side encryption status on the specified table.

## Contents

###### Note

In the following list, the required parameters are described first.

**InaccessibleEncryptionDateTime**

Indicates the time, in UNIX epoch date format, when DynamoDB detected that
the table's AWS KMS key was inaccessible. This attribute will automatically
be cleared when DynamoDB detects that the table's AWS KMS key is accessible
again. DynamoDB will initiate the table archival process when table's AWS KMS key remains inaccessible for more than seven days from this date.

Type: Timestamp

Required: No

**KMSMasterKeyArn**

The AWS KMS key ARN used for the AWS KMS encryption.

Type: String

Required: No

**SSEType**

Server-side encryption type. The only supported value is:

- `KMS` \- Server-side encryption that uses AWS Key Management Service. The
key is stored in your account and is managed by AWS KMS (AWS KMS charges apply).

Type: String

Valid Values: `AES256 | KMS`

Required: No

**Status**

Represents the current state of server-side encryption. The only supported values
are:

- `ENABLED` \- Server-side encryption is enabled.

- `UPDATING` \- Server-side encryption is being updated.

Type: String

Valid Values: `ENABLING | ENABLED | DISABLING | DISABLED | UPDATING`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/ssedescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/ssedescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/ssedescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceTableFeatureDetails

SSESpecification

All content copied from https://docs.aws.amazon.com/.
