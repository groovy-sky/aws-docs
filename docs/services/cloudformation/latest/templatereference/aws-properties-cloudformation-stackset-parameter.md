This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::StackSet Parameter

The Parameter data type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ParameterKey" : String,
  "ParameterValue" : String
}

```

### YAML

```yaml

  ParameterKey: String
  ParameterValue: String

```

## Properties

`ParameterKey`

The key associated with the parameter. If you don't specify a key and value for a
particular parameter, CloudFormation uses the default value that's specified in
your template.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterValue`

The input value associated with the parameter.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OperationPreferences

StackInstances

All content copied from https://docs.aws.amazon.com/.
