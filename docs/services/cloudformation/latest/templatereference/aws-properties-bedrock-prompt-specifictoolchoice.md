This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt SpecificToolChoice

The model must request a specific tool. For example, `{"tool" : {"name" : "Your
            tool name"}}`. For more information, see [Call a tool with the Converse API](../../../bedrock/latest/userguide/tool-use.md)
in the Amazon Bedrock User Guide

###### Note

This field is only supported by Anthropic Claude 3 models.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String
}

```

### YAML

```yaml

  Name: String

```

## Properties

`Name`

The name of the tool that the model must request.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]*$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PromptVariant

SystemContentBlock

All content copied from https://docs.aws.amazon.com/.
