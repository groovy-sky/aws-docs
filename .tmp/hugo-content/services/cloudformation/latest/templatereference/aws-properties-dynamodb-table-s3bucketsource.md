This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table S3BucketSource

The S3 bucket that is being imported from.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Bucket" : String,
  "S3BucketOwner" : String,
  "S3KeyPrefix" : String
}

```

### YAML

```yaml

  S3Bucket: String
  S3BucketOwner: String
  S3KeyPrefix: String

```

## Properties

`S3Bucket`

The S3 bucket that is being imported from.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9A-Z]+[\.\-\w]*[a-z0-9A-Z]+$`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3BucketOwner`

The account number of the S3 bucket that is being imported from. If the bucket is
owned by the requester this is optional.

_Required_: No

_Type_: String

_Pattern_: `[0-9]{12}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3KeyPrefix`

The key prefix shared by all S3 Objects that are being imported.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourcePolicy

SSESpecification

All content copied from https://docs.aws.amazon.com/.
