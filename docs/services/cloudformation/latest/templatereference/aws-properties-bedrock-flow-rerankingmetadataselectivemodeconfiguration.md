This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow RerankingMetadataSelectiveModeConfiguration

Configuration for selectively including or excluding metadata fields during the reranking process. This allows you to control which metadata attributes are considered when reordering search results.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldsToExclude" : [ FieldForReranking, ... ],
  "FieldsToInclude" : [ FieldForReranking, ... ]
}

```

### YAML

```yaml

  FieldsToExclude:
    - FieldForReranking
  FieldsToInclude:
    - FieldForReranking

```

## Properties

`FieldsToExclude`

A list of metadata field names to explicitly exclude from the reranking process. All metadata fields except these will be considered when reordering search results. This parameter cannot be used together with fieldsToInclude.

_Required_: No

_Type_: Array of [FieldForReranking](aws-properties-bedrock-flow-fieldforreranking.md)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldsToInclude`

A list of metadata field names to explicitly include in the reranking process. Only these fields will be considered when reordering search results. This parameter cannot be used together with fieldsToExclude.

_Required_: No

_Type_: Array of [FieldForReranking](aws-properties-bedrock-flow-fieldforreranking.md)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PromptTemplateConfiguration

RetrievalFlowNodeConfiguration

All content copied from https://docs.aws.amazon.com/.
