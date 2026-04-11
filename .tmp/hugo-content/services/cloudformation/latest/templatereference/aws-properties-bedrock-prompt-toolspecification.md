This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt ToolSpecification

The specification for the tool. For more information, see [Call a tool with the Converse API](../../../bedrock/latest/userguide/tool-use.md)
in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "InputSchema" : ToolInputSchema,
  "Name" : String
}

```

### YAML

```yaml

  Description: String
  InputSchema:
    ToolInputSchema
  Name: String

```

## Properties

`Description`

The description for the tool.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputSchema`

The input schema for the tool in JSON format.

_Required_: Yes

_Type_: [ToolInputSchema](aws-properties-bedrock-prompt-toolinputschema.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name for the tool.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]*$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ToolInputSchema

AWS::Bedrock::PromptVersion

All content copied from https://docs.aws.amazon.com/.
