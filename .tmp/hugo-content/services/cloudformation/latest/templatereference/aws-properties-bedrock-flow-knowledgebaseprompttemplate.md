This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow KnowledgeBasePromptTemplate

Defines a custom prompt template for orchestrating the retrieval and generation process.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TextPromptTemplate" : String
}

```

### YAML

```yaml

  TextPromptTemplate: String

```

## Properties

`TextPromptTemplate`

The text of the prompt template.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KnowledgeBaseOrchestrationConfiguration

LambdaFunctionFlowNodeConfiguration

All content copied from https://docs.aws.amazon.com/.
