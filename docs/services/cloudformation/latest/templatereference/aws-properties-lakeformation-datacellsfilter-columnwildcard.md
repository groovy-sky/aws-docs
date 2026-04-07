This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::DataCellsFilter ColumnWildcard

A wildcard object, consisting of an optional list of excluded column names or indexes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExcludedColumnNames" : [ String, ... ]
}

```

### YAML

```yaml

  ExcludedColumnNames:
    - String

```

## Properties

`ExcludedColumnNames`

Excludes column names. Any column with this name will be excluded.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::LakeFormation::DataCellsFilter

RowFilter
