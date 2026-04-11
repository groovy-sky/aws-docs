This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kinesis::Stream StreamEncryption

Enables or updates server-side encryption using an AWS KMS key for a
specified stream.

###### Note

When invoking this API, you must use either the `StreamARN` or the
`StreamName` parameter, or both. It is recommended that you use the
`StreamARN` input parameter when you invoke this API.

Starting encryption is an asynchronous operation. Upon receiving the request, Kinesis
Data Streams returns immediately and sets the status of the stream to
`UPDATING`. After the update is complete, Kinesis Data Streams sets the
status of the stream back to `ACTIVE`. Updating or applying encryption
normally takes a few seconds to complete, but it can take minutes. You can continue to
read and write data to your stream while its status is `UPDATING`. Once the
status of the stream is `ACTIVE`, encryption begins for records written to
the stream.

API Limits: You can successfully apply a new AWS KMS key for
server-side encryption 25 times in a rolling 24-hour period.

Note: It can take up to 5 seconds after the stream is in an `ACTIVE` status
before all records written to the stream are encrypted. After you enable encryption, you
can verify that encryption is applied by inspecting the API response from
`PutRecord` or `PutRecords`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionType" : String,
  "KeyId" : String
}

```

### YAML

```yaml

  EncryptionType: String
  KeyId: String

```

## Properties

`EncryptionType`

The encryption type to use. The only valid value is `KMS`.

_Required_: Yes

_Type_: String

_Allowed values_: `KMS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyId`

The GUID for the customer-managed AWS KMS key to use for encryption.
This value can be a globally unique identifier, a fully specified Amazon Resource Name
(ARN) to either an alias or a key, or an alias name prefixed by "alias/".You can also
use a master key owned by Kinesis Data Streams by specifying the alias
`aws/kinesis`.

- Key ARN example:
`arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012`

- Alias ARN example:
`arn:aws:kms:us-east-1:123456789012:alias/MyAliasName`

- Globally unique key ID example:
`12345678-1234-1234-1234-123456789012`

- Alias name example: `alias/MyAliasName`

- Master key owned by Kinesis Data Streams:
`alias/aws/kinesis`

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Kinesis::Stream

StreamModeDetails

All content copied from https://docs.aws.amazon.com/.
