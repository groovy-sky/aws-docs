This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent ParameterDetail

Contains details about a parameter in a function for an action group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Required" : Boolean,
  "Type" : String
}

```

### YAML

```yaml

  Description: String
  Required: Boolean
  Type: String

```

## Properties

`Description`

A description of the parameter. Helps the foundation model determine how to elicit the parameters from the user.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Required`

Whether the parameter is required for the agent to complete the function for action group invocation.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The data type of the parameter.

_Required_: Yes

_Type_: String

_Allowed values_: `string | number | integer | boolean | array`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OrchestrationExecutor

PromptConfiguration

All content copied from https://docs.aws.amazon.com/.
