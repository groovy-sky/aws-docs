This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::Collection EncryptionConfig

Encryption settings for the collection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AWSOwnedKey" : Boolean,
  "KmsKeyArn" : String
}

```

### YAML

```yaml

  AWSOwnedKey: Boolean
  KmsKeyArn: String

```

## Properties

`AWSOwnedKey`

Indicates whether to use an AWS-owned key for encryption.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyArn`

The ARN of the AWS KMS key used to encrypt the collection.

_Required_: No

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):kms:[a-z0-9-]+:[0-9]{12}:key/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

_Minimum_: `10`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::OpenSearchServerless::Collection

FipsEndpoints

All content copied from https://docs.aws.amazon.com/.
