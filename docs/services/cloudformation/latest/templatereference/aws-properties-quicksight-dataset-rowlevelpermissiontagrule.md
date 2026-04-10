This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet RowLevelPermissionTagRule

A set of rules associated with a tag.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnName" : String,
  "MatchAllValue" : String,
  "TagKey" : String,
  "TagMultiValueDelimiter" : String
}

```

### YAML

```yaml

  ColumnName: String
  MatchAllValue: String
  TagKey: String
  TagMultiValueDelimiter: String

```

## Properties

`ColumnName`

The column name that a tag key is assigned to.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchAllValue`

A string that you want to use to filter by all the values in a column in the dataset and don’t want to list the values one by one. For example, you can use an asterisk as your match all value.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagKey`

The unique key for a tag.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagMultiValueDelimiter`

A string that you want to use to delimit the values when you pass the values at run time. For example, you can delimit the values with a comma.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RowLevelPermissionTagConfiguration

S3Source

All content copied from https://docs.aws.amazon.com/.
