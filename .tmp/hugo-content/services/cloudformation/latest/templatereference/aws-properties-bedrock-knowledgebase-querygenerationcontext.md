This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase QueryGenerationContext

>Contains configurations for context to use during query generation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CuratedQueries" : [ CuratedQuery, ... ],
  "Tables" : [ QueryGenerationTable, ... ]
}

```

### YAML

```yaml

  CuratedQueries:
    - CuratedQuery
  Tables:
    - QueryGenerationTable

```

## Properties

`CuratedQueries`

An array of objects, each of which defines information about example queries to help the query engine generate appropriate SQL queries.

_Required_: No

_Type_: Array of [CuratedQuery](aws-properties-bedrock-knowledgebase-curatedquery.md)

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tables`

An array of objects, each of which defines information about a table in the database.

_Required_: No

_Type_: Array of [QueryGenerationTable](aws-properties-bedrock-knowledgebase-querygenerationtable.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryGenerationConfiguration

QueryGenerationTable

All content copied from https://docs.aws.amazon.com/.
