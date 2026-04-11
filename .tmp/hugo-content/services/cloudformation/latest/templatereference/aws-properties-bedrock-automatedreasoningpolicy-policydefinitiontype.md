This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::AutomatedReasoningPolicy PolicyDefinitionType

A custom type definition within the policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Name" : String,
  "Values" : [ PolicyDefinitionTypeValue, ... ]
}

```

### YAML

```yaml

  Description: String
  Name: String
  Values:
    - PolicyDefinitionTypeValue

```

## Properties

`Description`

A description of the custom type defined in the policy.

_Required_: No

_Type_: String

_Pattern_: `^[\s\S]+$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a custom type defined in the policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z][A-Za-z0-9_]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The possible values for a custom type defined in the policy.

_Required_: Yes

_Type_: Array of [PolicyDefinitionTypeValue](aws-properties-bedrock-automatedreasoningpolicy-policydefinitiontypevalue.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PolicyDefinitionRule

PolicyDefinitionTypeValue

All content copied from https://docs.aws.amazon.com/.
