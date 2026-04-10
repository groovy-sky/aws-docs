This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::SecurityControl ParameterConfiguration

An object that provides the current value of a security control parameter and identifies whether it has been customized.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Value" : ParameterValue,
  "ValueType" : String
}

```

### YAML

```yaml

  Value:
    ParameterValue
  ValueType: String

```

## Properties

`Value`

The current value of a control parameter.

_Required_: No

_Type_: [ParameterValue](aws-properties-securityhub-securitycontrol-parametervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueType`

Identifies whether a control parameter uses a custom user-defined value or subscribes to the default
AWS Security Hub CSPM behavior.

When `ValueType` is set equal to `DEFAULT`, the default
behavior can be a specific Security Hub CSPM default value, or the default behavior can be to ignore a specific parameter.
When `ValueType` is set equal to `DEFAULT`, Security Hub CSPM ignores user-provided input for
the `Value` field.

When `ValueType` is set equal to `CUSTOM`, the `Value` field can't be empty.

_Required_: Yes

_Type_: String

_Allowed values_: `DEFAULT | CUSTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SecurityHub::SecurityControl

ParameterValue

All content copied from https://docs.aws.amazon.com/.
