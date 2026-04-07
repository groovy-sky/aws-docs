This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet ColumnLevelPermissionRule

A rule defined to grant access on one or more restricted columns. Each dataset can have multiple rules. To
create a restricted column, you add it to one or more rules. Each rule must contain at least one column and at least
one user or group. To be able to see a restricted column, a user or group needs to be added to a rule for that
column.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnNames" : [ String, ... ],
  "Principals" : [ String, ... ]
}

```

### YAML

```yaml

  ColumnNames:
    - String
  Principals:
    - String

```

## Properties

`ColumnNames`

An array of column names.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Principals`

An array of Amazon Resource Names (ARNs) for Quick users or groups.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ColumnGroup

ColumnTag
