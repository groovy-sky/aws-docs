This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase QueryGenerationConfiguration

Contains configurations for query generation. For more information, see [Build a knowledge base by connecting to a structured data source](../../../bedrock/latest/userguide/knowledge-base-build-structured.md) in the Amazon Bedrock User Guide..

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExecutionTimeoutSeconds" : Integer,
  "GenerationContext" : QueryGenerationContext
}

```

### YAML

```yaml

  ExecutionTimeoutSeconds: Integer
  GenerationContext:
    QueryGenerationContext

```

## Properties

`ExecutionTimeoutSeconds`

The time after which query generation will time out.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GenerationContext`

Specifies configurations for context to use during query generation.

_Required_: No

_Type_: [QueryGenerationContext](aws-properties-bedrock-knowledgebase-querygenerationcontext.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryGenerationColumn

QueryGenerationContext

All content copied from https://docs.aws.amazon.com/.
