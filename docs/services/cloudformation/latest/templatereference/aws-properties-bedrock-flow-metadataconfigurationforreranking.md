This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow MetadataConfigurationForReranking

Configuration for how metadata should be used during the reranking process in Knowledge Base vector searches. This determines which metadata fields are included or excluded when reordering search results.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SelectionMode" : String,
  "SelectiveModeConfiguration" : RerankingMetadataSelectiveModeConfiguration
}

```

### YAML

```yaml

  SelectionMode: String
  SelectiveModeConfiguration:
    RerankingMetadataSelectiveModeConfiguration

```

## Properties

`SelectionMode`

The mode for selecting which metadata fields to include in the reranking process. Valid values are ALL (use all available metadata fields) or SELECTIVE (use only specified fields).

_Required_: Yes

_Type_: String

_Allowed values_: `SELECTIVE | ALL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectiveModeConfiguration`

Configuration for selective mode, which allows you to explicitly include or exclude specific metadata fields during reranking. This is only used when selectionMode is set to SELECTIVE.

_Required_: No

_Type_: [RerankingMetadataSelectiveModeConfiguration](aws-properties-bedrock-flow-rerankingmetadataselectivemodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoopFlowNodeConfiguration

PerformanceConfiguration

All content copied from https://docs.aws.amazon.com/.
