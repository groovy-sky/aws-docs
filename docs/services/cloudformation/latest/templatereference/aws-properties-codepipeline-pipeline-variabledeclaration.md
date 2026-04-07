This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline VariableDeclaration

A variable declared at the pipeline level.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultValue" : String,
  "Description" : String,
  "Name" : String
}

```

### YAML

```yaml

  DefaultValue: String
  Description: String
  Name: String

```

## Properties

`DefaultValue`

The value of a pipeline-level variable.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of a pipeline-level variable. It's used to add additional context
about the variable, and not being used at time when pipeline executes.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a pipeline-level variable.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9@\-_]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::CodePipeline::Webhook
