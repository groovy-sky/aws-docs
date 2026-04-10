This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase CuratedQuery

Contains configurations for a query, each of which defines information about example queries to help the query engine generate appropriate SQL queries.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NaturalLanguage" : String,
  "Sql" : String
}

```

### YAML

```yaml

  NaturalLanguage: String
  Sql: String

```

## Properties

`NaturalLanguage`

An example natural language query.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sql`

The SQL equivalent of the natural language query.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BedrockEmbeddingModelConfiguration

EmbeddingModelConfiguration

All content copied from https://docs.aws.amazon.com/.
