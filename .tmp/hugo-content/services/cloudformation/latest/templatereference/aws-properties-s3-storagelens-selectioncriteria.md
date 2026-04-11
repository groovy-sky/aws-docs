This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens SelectionCriteria

This resource contains the details of the Amazon S3 Storage Lens selection
criteria.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Delimiter" : String,
  "MaxDepth" : Integer,
  "MinStorageBytesPercentage" : Number
}

```

### YAML

```yaml

  Delimiter: String
  MaxDepth: Integer
  MinStorageBytesPercentage: Number

```

## Properties

`Delimiter`

This property contains the details of the S3 Storage Lens delimiter being used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxDepth`

This property contains the details of the max depth that S3 Storage Lens will collect
metrics up to.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinStorageBytesPercentage`

This property contains the details of the minimum storage bytes percentage threshold that
S3 Storage Lens will collect metrics up to.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3BucketDestination

SSEKMS

All content copied from https://docs.aws.amazon.com/.
