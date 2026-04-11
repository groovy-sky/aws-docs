This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt ToolInputSchema

The schema for the tool. The top level schema type must be `object`. For more
information, see [Call a tool with the Converse API](../../../bedrock/latest/userguide/tool-use.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Json" : Json
}

```

### YAML

```yaml

  Json:
    Json

```

## Properties

`Json`

The JSON schema for the tool. For more information, see [JSON Schema Reference](https://json-schema.org/understanding-json-schema/reference).

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ToolConfiguration

ToolSpecification

All content copied from https://docs.aws.amazon.com/.
