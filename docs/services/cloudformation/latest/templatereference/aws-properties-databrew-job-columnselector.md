This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Job ColumnSelector

Selector of a column from a dataset for profile job configuration.
One selector includes either a column name or a regular expression.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Regex" : String
}

```

### YAML

```yaml

  Name: String
  Regex: String

```

## Properties

`Name`

The name of a column from a dataset.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Regex`

A regular expression for selecting a column from a dataset.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AllowedStatistics

ColumnStatisticsConfiguration

All content copied from https://docs.aws.amazon.com/.
