This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::ScheduledQuery DimensionMapping

This type is used to map column(s) from the query result to a dimension in the destination
table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DimensionValueType" : String,
  "Name" : String
}

```

### YAML

```yaml

  DimensionValueType: String
  Name: String

```

## Properties

`DimensionValueType`

Type for the dimension: VARCHAR

_Required_: Yes

_Type_: String

_Allowed values_: `VARCHAR`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

Column name from query result.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Timestream::ScheduledQuery

ErrorReportConfiguration

All content copied from https://docs.aws.amazon.com/.
