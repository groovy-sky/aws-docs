This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard BinWidthOptions

The options that determine the bin width of a histogram.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BinCountLimit" : Number,
  "Value" : Number
}

```

### YAML

```yaml

  BinCountLimit: Number
  Value: Number

```

## Properties

`BinCountLimit`

The options that determine the bin count limit.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The options that determine the bin width value.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BinCountOptions

BodySectionConfiguration

All content copied from https://docs.aws.amazon.com/.
