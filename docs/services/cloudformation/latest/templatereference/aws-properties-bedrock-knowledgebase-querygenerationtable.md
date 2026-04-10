This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase QueryGenerationTable

Contains information about a table for the query engine to consider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Columns" : [ QueryGenerationColumn, ... ],
  "Description" : String,
  "Inclusion" : String,
  "Name" : String
}

```

### YAML

```yaml

  Columns:
    - QueryGenerationColumn
  Description: String
  Inclusion: String
  Name: String

```

## Properties

`Columns`

An array of objects, each of which defines information about a column in the table.

_Required_: No

_Type_: Array of [QueryGenerationColumn](aws-properties-bedrock-knowledgebase-querygenerationcolumn.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the table that helps the query engine understand the contents of the table.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Inclusion`

Specifies whether to include or exclude the table during query generation. If you specify `EXCLUDE`, the table will be ignored. If you specify `INCLUDE`, all other tables will be ignored.

_Required_: No

_Type_: String

_Allowed values_: `INCLUDE | EXCLUDE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the table for which the other fields in this object apply.

_Required_: Yes

_Type_: String

_Pattern_: `^.*\..*\..*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryGenerationContext

RdsConfiguration

All content copied from https://docs.aws.amazon.com/.
