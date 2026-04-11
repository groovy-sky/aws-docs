This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::ProjectProfile EnvironmentConfigurationParameter

The environment configuration parameter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IsEditable" : Boolean,
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  IsEditable: Boolean
  Name: String
  Value: String

```

## Properties

`IsEditable`

Specifies whether the environment parameter is editable.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the environment configuration parameter.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z_][a-zA-Z0-9_]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the environment configuration parameter.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnvironmentConfiguration

EnvironmentConfigurationParametersDetails

All content copied from https://docs.aws.amazon.com/.
