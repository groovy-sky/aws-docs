This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataSetStringListFilterValue

Represents a list of string values used in filter conditions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StaticValues" : [ String, ... ]
}

```

### YAML

```yaml

  StaticValues:
    - String

```

## Properties

`StaticValues`

A list of static string values used for filtering.

_Required_: No

_Type_: Array of String

_Minimum_: `0 | 0`

_Maximum_: `512 | 1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSetStringListFilterCondition

DataSetUsageConfiguration

All content copied from https://docs.aws.amazon.com/.
