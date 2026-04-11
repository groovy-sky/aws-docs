This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel FecOutputSettings

The settings for FEC.

The parent of this entity is UdpOutputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnDepth" : Integer,
  "IncludeFec" : String,
  "RowLength" : Integer
}

```

### YAML

```yaml

  ColumnDepth: Integer
  IncludeFec: String
  RowLength: Integer

```

## Properties

`ColumnDepth`

The parameter D from SMPTE 2022-1. The height of the FEC protection matrix. The number of transport stream
packets per column error correction packet. The number must be between 4 and 20, inclusive.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeFec`

Enables column only or column and row-based FEC.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowLength`

The parameter L from SMPTE 2022-1. The width of the FEC protection matrix. Must be between 1 and 20, inclusive.
If only Column FEC is used, then larger values increase robustness. If Row FEC is used, then this is the number of
transport stream packets per row error correction packet, and the value must be between 4 and 20, inclusive, if
includeFec is columnAndRow. If includeFec is column, this value must be 1 to 20, inclusive.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FeatureActivations

Fmp4HlsSettings

All content copied from https://docs.aws.amazon.com/.
