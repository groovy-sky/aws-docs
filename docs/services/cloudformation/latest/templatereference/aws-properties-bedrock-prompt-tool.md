This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt Tool

Information about a tool that you can use with the Converse API. For more information,
see [Call a tool with the Converse API](../../../bedrock/latest/userguide/tool-use.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CachePoint" : CachePointBlock,
  "ToolSpec" : ToolSpecification
}

```

### YAML

```yaml

  CachePoint:
    CachePointBlock
  ToolSpec:
    ToolSpecification

```

## Properties

`CachePoint`

CachePoint to include in the tool configuration.

_Required_: No

_Type_: [CachePointBlock](aws-properties-bedrock-prompt-cachepointblock.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToolSpec`

The specfication for the tool.

_Required_: No

_Type_: [ToolSpecification](aws-properties-bedrock-prompt-toolspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TextS3Location

ToolChoice

All content copied from https://docs.aws.amazon.com/.
