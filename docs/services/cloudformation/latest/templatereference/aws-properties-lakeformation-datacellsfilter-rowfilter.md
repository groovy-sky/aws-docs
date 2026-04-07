This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::DataCellsFilter RowFilter

A PartiQL predicate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllRowsWildcard" : Json,
  "FilterExpression" : String
}

```

### YAML

```yaml

  AllRowsWildcard: Json
  FilterExpression: String

```

## Properties

`AllRowsWildcard`

A wildcard for all rows.

_Required_: No

_Type_: Json

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FilterExpression`

A filter expression.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ColumnWildcard

AWS::LakeFormation::DataLakeSettings
