This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table SSESpecification

Represents the settings used to enable server-side encryption.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KMSMasterKeyId" : String,
  "SSEEnabled" : Boolean,
  "SSEType" : String
}

```

### YAML

```yaml

  KMSMasterKeyId: String
  SSEEnabled: Boolean
  SSEType: String

```

## Properties

`KMSMasterKeyId`

The AWS KMS key that should be used for the AWS KMS encryption.
To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN.
Note that you should only provide this parameter if the key is different from the
default DynamoDB key `alias/aws/dynamodb`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SSEEnabled`

Indicates whether server-side encryption is done using an AWS managed
key or an AWS owned key. If enabled (true), server-side encryption type
is set to `KMS` and an AWS managed key is used (AWS KMS charges apply). If disabled (false) or not specified, server-side
encryption is set to AWS owned key.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SSEType`

Server-side encryption type. The only supported value is:

- `KMS` \- Server-side encryption that uses AWS Key Management Service. The
key is stored in your account and is managed by AWS KMS (AWS KMS charges apply).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3BucketSource

StreamSpecification

All content copied from https://docs.aws.amazon.com/.
