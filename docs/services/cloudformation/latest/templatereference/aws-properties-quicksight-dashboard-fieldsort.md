This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard FieldSort

The sort configuration for a field in a
field well.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Direction" : String,
  "FieldId" : String
}

```

### YAML

```yaml

  Direction: String
  FieldId: String

```

## Properties

`Direction`

The sort direction. Choose one of the following
options:

- `ASC`: Ascending

- `DESC`: Descending

_Required_: Yes

_Type_: String

_Allowed values_: `ASC | DESC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldId`

The sort configuration target field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldSeriesItem

FieldSortOptions

All content copied from https://docs.aws.amazon.com/.
