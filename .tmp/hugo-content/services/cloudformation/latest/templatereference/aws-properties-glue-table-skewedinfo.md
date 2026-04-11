This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Table SkewedInfo

Specifies skewed values in a table. Skewed values are those that occur with very high
frequency.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SkewedColumnNames" : [ String, ... ],
  "SkewedColumnValueLocationMaps" : Json,
  "SkewedColumnValues" : [ String, ... ]
}

```

### YAML

```yaml

  SkewedColumnNames:
    - String
  SkewedColumnValueLocationMaps: Json
  SkewedColumnValues:
    - String

```

## Properties

`SkewedColumnNames`

A list of names of columns that contain skewed values.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SkewedColumnValueLocationMaps`

A mapping of skewed values to the columns that contain them.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SkewedColumnValues`

A list of values that appear so frequently as to be considered
skewed.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SerdeInfo

StorageDescriptor

All content copied from https://docs.aws.amazon.com/.
