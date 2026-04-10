This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::AutomatedReasoningPolicy PolicyDefinitionTypeValue

A value associated with a custom type in the policy definition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Value" : String
}

```

### YAML

```yaml

  Description: String
  Value: String

```

## Properties

`Description`

A description of the policy definition type value.

_Required_: No

_Type_: String

_Pattern_: `^[\s\S]+$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value associated with a policy definition type.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z][A-Za-z0-9_]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PolicyDefinitionType

PolicyDefinitionVariable

All content copied from https://docs.aws.amazon.com/.
