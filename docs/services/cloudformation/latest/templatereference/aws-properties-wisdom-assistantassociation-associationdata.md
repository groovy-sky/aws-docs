This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AssistantAssociation AssociationData

A union type that currently has a single argument, which is the knowledge base
ID.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExternalBedrockKnowledgeBaseConfig" : ExternalBedrockKnowledgeBaseConfig,
  "KnowledgeBaseId" : String
}

```

### YAML

```yaml

  ExternalBedrockKnowledgeBaseConfig:
    ExternalBedrockKnowledgeBaseConfig
  KnowledgeBaseId: String

```

## Properties

`ExternalBedrockKnowledgeBaseConfig`

Property description not available.

_Required_: No

_Type_: [ExternalBedrockKnowledgeBaseConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wisdom-assistantassociation-externalbedrockknowledgebaseconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KnowledgeBaseId`

The identifier of the knowledge base.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Wisdom::AssistantAssociation

ExternalBedrockKnowledgeBaseConfig
