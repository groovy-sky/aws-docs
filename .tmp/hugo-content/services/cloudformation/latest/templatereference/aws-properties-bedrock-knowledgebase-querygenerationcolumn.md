This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase QueryGenerationColumn

Contains information about a column in the current table for the query engine to consider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Inclusion" : String,
  "Name" : String
}

```

### YAML

```yaml

  Description: String
  Inclusion: String
  Name: String

```

## Properties

`Description`

A description of the column that helps the query engine understand the contents of the column.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Inclusion`

Specifies whether to include or exclude the column during query generation. If you specify `EXCLUDE`, the column will be ignored. If you specify `INCLUDE`, all other columns in the table will be ignored.

_Required_: No

_Type_: String

_Allowed values_: `INCLUDE | EXCLUDE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the column for which the other fields in this object apply.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PineconeFieldMapping

QueryGenerationConfiguration

All content copied from https://docs.aws.amazon.com/.
