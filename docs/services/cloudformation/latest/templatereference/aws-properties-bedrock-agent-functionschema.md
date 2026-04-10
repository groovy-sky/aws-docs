This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent FunctionSchema

Contains details about the function schema for the action group or the JSON or YAML-formatted payload defining the schema.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Functions" : [ Function, ... ]
}

```

### YAML

```yaml

  Functions:
    - Function

```

## Properties

`Functions`

A list of functions that each define an action in the action group.

_Required_: Yes

_Type_: Array of [Function](aws-properties-bedrock-agent-function.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Function

GuardrailConfiguration

All content copied from https://docs.aws.amazon.com/.
