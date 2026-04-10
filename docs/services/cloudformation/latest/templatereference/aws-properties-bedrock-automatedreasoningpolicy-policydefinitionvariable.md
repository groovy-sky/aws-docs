This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::AutomatedReasoningPolicy PolicyDefinitionVariable

A variable defined within the policy that can be used in rules.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Name" : String,
  "Type" : String
}

```

### YAML

```yaml

  Description: String
  Name: String
  Type: String

```

## Properties

`Description`

A description of a variable defined in the policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[\s\S]+$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a variable defined in the policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z][A-Za-z0-9_]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The data type of a variable defined in the policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z][A-Za-z0-9_]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PolicyDefinitionTypeValue

Tag

All content copied from https://docs.aws.amazon.com/.
