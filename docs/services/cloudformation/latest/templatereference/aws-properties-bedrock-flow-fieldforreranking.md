This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow FieldForReranking

Specifies a field to be used during the reranking process in a Knowledge Base vector search. This structure identifies metadata fields that should be considered when reordering search results to improve relevance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldName" : String
}

```

### YAML

```yaml

  FieldName: String

```

## Properties

`FieldName`

The name of the metadata field to be used during the reranking process.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConditionFlowNodeConfiguration

FlowCondition

All content copied from https://docs.aws.amazon.com/.
