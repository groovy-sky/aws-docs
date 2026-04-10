This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt PromptMetadataEntry

Contains a key-value pair that defines a metadata tag and value to attach to a prompt variant. For more information, see [Create a prompt using Prompt management](../../../bedrock/latest/userguide/prompt-management-create.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The key of a metadata tag for a prompt variant.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of a metadata tag for a prompt variant.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PromptInputVariable

PromptModelInferenceConfiguration

All content copied from https://docs.aws.amazon.com/.
