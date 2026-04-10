This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis CalculatedMeasureField

The table calculation measure field for pivot tables.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Expression" : String,
  "FieldId" : String
}

```

### YAML

```yaml

  Expression: String
  FieldId: String

```

## Properties

`Expression`

The expression in the table calculation.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldId`

The custom field ID.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CalculatedField

CascadingControlConfiguration

All content copied from https://docs.aws.amazon.com/.
