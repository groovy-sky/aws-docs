This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLensGroup MatchObjectSize

This resource filters objects that match the specified object size range.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BytesGreaterThan" : Integer,
  "BytesLessThan" : Integer
}

```

### YAML

```yaml

  BytesGreaterThan: Integer
  BytesLessThan: Integer

```

## Properties

`BytesGreaterThan`

This property specifies the minimum object size in bytes. The value must be a positive
number, greater than 0 and less than 5 TB.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BytesLessThan`

This property specifies the maximum object size in bytes. The value must be a positive
number, greater than the minimum object size and less than 5 TB.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MatchObjectAge

Or

All content copied from https://docs.aws.amazon.com/.
